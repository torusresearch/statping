package database

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/statping/statping/types"
	"github.com/statping/statping/utils"
)

type GroupBy struct {
	db    Database
	query *GroupQuery
}

type GroupByer interface {
	ToTimeValue() (*TimeVar, error)
}

type By string

func (b By) String() string {
	return string(b)
}

type GroupQuery struct {
	Start     time.Time
	End       time.Time
	Group     time.Duration
	Order     string
	Limit     int
	Offset    int
	FillEmpty bool

	db Database
}

func (b GroupQuery) Find(data interface{}) error {
	return b.db.Order("id DESC").Find(data).Error()
}

func (b GroupQuery) Database() Database {
	return b.db
}

var (
	ByCount   = By("COUNT(id) as amount")
	ByAverage = func(column string, multiplier int) By {
		switch database.DbType() {
		case "mysql":
			return By(fmt.Sprintf("CAST(AVG(%s) as UNSIGNED INT) as amount", column))
		case "postgres":
			return By(fmt.Sprintf("cast(AVG(%s) as int) as amount", column))
		default:
			return By(fmt.Sprintf("cast(AVG(%s) as int) as amount", column))
		}
	}
)

type TimeVar struct {
	g    *GroupQuery
	data []*TimeValue
}

func (t *TimeVar) ToValues() ([]*TimeValue, error) {
	return t.data, nil
}

// GraphData will return all hits or failures
func (b *GroupQuery) GraphData(by By) ([]*TimeValue, error) {
	b.db = b.db.MultipleSelects(
		b.db.SelectByTime(b.Group),
		by.String(),
	).Group("timeframe").Order("timeframe", true)

	caller, err := b.ToTimeValue()
	if err != nil {
		return nil, err
	}

	if b.FillEmpty {
		return caller.FillMissing(b.Start, b.End)
	}
	return caller.ToValues()
}

// ToTimeValue will format the SQL rows into a JSON format for the API.
// [{"timestamp": "2006-01-02T15:04:05Z", "amount": 468293}]
// TODO redo this entire function, use better SQL query to group by time
func (b *GroupQuery) ToTimeValue() (*TimeVar, error) {
	rows, err := b.db.Rows()
	if err != nil {
		return nil, err
	}
	var data []*TimeValue
	for rows.Next() {
		var timeframe string
		var amount int64
		if err := rows.Scan(&timeframe, &amount); err != nil {
			log.Error(err, timeframe)
		}
		trueTime, _ := b.db.ParseTime(timeframe)
		newTs := types.FixedTime(trueTime, b.Group)

		tv := &TimeValue{
			Timeframe: newTs,
			Amount:    amount,
		}
		data = append(data, tv)
	}
	return &TimeVar{b, data}, nil
}

func (t *TimeVar) FillMissing(current, end time.Time) ([]*TimeValue, error) {
	timeMap := make(map[string]int64)
	var validSet []*TimeValue
	for _, v := range t.data {
		timeMap[v.Timeframe] = v.Amount
	}

	for {
		currentStr := types.FixedTime(current, t.g.Group)

		var amount int64
		if timeMap[currentStr] != 0 {
			amount = timeMap[currentStr]
		}

		validSet = append(validSet, &TimeValue{
			Timeframe: currentStr,
			Amount:    amount,
		})
		current = current.Add(t.g.Group)
		if current.After(end) {
			break
		}
	}

	return validSet, nil
}

type isObject interface {
	Db() Database
}

func ParseRequest(r *http.Request) (*GroupQuery, error) {
	fields := parseGet(r)
	grouping := fields.Get("group")
	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))
	limit := utils.ToInt(fields.Get("limit"))
	offset := utils.ToInt(fields.Get("offset"))
	fill, _ := strconv.ParseBool(fields.Get("fill"))
	orderBy := fields.Get("order")
	if limit == 0 {
		limit = 10000
	}

	if grouping == "" {
		grouping = "1h"
	}
	groupDur, err := time.ParseDuration(grouping)
	if err != nil {
		log.Errorln(err)
		groupDur = 1 * time.Hour
	}

	query := &GroupQuery{
		Start:     time.Unix(startField, 0).UTC(),
		End:       time.Unix(endField, 0).UTC(),
		Group:     groupDur,
		Order:     orderBy,
		Limit:     int(limit),
		Offset:    int(offset),
		FillEmpty: fill,
	}

	if query.Start.After(query.End) {
		return nil, errors.New("start time is after ending time")
	}

	return query, nil
}

func ParseQueries(r *http.Request, o isObject) (*GroupQuery, error) {
	fields := parseGet(r)
	grouping := fields.Get("group")
	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))
	limit := utils.ToInt(fields.Get("limit"))
	offset := utils.ToInt(fields.Get("offset"))
	fill, _ := strconv.ParseBool(fields.Get("fill"))
	orderBy := fields.Get("order")
	if limit == 0 {
		limit = 10000
	}

	q := o.Db()

	if grouping == "" {
		grouping = "1h"
	}
	groupDur, err := time.ParseDuration(grouping)
	if err != nil {
		log.Errorln(err)
		groupDur = 1 * time.Hour
	}
	if endField == 0 {
		endField = utils.Now().Unix()
	}

	query := &GroupQuery{
		Start:     time.Unix(startField, 0).UTC(),
		End:       time.Unix(endField, 0).UTC(),
		Group:     groupDur,
		Order:     orderBy,
		Limit:     int(limit),
		Offset:    int(offset),
		FillEmpty: fill,
		db:        q,
	}

	if query.Start.After(query.End) {
		return nil, errors.New("start time is after ending time")
	}

	if startField == 0 {
		query.Start = utils.Now().Add(-7 * types.Day)
	}
	if endField == 0 {
		query.End = utils.Now()
	}
	if query.Limit != 0 {
		q = q.Limit(query.Limit)
	}
	if query.Offset > 0 {
		q = q.Offset(query.Offset)
	}

	q = q.Where("created_at BETWEEN ? AND ?", q.FormatTime(query.Start), q.FormatTime(query.End))

	if query.Order != "" {
		q = q.Order(query.Order)
	}
	query.db = q

	return query, nil
}

func parseForm(r *http.Request) url.Values {
	r.ParseForm()
	return r.PostForm
}

func parseGet(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
