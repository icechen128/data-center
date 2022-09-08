package mysql

import (
	"fmt"
	"github.com/icechen128/data-center/internal/pkg/database/common"
	"github.com/spf13/cast"

	_ "github.com/go-sql-driver/mysql"
)

const DriverName = "mysql"

type Mysql struct {
	common.Common
}

func New(host, port, user, password, dbName string) *Mysql {
	MysqlInstance := Mysql{
		Common: common.Common{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			DBName:   dbName,
		},
	}

	MysqlInstance.Instance = &MysqlInstance
	return &MysqlInstance
}

// DSN data source name
func (c *Mysql) DSN() (driverName string, dataSource string) {
	driverName = DriverName
	dataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.DBName)
	return
}

func (c *Mysql) DBStruct() ([]common.Table, error) {
	// 	sqlStr := `SELECT TABLE_NAME tableName ,TABLE_COMMENT tableDesc
	// From INFORMATION_SCHEMA.TABLES
	// WHERE UPPER(table_type)='BASE TABLE'
	// AND LOWER(table_schema) = ?
	// ORDER BY table_name`
	// rows, err := c.DB.Query(sqlStr, c.DBName)

	sqlStr := `SHOW FULL TABLES`
	var result []common.Table
	rows, err := c.DB.Queryx(sqlStr)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	for rows.Next() {
		row, err := rows.SliceScan()
		if err != nil {
			fmt.Println(err)
			return result, err
		}

		result = append(result, common.Table{
			TableName: cast.ToString(row[0]),
			TableType: common.TTableType(cast.ToString(row[1])),
		})
	}

	return result, nil
}

// type Field struct {
// 	DatabaseName           string       `json:"database_name"`            // 数据库名称 TABLE_SCHEMA
// 	TableName              string       `json:"table_name"`               // 数据表名称 TABLE_NAME
// 	ColumnName             *string      `json:"column_name"`              // 字段名 COLUMN_NAME
// 	OrdinalPosition        int64        `json:"ordinal_position"`         // 排序序号 ORDINAL_POSITION
// 	ColumnDefault          *interface{} `json:"column_default"`           // 字段默认值 COLUMN_DEFAULT
// 	IsNullable             string       `json:"is_nullable"`              // 是否可以为 NULL IS_NULLABLE
// 	DataType               *string      `json:"data_type"`                // 字段数据类型 DATA_TYPE
// 	CharacterMaximumLength *int64       `json:"character_maximum_length"` // 字段的最大字符数 CHARACTER_MAXIMUM_LENGTH
// 	CharacterOctetLength   *int64       `json:"character_octet_length"`   // 字段的最大字节数 CHARACTER_OCTET_LENGTH
// 	NumericPrecision       *int64       `json:"numeric_precision"`        // 数字精度 NUMERIC_PRECISION
// 	NumericScale           *int64       `json:"numeric_scale"`            // 小数位数 NUMERIC_SCALE
// 	DateTimePrecision      *int64       `json:"date_time_precision"`      // 日期精度（datetime 类型和 SQL-92interval 类型数据库的子类型代码） DATETIME_PRECISION
// 	CharacterSetName       *string      `json:"character_set_name"`       // 字符集 CHARACTER_SET_NAME
// 	CollationName          *string      `json:"collation_name"`           // 字符集排序规则 COLLATION_NAME
// 	ColumnType             string       `json:"column_type"`              // 字段类型 COLUMN_TYPE
// 	ColumnKey              string       `json:"column_key"`               // 索引类型 COLUMN_KEY
// 	Extra                  *string      `json:"extra"`                    // 其它信息(auto_increment 等) EXTRA
// 	Privileges             *string      `json:"privileges"`               // 权限 PRIVILEGES
// 	ColumnComment          string       `json:"column_comment"`           // 字段注释 COLUMN_COMMENT
// }

// sqlStr := `SELECT TABLE_SCHEMA, TABLE_NAME, COLUMN_NAME, ORDINAL_POSITION, COLUMN_DEFAULT, IS_NULLABLE, DATA_TYPE,
// 	CHARACTER_MAXIMUM_LENGTH, CHARACTER_OCTET_LENGTH, NUMERIC_PRECISION, NUMERIC_SCALE, DATETIME_PRECISION, CHARACTER_SET_NAME,
//    	COLLATION_NAME, COLUMN_TYPE, COLUMN_KEY, EXTRA, PRIVILEGES, COLUMN_COMMENT
// 	FROM information_schema.columns
// 	WHERE table_schema = ? AND table_name = ?`

func (c *Mysql) TableStruct(tableName string) ([]common.Field, error) {
	sqlStr := `SHOW FULL COLUMNS FROM ` + tableName

	var result []common.Field
	err := c.DB.Select(&result, sqlStr)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	return result, nil
}
