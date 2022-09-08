package temp

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetData(tableName string) error {
	db, err := sql.Open("mysql", "root:3541213@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	r, err := db.Query("select * from db")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(r)
	r.Err()
	return err
}
