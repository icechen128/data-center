package mysql

import (
	"github.com/icechen128/data-center/internal/pkg/database/common"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var defaultMysqlDB = New("127.0.0.1", "3306", "root", "3541213", "test")

func TestMysql_DBStruct(t *testing.T) {
	tests := []struct {
		name string
		db   *Mysql
		want []common.Table
	}{
		{
			name: "test get tables",
			db:   defaultMysqlDB,
			want: []common.Table{
				{
					TableName: "users",
					TableType: common.TableTypeBaseTable},
				{
					TableName: "usersView",
					TableType: common.TableTypeView,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.db
			err := c.Open()
			if err != nil {
				t.Error(err)
			}
			if got, _ := c.DBStruct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mysql.DBStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMysql_TableStruct(t *testing.T) {
	tests := []struct {
		name string
		db   *Mysql
		args struct{ tableName string }
		want map[string]string
	}{
		{
			name: "test get column",
			db:   defaultMysqlDB,
			args: struct{ tableName string }{tableName: "users"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.db
			err := c.Open()
			if err != nil {
				t.Error(err)
			}
			if got, _ := c.TableStruct(tt.args.tableName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mysql.TableStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
