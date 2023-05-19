package toolkit

import (
	"fmt"
	"testing"
	"time"
)

func TestStr3dToXlsx(t *testing.T) {
	header := []string{"header", "header", "header", "header", "header"}
	row1 := []string{"row1", "row1", "row1"}
	row2 := []string{"row2", "row2", "row2", "row2"}
	var str3d = make([][][]string, 3, 3)
	for i := 0; i < len(str3d); i++ {
		str3d[i] = [][]string{header, row1, row2}
	}
	str3d[0] = append(str3d[0], header)
	path, err := Str3dToXlsx(str3d, "./cache", time.Now().Format("str3dtest.150405.xlsx"))
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("	File saved at ", path)
	}
}
