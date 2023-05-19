package toolkit

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"path"
)

// 三维字符串数组转换为xlsx,错误则返回空路径和err
func Str3dToXlsx(str3d [][][]string, dir string, fileName string) (realativeFilePath string, err error) {
	//str3d 为空则返回
	if len(str3d) == 0 {
		return "", errors.New("str3d empty to xlsx")
	}

	//目录不存在则创建目录
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	//拼接目录、文件名
	realativeFilePath = path.Join(dir, fileName)
	if realativeFilePath == "" {
		return "", fmt.Errorf("empty realativeFilePath after join , dir=%s, fileName=%s", dir, fileName)
	}

	//遍历str3d写表
	f := xlsx.NewFile()
	for i := 0; i < len(str3d); i++ {
		if len(str3d[i]) == 0 { //str3d[i]为空，略过
			continue
		}
		sheet, err := f.AddSheet("sheet" + ToString(i))
		if err != nil {
			return "", err
		}
		for j := 0; j < len(str3d[i]); j++ {
			sheet.AddRow().WriteSlice(&(str3d[i][j]), -1)
		}
	}
	err = f.Save(realativeFilePath)
	if err != nil {
		return "", err
	}

	return realativeFilePath, nil

}
