package checkins

import (
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
