package database

import (
	"encoding/json"
	"testing"
)

func TestCore_Open(t *testing.T) {
	c := NewDatabase("mysql", "localhost", "3306", "root", "3541213", "test")
	c.Open()
	defer c.Close()
	result := c.DBStruct("test OR 1=1")
	t.Log(result)

	tables := make(map[string][]Field)
	for tableName := range result {
		fields := c.TableStruct("test", tableName)
		tables[tableName] = fields
	}
	t.Logf("%+v", tables)
	jsonData, _ := json.Marshal(tables)
	t.Log(string(jsonData))
}

func BenchmarkCore_TableStruct(b *testing.B) {
	c := Core{}
	c.Open()
	defer c.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := c.DBStruct("test")
		b.Log(result)

		tables := make(map[string][]Field)
		for tableName := range result {
			fields := c.TableStruct("test", tableName)
			tables[tableName] = fields
		}
		b.Log(tables)
	}
}
