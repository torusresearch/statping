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
	"time"
)

var (
	telegramToken   string
	telegramChannel string
)

func TestTelegramNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)

	t.Parallel()

	telegramToken = utils.Params.GetString("TELEGRAM_TOKEN")
	telegramChannel = utils.Params.GetString("TELEGRAM_CHANNEL")
	if telegramToken == "" || telegramChannel == "" {
		t.Log("Telegram notifier testing skipped, missing TELEGRAM_TOKEN and TELEGRAM_CHANNEL environment variable")
		t.SkipNow()
	}

	Telegram.ApiSecret = null.NewNullString(telegramToken)
	Telegram.Var1 = null.NewNullString(telegramChannel)

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	t.Run("Load Telegram", func(t *testing.T) {
		Telegram.ApiSecret = null.NewNullString(telegramToken)
		Telegram.Var1 = null.NewNullString(telegramChannel)
		Telegram.Delay = time.Duration(1 * time.Second)
		Telegram.Enabled = null.NewNullBool(true)

		Add(Telegram)

		assert.Equal(t, "Hunter Long", Telegram.Author)
		assert.Equal(t, telegramToken, Telegram.ApiSecret.String)
		assert.Equal(t, telegramChannel, Telegram.Var1.String)
	})

	t.Run("Telegram Within Limits", func(t *testing.T) {
		assert.True(t, Telegram.CanSend())
	})

	t.Run("Telegram OnSave", func(t *testing.T) {
		_, err := Telegram.OnSave()
		assert.Nil(t, err)
	})

	t.Run("Telegram OnFailure", func(t *testing.T) {
		_, err := Telegram.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("Telegram OnSuccess", func(t *testing.T) {
		_, err := Telegram.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("Telegram Test", func(t *testing.T) {
		_, err := Telegram.OnTest()
		assert.Nil(t, err)
	})

}
