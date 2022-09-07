package common

type Database interface {
	DBStruct()
	TableStruct(tableName string)

	Find(tableName string, page, pageSize int64, soryBy string, where ...string)
	First(tableName string, where ...string)
	Update(tableName string)
	Insert(tableName string, obj interface{})
	Delete(tableName string, obj interface{})
}
