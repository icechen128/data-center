package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type Core struct {
	db *sql.DB // db实例

	driverName string // 数据库驱动
	host       string // 数据库地址
	port       string // 数据库端口
	user       string // 数据库用户名
	password   string // 数据库密码
	dbName     string // 数据库名
}

func NewDatabase(driverName string, host, port, user, password, dbName string) Core {
	return Core{
		driverName: driverName,
		host:       host,
		port:       port,
		user:       user,
		password:   password,
		dbName:     dbName,
	}
}

func (c *Core) Open() error {
	var dataSource string
	switch c.driverName {
	case "mysql":
		dataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			c.user, c.password, c.host, c.port, c.dbName)
	default:
		return errors.New("unknown driver name")
	}

	db, err := sql.Open(c.driverName, dataSource)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *Core) Close() error {
	return c.db.Close()
}

func (c *Core) DBStruct(dbName string) map[string]string {
	sqlStr := `SELECT TABLE_NAME tableName ,TABLE_COMMENT tableDesc
From INFORMATION_SCHEMA.TABLES
WHERE UPPER(table_type)='BASE TABLE'
AND LOWER(table_schema) = ?
ORDER BY table_name`
	var result = make(map[string]string)
	rows, err := c.db.Query(sqlStr, dbName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for rows.Next() {
		var tableName, tableDesc string
		err = rows.Scan(&tableName, &tableDesc)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		result[tableName] = tableDesc
	}
	return result
}

type Field struct {
	DatabaseName           string       `json:"database_name"`            // 数据库名称 TABLE_SCHEMA
	TableName              string       `json:"table_name"`               // 数据表名称 TABLE_NAME
	ColumnName             *string      `json:"column_name"`              // 字段名 COLUMN_NAME
	OrdinalPosition        int64        `json:"ordinal_position"`         // 排序序号 ORDINAL_POSITION
	ColumnDefault          *interface{} `json:"column_default"`           // 字段默认值 COLUMN_DEFAULT
	IsNullable             string       `json:"is_nullable"`              // 是否可以为 NULL IS_NULLABLE
	DataType               *string      `json:"data_type"`                // 字段数据类型 DATA_TYPE
	CharacterMaximumLength *int64       `json:"character_maximum_length"` // 字段的最大字符数 CHARACTER_MAXIMUM_LENGTH
	CharacterOctetLength   *int64       `json:"character_octet_length"`   // 字段的最大字节数 CHARACTER_OCTET_LENGTH
	NumericPrecision       *int64       `json:"numeric_precision"`        // 数字精度 NUMERIC_PRECISION
	NumericScale           *int64       `json:"numeric_scale"`            // 小数位数 NUMERIC_SCALE
	DateTimePrecision      *int64       `json:"date_time_precision"`      // 日期精度（datetime 类型和 SQL-92interval 类型数据库的子类型代码） DATETIME_PRECISION
	CharacterSetName       *string      `json:"character_set_name"`       // 字符集 CHARACTER_SET_NAME
	CollationName          *string      `json:"collation_name"`           // 字符集排序规则 COLLATION_NAME
	ColumnType             string       `json:"column_type"`              // 字段类型 COLUMN_TYPE
	ColumnKey              string       `json:"column_key"`               // 索引类型 COLUMN_KEY
	Extra                  *string      `json:"extra"`                    // 其它信息(auto_increment 等) EXTRA
	Privileges             *string      `json:"privileges"`               // 权限 PRIVILEGES
	ColumnComment          string       `json:"column_comment"`           // 字段注释 COLUMN_COMMENT
}

func (c *Core) TableStruct(dbName, tableName string) []Field {
	sqlStr := `SELECT TABLE_SCHEMA, TABLE_NAME, COLUMN_NAME, ORDINAL_POSITION, COLUMN_DEFAULT, IS_NULLABLE, DATA_TYPE,
    	CHARACTER_MAXIMUM_LENGTH, CHARACTER_OCTET_LENGTH, NUMERIC_PRECISION, NUMERIC_SCALE, DATETIME_PRECISION, CHARACTER_SET_NAME,
       	COLLATION_NAME, COLUMN_TYPE, COLUMN_KEY, EXTRA, PRIVILEGES, COLUMN_COMMENT
		FROM information_schema.columns
		WHERE table_schema = ? AND table_name = ?`

	var result []Field
	rows, err := c.db.Query(sqlStr, dbName, tableName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for rows.Next() {
		var f Field
		err = rows.Scan(&f.DatabaseName, &f.TableName, &f.ColumnName, &f.OrdinalPosition, &f.ColumnDefault,
			&f.IsNullable, &f.DataType, &f.CharacterMaximumLength, &f.CharacterOctetLength, &f.NumericPrecision,
			&f.NumericScale, &f.DateTimePrecision, &f.CharacterSetName, &f.CollationName, &f.ColumnType,
			&f.ColumnKey, &f.Extra, &f.Privileges, &f.ColumnComment)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		result = append(result, f)
	}

	return result
}

func (c Core) Find(tableName string, where string) {
	rows, err := c.db.Query("select * from ? where name = ?", tableName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	rows.Next()
}
