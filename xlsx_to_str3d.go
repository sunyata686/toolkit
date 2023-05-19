package toolkit

import "github.com/tealeg/xlsx"

func XlsxToStr3d(path string) ([][][]string, error) {
	xFile, err := xlsx.OpenFile(path)
	if err != nil {
		return [][][]string{}, err
	}

	out, err := xFile.ToSlice()
	if err != nil {
		return [][][]string{}, err
	}
	return out, nil
}
