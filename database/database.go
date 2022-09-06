package database

type Database interface {
	DBStruct(dbName string)
	TableStruct(dbName, tableName string)

	Find(tableName string, page, pageSize int64, soryBy string, where ...string)
	First(tableName string, where ...string)
	Update(tableName string)
	Insert(tableName string, obj interface{})
	Delete(tableName string, obj interface{})
}
