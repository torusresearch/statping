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

func TestCommandNotifier(t *testing.T) {
	t.SkipNow()
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&notifications.Notification{})
	notifications.SetDB(db)
	core.Example()

	t.Run("Load Command", func(t *testing.T) {
		Command.Host = "/bin/echo"
		Command.Var1 = "service {{.Service.Domain}} is online"
		Command.Var2 = "service {{.Service.Domain}} is offline"
		Command.Delay = time.Duration(100 * time.Millisecond)
		Command.Limits = 99
		Command.Enabled = null.NewNullBool(true)

		Add(Command)

		assert.Equal(t, "Hunter Long", Command.Author)
		assert.Equal(t, "/bin/echo", Command.Host)
	})

	t.Run("Command Notifier Tester", func(t *testing.T) {
		assert.True(t, Command.CanSend())
	})

	t.Run("Command OnSave", func(t *testing.T) {
		_, err := Command.OnSave()
		assert.Nil(t, err)
	})

	t.Run("Command OnFailure", func(t *testing.T) {
		_, err := Command.OnFailure(services.Example(false), failures.Example())
		assert.Nil(t, err)
	})

	t.Run("Command OnSuccess", func(t *testing.T) {
		_, err := Command.OnSuccess(services.Example(true))
		assert.Nil(t, err)
	})

	t.Run("Command Test", func(t *testing.T) {
		_, err := Command.OnTest()
		assert.Nil(t, err)
	})

}
