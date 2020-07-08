package configs

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"github.com/torusresearch/statping/source"
	"github.com/torusresearch/statping/types/notifications"
	"github.com/torusresearch/statping/utils"

	"github.com/torusresearch/statping/types/checkins"
	"github.com/torusresearch/statping/types/core"
	"github.com/torusresearch/statping/types/failures"
	"github.com/torusresearch/statping/types/groups"
	"github.com/torusresearch/statping/types/hits"
	"github.com/torusresearch/statping/types/incidents"
	"github.com/torusresearch/statping/types/messages"
	"github.com/torusresearch/statping/types/services"
	"github.com/torusresearch/statping/types/users"
)

func (d *DbConfig) DatabaseChanges() error {
	var cr core.Core
	d.Db.Model(&core.Core{}).Find(&cr)

	if latestMigration > cr.MigrationId {
		log.Infof("Statping database is out of date, migrating to: %d", latestMigration)

		switch d.Db.DbType() {
		case "mysql":
			if err := d.genericMigration("MODIFY", false); err != nil {
				return err
			}
		case "postgres":
			if err := d.genericMigration("ALTER", true); err != nil {
				return err
			}
		default:
			if err := d.sqliteMigration(); err != nil {
				return err
			}
		}

		if err := d.Db.Exec(fmt.Sprintf("UPDATE core SET migration_id = %d", latestMigration)).Error(); err != nil {
			return err
		}

		if err := d.BackupAssets(); err != nil {
			return err
		}
	}
	return nil
}

// BackupAssets is a temporary function (to version 0.90.*) to backup your customized theme
// to a new folder called 'assets_backup'.
func (d *DbConfig) BackupAssets() error {
	if source.UsingAssets(utils.Directory) {
		log.Infof("Backing up 'assets' folder to 'assets_backup'")
		if err := utils.RenameDirectory(utils.Directory+"/assets", utils.Directory+"/assets_backup"); err != nil {
			return err
		}
		log.Infof("Old assets are now stored in: " + utils.Directory + "/assets_backup")
	}
	return nil
}

//MigrateDatabase will migrate the database structure to current version.
//This function will NOT remove previous records, tables or columns from the database.
//If this function has an issue, it will ROLLBACK to the previous state.
func (d *DbConfig) MigrateDatabase() error {
	var DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}}

	log.Infoln("Migrating Database Tables...")
	tx := d.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, table := range DbModels {
		tx = tx.AutoMigrate(table)
		if tx.Error() != nil {
			log.Errorln(tx.Error())
			return tx.Error()
		}
	}

	log.Infof("Migrating App to version: %s", core.App.Version)
	if err := tx.Table("core").AutoMigrate(&core.Core{}); err.Error() != nil {
		tx.Rollback()
		log.Errorln(fmt.Sprintf("Statping Database could not be migrated: %v", tx.Error()))
		return tx.Error()
	}

	if err := tx.Commit().Error(); err != nil {
		return err
	}

	d.Db.Table("core").Model(&core.Core{}).Update("version", core.App.Version)

	log.Infoln("Statping Database Tables Migrated")

	if err := d.Db.Model(&hits.Hit{}).AddIndex("idx_service_hit", "service").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&hits.Hit{}).AddIndex("hit_created_at", "created_at").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&failures.Failure{}).AddIndex("fail_created_at", "created_at").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&failures.Failure{}).AddIndex("idx_service_fail", "service").Error(); err != nil {
		log.Errorln(err)
	}

	if err := d.Db.Model(&failures.Failure{}).AddIndex("idx_checkin_fail", "checkin").Error(); err != nil {
		log.Errorln(err)
	}
	log.Infoln("Database Indexes Created")

	return nil
}
