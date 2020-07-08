package failures

import (
	"github.com/torusresearch/statping/database"
	"github.com/torusresearch/statping/types/metrics"
)

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Failure{})
}

func DB() database.Database {
	return db
}

func (f *Failure) AfterFind() {
	metrics.Query("failure", "find")
}

func (f *Failure) AfterUpdate() {
	metrics.Query("failure", "update")
}

func (f *Failure) AfterDelete() {
	metrics.Query("failure", "delete")
}

func (f *Failure) AfterCreate() {
	metrics.Query("failure", "create")
}

func All() []*Failure {
	var failures []*Failure
	db.Find(&failures)
	return failures
}

func (f *Failure) Create() error {
	q := db.Create(f)
	return q.Error()
}

func (f *Failure) Update() error {
	q := db.Update(f)
	return q.Error()
}

func (f *Failure) Delete() error {
	q := db.Delete(f)
	return q.Error()
}
