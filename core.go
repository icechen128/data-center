package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Core struct {
	Type string `json:"type"` // 数据库类型
	db   *sql.DB
}

func (c *Core) Open() {
	db, err := sql.Open("mysql", "root:3541213@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	c.db = db
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
	fieldName string
	fieldDesc string
	dataType  string
	isNull    string
	length    int
}

func (c *Core) TableStruct(dbName, tableName string) []Field {
	sqlStr := `SELECT COLUMN_NAME fName,COLUMN_COMMENT fDesc,DATA_TYPE dataType,
    IS_NULLABLE isNull, IFNULL(CHARACTER_MAXIMUM_LENGTH, 0) sLength
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
		err = rows.Scan(&f.fieldName, &f.fieldDesc, &f.dataType, &f.isNull, &f.length)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		result = append(result, f)
	}

	return result
}
