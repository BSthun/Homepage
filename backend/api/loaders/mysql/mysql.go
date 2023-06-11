package mysql

import (
	"backend/loaders/config"
	"log"
	"os"
	model2 "share/types/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	loggerUtils "backend/utils/logger"
)

var DB *gorm.DB
var StrapiDB *gorm.DB

func Init() {
	// Initialize GORM instance using previously opened SQL connection
	gormLogLevel := []logger.LogLevel{
		logger.Silent,
		logger.Error,
		logger.Error,
		logger.Warn,
		logger.Info,
		logger.Info,
		logger.Info,
	}

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             100 * time.Millisecond,
			LogLevel:                  gormLogLevel[config.C.LogLevel],
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open SQL connection
	dialector := mysql.New(
		mysql.Config{
			DSN:               config.C.MySqlDsn,
			DefaultStringSize: 255,
		},
	)

	strapiDialector := mysql.New(
		mysql.Config{
			DSN:               config.C.MySqlStrapiDsn,
			DefaultStringSize: 255,
		},
	)

	// * Open main database connection
	if db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	}); err != nil {
		loggerUtils.Log(logrus.Fatal, "UNABLE TO LOAD GORM MYSQL DATABASE")
	} else {
		DB = db
	}

	// * Open strapi database connection
	if db, err := gorm.Open(strapiDialector, &gorm.Config{
		Logger: gormLogger,
	}); err != nil {
		loggerUtils.Log(logrus.Fatal, "UNABLE TO LOAD STRAPI GORM MYSQL DATABASE")
	} else {
		StrapiDB = db
	}

	// Initialize model migrations
	if config.C.MySqlMigrate {
		if err := migrate(); err != nil {
			loggerUtils.Log(logrus.Fatal, "UNABLE TO MIGRATE GORM MODEL")
		}
	}

	loggerUtils.Log(logrus.Debug, "INITIALIZED MYSQL CONNECTION")
}

func migrate() error {
	// * Migrate model
	if err := DB.AutoMigrate(
		new(model2.PhotoAlbum),
		new(model2.PhotoItem),
		new(model2.PhotoPerson),
		new(model2.PhotoSection),
		new(model2.PhotoSessionAlbum),
		new(model2.DiarySpotifyAnalytics),
		new(model2.TrackSession),
		new(model2.TrackLog),
	); err != nil {
		return err
	}

	return nil
}
