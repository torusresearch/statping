package notifications

import (
	"github.com/torusresearch/statping/types/metrics"
	"github.com/torusresearch/statping/utils"
)

// AfterFind for Notification will set the timezone
func (n *Notification) AfterFind() (err error) {
	n.CreatedAt = utils.Now()
	n.UpdatedAt = utils.Now()
	metrics.Query("notifier", "find")
	return
}

func (n *Notification) AfterCreate() {
	metrics.Query("notifier", "create")
}

func (n *Notification) AfterUpdate() {
	metrics.Query("notifier", "update")
}

func (n *Notification) AfterDelete() {
	metrics.Query("notifier", "delete")
}
