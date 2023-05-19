package toolkit

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestSqlToStr2d(t *testing.T) {
	dsn := "user:pwd@tcp(10.9.1.1:3306)/dbName?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		t.Fatal(err)
	}
	sql := `****`
	str2dP, err := SqlToStr2d(db, sql, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range *str2dP {
		fmt.Println(v)
	}
}
