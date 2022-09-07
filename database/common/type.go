package common

type Database interface {
	DBStruct() map[string]string
	TableStruct(tableName string) []Field

	//Find(tableName string, page, pageSize int64, soryBy string, where ...string)
	//First(tableName string, where ...string)
	//Update(tableName string)
	//Insert(tableName string, obj interface{})
	//Delete(tableName string, obj interface{})
}

type Field struct {
	Field      string      `json:"Field" db:"Field"`
	Type       string      `json:"Type" db:"Type"`
	Collation  *string     `json:"Collation" db:"Collation"`
	Null       string      `json:"Null" db:"Null"`
	Key        string      `json:"Key" db:"Key"`
	Default    interface{} `json:"Default" db:"Default"`
	Extra      string      `json:"Extra" db:"Extra"`
	Privileges string      `json:"Privileges" db:"Privileges"`
	Comment    string      `json:"Comment" db:"Comment"`
}
