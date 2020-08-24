package main

import (
	"testing"
)

func TestCore_Open(t *testing.T) {
	c := Core{}
	c.Open()
	result := c.DBStruct("test")
	t.Log(result)

	tables := make(map[string][]Field)
	for tableName := range result {
		fields := c.TableStruct("test", tableName)
		tables[tableName] = fields
	}
	t.Log(tables)
}

func BenchmarkCore_TableStruct(b *testing.B) {
	c := Core{}
	c.Open()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := c.DBStruct("test")

		tables := make(map[string][]Field)
		for tableName := range result {
			fields := c.TableStruct("test", tableName)
			tables[tableName] = fields
		}
	}
}
