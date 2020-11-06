package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/torusresearch/statping/database"
	"github.com/torusresearch/statping/types/checkins"
	"github.com/torusresearch/statping/types/core"
	"github.com/torusresearch/statping/types/failures"
	"github.com/torusresearch/statping/types/groups"
	"github.com/torusresearch/statping/types/hits"
	"github.com/torusresearch/statping/types/incidents"
	"github.com/torusresearch/statping/types/messages"
	"github.com/torusresearch/statping/types/notifications"
	"github.com/torusresearch/statping/types/null"
	"github.com/torusresearch/statping/types/services"
	"github.com/torusresearch/statping/types/users"
	"github.com/torusresearch/statping/utils"
	"time"
)

// Connect will attempt to connect to the sqlite, postgres, or mysql database
func Connect(configs *DbConfig, retry bool) error {
	conn := configs.ConnectionString()

	log.WithFields(utils.ToFields(configs, conn)).Debugln("attempting to connect to database")

	dbSession, err := database.Openw(configs.DbConn, conn)
	if err != nil {
		log.Errorf(fmt.Sprintf("Database connection error %s", err))
		if retry {
			log.Warnln(fmt.Sprintf("Database %s connection to '%s' is not available, trying again in 5 seconds...", configs.DbConn, configs.DbHost))
			time.Sleep(5 * time.Second)
			return Connect(configs, retry)
		} else {
			return err
		}
	}

	configs.ApiSecret = utils.Params.GetString("API_SECRET")

	log.WithFields(utils.ToFields(dbSession)).Debugln("connected to database")

	db := dbSession.DB()
	db.SetMaxOpenConns(utils.Params.GetInt("MAX_OPEN_CONN"))
	db.SetMaxIdleConns(utils.Params.GetInt("MAX_IDLE_CONN"))
	db.SetConnMaxLifetime(utils.Params.GetDuration("MAX_LIFE_CONN"))

	if db.Ping() == nil {
		if utils.VerboseMode >= 4 {
			dbSession.LogMode(true).Debug().SetLogger(gorm.Logger{log})
		}
		log.Infoln(fmt.Sprintf("Database %s connection was successful.", configs.DbConn))
	}

	configs.Db = dbSession

	initModels(configs.Db)

	return err
}

func initModels(db database.Database) {
	core.SetDB(db)
	services.SetDB(db)
	hits.SetDB(db)
	failures.SetDB(db)
	checkins.SetDB(db)
	notifications.SetDB(db)
	incidents.SetDB(db)
	users.SetDB(db)
	messages.SetDB(db)
	groups.SetDB(db)
}

func CreateAdminUser(c *DbConfig) error {
	adminUser := utils.Params.GetString("ADMIN_USER")
	adminPass := utils.Params.GetString("ADMIN_PASSWORD")
	adminEmail := utils.Params.GetString("ADMIN_EMAIL")

	if adminUser == "" || adminPass == "" {
		adminUser = "admin"
		adminPass = "admin"
	}

	admin := &users.User{
		Username: adminUser,
		Password: adminPass,
		Email:    adminEmail,
		Admin:    null.NewNullBool(true),
	}

	if err := admin.Create(); err != nil {
		return errors.Wrap(err, "error creating admin")
	}

	return nil
}
