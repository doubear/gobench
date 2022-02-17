package db

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var ELQ1 *gorm.DB

func Test_Dbinit(t *testing.T) {
	var err error
	ELQ1, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err != nil {
		fmt.Println("mysql connect fail")
	}
	if ELQ1.Error != nil {
		fmt.Println("database error ")
		fmt.Println(ELQ1.Error)
	}
	t.Log("hello test")
	defer ELQ1.Close()
}
