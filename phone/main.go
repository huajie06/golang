package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var phone = `1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892`

	for i, v := range strings.Split(phone, "\n") {
		fmt.Printf("#%v: %v\n", i, numOnly(v))
	}

	var err error
	db, err := sql.Open("mysql", "hzhang:12345@tcp(192.168.1.137:3306)/stock")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func numOnly(s string) string {
	var ret string
	for _, v := range s {
		if v >= '0' && v <= '9' {
			ret += string(v)
		}
	}
	return ret
}
