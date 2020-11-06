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
	DISCORD_URL string
)

func TestDiscordNotifier(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	DISCORD_URL = utils.Params.GetString("DISCORD_URL")

	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	if DISCORD_URL == "" {
		t.Log("discord notifier testing skipped, missing DISCORD_URL environment variable")
		t.SkipNow()
	}

	t.Run("Load discord", func(t *testing.T) {
		Discorder.Host = DISCORD_URL
		Discorder.Delay = time.Duration(100 * time.Millisecond)
		Discorder.Enabled = null.NewNullBool(true)

		Add(Discorder)

		assert.Equal(t, "Hunter Long", Discorder.Author)
		assert.Equal(t, DISCORD_URL, Discorder.Host)
	})

	t.Run("discord Notifier Tester", func(t *testing.T) {
		assert.True(t, Discorder.CanSend())
	})

	t.Run("discord Notifier Tester OnSave", func(t *testing.T) {
		_, err := Discorder.OnSave()
		assert.Nil(t, err)
	})

	t.Run("discord OnFailure", func(t *testing.T) {
		_, err := Discorder.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("discord OnSuccess", func(t *testing.T) {
		_, err := Discorder.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("discord Test", func(t *testing.T) {
		_, err := Discorder.OnTest()
		assert.Nil(t, err)
	})

}
