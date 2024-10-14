package svc

import (
	"fmt"
	"github.com/smallq_class/app/internal/config"
	"github.com/smallq_class/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logger.Silent,          // Log level
			IgnoreRecordNotFoundError: true,                   // 忽略 ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                  // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		model.Banner{},
		model.Student{},
		model.ClassInfo{},
		model.ClassInfoStu{},
		model.ClassInfoTalk{},
		model.ClassRoom{},
		model.Course{},
		model.CoursePackage{},
		model.CoursePackageLog{},
		model.Family{},
		model.Notice{},
		model.Teacher{},
		model.User{},
	)
	if err != nil {
		fmt.Printf("AutoMigrate ERR: %s \n", err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
