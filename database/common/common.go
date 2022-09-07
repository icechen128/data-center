package common

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Common struct {
	DB *sqlx.DB // db实例

	Host     string // 数据库地址
	Port     string // 数据库端口
	User     string // 数据库用户名
	Password string // 数据库密码
	DBName   string // 数据库名

	Instance InstacneInterface
}

type InstacneInterface interface {
	DSN() (driverName string, dataSource string)
}

func (c *Common) Open() error {
	driverName, dataSource := c.Instance.DSN()
	db, err := sqlx.Connect(driverName, dataSource)
	if err != nil {
		return err
	}
	c.DB = db
	return nil
}

func (c Common) Close() error {
	return c.DB.Close()
}

func (c Common) Find(tableName string, page, pageSize int64, soryBy string, where ...string) error {
	rows, err := c.DB.Query("select * from ?", tableName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan()
	}

	return nil
}
