package checkins

import (
	_ "github.com/torusresearch/statping/utils"
	"time"

	"github.com/statping/statping/types/failures"
)

func (c *Checkin) CreateFailure(f *failures.Failure) error {
	f.Checkin = c.Id
	c.Failing = true
	return failures.DB().Create(f).Error()
}

func (c *Checkin) FailuresColumnID() (string, int64) {
	return "checkin", c.Id
}

func (c *Checkin) Failures() failures.Failurer {
	return failures.AllFailures(c)
}

func (c *Checkin) FailuresSince(t time.Time) failures.Failurer {
	return failures.Since(t, c)
}

func (c *Checkin) RecordSuccess() error {
	//service, err := services.Find(c.ServiceId)
	//if err != nil {
	//	return fmt.Errorf("couldn't find the corresponding service for the checkin: %v",err)
	//}
	//
	//service.Online = true
	//service.LastOnline = utils.Now()
	//if err := service.Update(); err != nil {
	//	return fmt.Errorf("couldn't update the service for the checkin: %v",err)
	//}

	return nil
}

func (c *Checkin) RecordFailure() error {
	//service, err := services.Find(c.ServiceId)
	//if err != nil {
	//	return fmt.Errorf("couldn't find the corresponding service for the checkin: %v",err)
	//}
	//
	//service.Online = false
	//service.LastOffline = utils.Now()
	//if err := service.Update(); err != nil {
	//	return fmt.Errorf("couldn't update the service for the checkin: %v",err)
	//}

	return nil
}
