package toolkit

import (
	"fmt"
	"testing"
)

func TestXlsxToStr3d(t *testing.T) {
	path := "test_files/str3dTest.xlsx"
	str3d, err := XlsxToStr3d(path)
	if err != nil {
		t.Fatal("Error while TestXlsxToStr3d: ", err)
	} else {
		fmt.Println(str3d)
	}
}
