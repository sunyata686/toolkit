package toolkit

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

// SqlToStr2d 执行sql，导出字符串二维数组,withHeader为true则str2d[1]为sql结果的列名
// dsn := "user:pwd@tcp(192.168.1.173:3306)/dbName?charset=utf8&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn))
func SqlToStr2d(Db *gorm.DB, sql string, withHeader bool) (str2dP *[][]string, resErr error) {
	var str2d [][]string
	str2dP = &str2d
	//执行sql
	rows, resErr := Db.Raw(sql).Rows()
	if resErr != nil {
		log.Println(resErr)
		return str2dP, resErr
	}

	//获取sql的结果的列名
	headerArr, err := rows.Columns()
	if err != nil {
		log.Println(resErr)
		return str2dP, err
	}

	//sql的列名插入表头
	if withHeader {
		str2d = [][]string{headerArr}
	}

	//----------------开始处理sql结果-----------------

	//声明逐行接受sql结果的【】interface{}，及其对应的指针数组
	scanRes := make([]interface{}, len(headerArr))
	scanResP := make([]interface{}, len(headerArr))
	for k := range scanRes {
		scanResP[k] = &scanRes[k]
	}
	//逐行处理
	for rows.Next() {
		strArr := make([]string, len(headerArr))
		err := rows.Scan(scanResP...)
		if err != nil {
			fmt.Println(err)
		}
		for i, v := range scanRes {
			strArr[i] = ToString(v)
		}
		str2d = append(str2d, strArr)
	}

	return str2dP, nil
}
