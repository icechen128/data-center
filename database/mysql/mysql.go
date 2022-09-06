package mysql

import (
	"data-center/database/common"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DriverName = "mysql"

type Mysql struct {
	common.Common
}

func New(host, port, user, password, dbName string) Mysql {
	return Mysql{
		Common: common.Common{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			DBName:   dbName,
		},
	}
}

func (c *Mysql) Open() error {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.DBName)

	db, err := sql.Open(DriverName, dataSource)
	if err != nil {
		return err
	}
	c.DB = db
	return nil
}
