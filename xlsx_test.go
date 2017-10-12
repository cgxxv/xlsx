package xlsx

import (
	"strconv"
	"testing"
)

func Test_OpenFile(t *testing.T) {
	t.Run("test open xlsx file", func(t *testing.T) {
		filename := "./testxlsx/sample.xlsx"
		_, err := OpenFile(filename)
		if err != nil {
			t.Errorf("OpenFile() error = %v", err)
			return
		}
	})
}

func Test_GetSheetIndex(t *testing.T) {
	t.Run("test open xlsx file", func(t *testing.T) {
		filename := "./testxlsx/sample.xlsx"
		f, _ := OpenFile(filename)
		_, err := f.GetSheetIndex("Sheet1")
		if err != nil {
			t.Errorf("GetSheetIndex() error = %v", err)
			return
		}
	})
}

func TestFile_GetRows(t *testing.T) {
	filename := "./testxlsx/sample.xlsx"
	xlsx, _ := OpenFile(filename)
	index, _ := xlsx.GetSheetIndex("Sheet2")
	// Get all the rows in a sheet.
	_, err := xlsx.GetRows("sheet" + strconv.Itoa(index))
	if err != nil {
		t.Errorf("GetRows() error = %v", err)
		return
	}
}
