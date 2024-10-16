package datasource

import (
	"baize/app/setting"
	"baize/app/utils/logger"
	"fmt"
	"github.com/baizeplus/sqly"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is datasource providers.
var ProviderSet = wire.NewSet(NewData)

// NewData .
func NewData(d *setting.Datasource) (*sqly.DB, func(), error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.Mysql.User, d.Mysql.Password, d.Mysql.Host, d.Mysql.Port, d.Mysql.DB) + "?parseTime=true&loc=Asia%2FShanghai"
	fmt.Println(dsn)
	db, err := sqly.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(d.Mysql.MaxOpenConns)
	db.SetMaxIdleConns(d.Mysql.MaxIdleConns)
	sqly.SetLog(new(logger.SqlyLog))
	return db, nil, err
}
