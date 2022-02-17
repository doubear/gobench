package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var ELQ1 *gorm.DB

func Dbinit() {
	var err error
	ELQ1, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err != nil {
		fmt.Println("mysql connect fail")
	}
	if ELQ1.Error != nil {
		fmt.Println("database error ")
		fmt.Println(ELQ1.Error)
	}
}

func Dbclose() {
	defer ELQ1.Close()
}
