package svc

import (
	"github.com/smallq_class/internal/model"
	"github.com/smallq_class/web/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
	//		LogLevel:                  logger.Silent,          // Log level
	//		IgnoreRecordNotFoundError: true,                   // 忽略 ErrRecordNotFound（记录未找到）错误
	//		Colorful:                  false,                  // 禁用彩色打印
	//	},
	//)
	Db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = Db.AutoMigrate(
		model.Family{},
		model.Parent{},
		model.Banner{},
		model.ClassInfo{},
		model.ClassInfoStu{},
		model.ClassInfoTalk{},
		model.Student{},
		model.Teacher{},
		model.ClassRoom{},
		model.Course{},
		model.CoursePackage{},
		model.CoursePackageLog{},
		model.Notice{},
	)
	return &ServiceContext{
		Config: c,
		DB:     Db,
	}
}
