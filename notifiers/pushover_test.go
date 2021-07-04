package notifiers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"testing"
)

var (
	PUSHOVER_TOKEN string
	PUSHOVER_API   string
)

func TestPushoverNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)

	t.Parallel()
	PUSHOVER_TOKEN = utils.Params.GetString("PUSHOVER_TOKEN")
	PUSHOVER_API = utils.Params.GetString("PUSHOVER_API")

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	if PUSHOVER_TOKEN == "" || PUSHOVER_API == "" {
		t.Log("Pushover notifier testing skipped, missing PUSHOVER_TOKEN and PUSHOVER_API environment variable")
		t.SkipNow()
	}

	t.Run("Load Pushover", func(t *testing.T) {
		Pushover.ApiKey = null.NewNullString(PUSHOVER_TOKEN)
		Pushover.ApiSecret = null.NewNullString(PUSHOVER_API)
		Pushover.Var1 = null.NewNullString("Normal")
		Pushover.Var2 = null.NewNullString("vibrate")
		Pushover.Enabled = null.NewNullBool(true)

		Add(Pushover)

		assert.Nil(t, err)
		assert.Equal(t, "Hunter Long", Pushover.Author)
		assert.Equal(t, PUSHOVER_TOKEN, Pushover.ApiKey.String)
	})

	t.Run("Pushover Within Limits", func(t *testing.T) {
		assert.True(t, Pushover.CanSend())
	})

	t.Run("Pushover OnSave", func(t *testing.T) {
		_, err := Pushover.OnSave()
		assert.Nil(t, err)
	})

	t.Run("Pushover OnFailure", func(t *testing.T) {
		_, err := Pushover.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("Pushover OnSuccess", func(t *testing.T) {
		_, err := Pushover.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("Pushover Test", func(t *testing.T) {
		_, err := Pushover.OnTest()
		assert.Nil(t, err)
	})

}
