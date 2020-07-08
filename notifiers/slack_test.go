package notifiers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/torusresearch/statping/database"
	"github.com/torusresearch/statping/types/core"
	"github.com/torusresearch/statping/types/failures"
	"github.com/torusresearch/statping/types/notifications"
	"github.com/torusresearch/statping/types/null"
	"github.com/torusresearch/statping/types/services"
	"github.com/torusresearch/statping/utils"
	"testing"
	"time"
)

var (
	SLACK_URL string
)

func TestSlackNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	SLACK_URL = utils.Params.GetString("SLACK_URL")
	slacker.Host = SLACK_URL
	slacker.Enabled = null.NewNullBool(true)

	if SLACK_URL == "" {
		t.Log("slack notifier testing skipped, missing SLACK_URL environment variable")
		t.SkipNow()
	}

	t.Run("Load slack", func(t *testing.T) {
		slacker.Host = SLACK_URL
		slacker.Delay = time.Duration(100 * time.Millisecond)
		slacker.Limits = 3
		Add(slacker)
		assert.Equal(t, "Hunter Long", slacker.Author)
		assert.Equal(t, SLACK_URL, slacker.Host)
	})

	t.Run("slack Within Limits", func(t *testing.T) {
		ok := slacker.CanSend()
		assert.True(t, ok)
	})

	t.Run("slack OnSave", func(t *testing.T) {
		_, err := slacker.OnSave()
		assert.Nil(t, err)
	})

	t.Run("slack OnFailure", func(t *testing.T) {
		_, err := slacker.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("slack OnSuccess", func(t *testing.T) {
		_, err := slacker.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

}
