package xlsx

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"strconv"

	"github.com/stackerzzq/xlsx/excel"
)

func readZip(r *zip.Reader) (map[string][]byte, int, error) {
	fl := make(map[string][]byte)
	wsc := 0
	for _, v := range r.File {
		fl[v.Name] = readFile(v)
		if len(v.Name) > 18 {
			if v.Name[0:19] == "xl/worksheets/sheet" {
				wsc++
			}
		}
	}
	return fl, wsc, nil
}

func readFile(file *zip.File) []byte {
	rc, err := file.Open()
	if err != nil {
		panic(err)
	}
	buff := bytes.NewBuffer(nil)
	io.Copy(buff, rc)
	rc.Close()
	return buff.Bytes()
}

func (f *File) readXML(name string) []byte {
	if content, ok := f.XLSX[name]; ok {
		return content
	}
	return nil
}

func (f *File) workbookReader() *excel.Workbook {
	if f.WorkBook == nil {
		var content excel.Workbook
		xml.Unmarshal(f.readXML("xl/workbook.xml"), &content)
		f.WorkBook = &content
	}
	return f.WorkBook
}

func (f *File) workbookRelsReader() *excel.Relationships {
	if f.WorkBookRels == nil {
		var content excel.Relationships
		xml.Unmarshal(f.readXML("xl/_rels/workbook.xml.rels"), &content)
		f.WorkBookRels = &content
	}
	return f.WorkBookRels
}

func (f *File) sharedStringsReader() *excel.Sst {
	if f.SharedStrings == nil {
		var sharedStrings excel.Sst
		xml.Unmarshal(f.readXML("xl/sharedStrings.xml"), &sharedStrings)
		f.SharedStrings = &sharedStrings
	}
	return f.SharedStrings
}

func (f *File) stylesReader() *excel.StyleSheet {
	if f.Styles == nil {
		var styleSheet excel.StyleSheet
		xml.Unmarshal([]byte(f.readXML("xl/styles.xml")), &styleSheet)
		f.Styles = &styleSheet
	}
	return f.Styles
}

func (f *File) getValueFrom(d *excel.Sst, xlsx *excel.C) (string, error) {
	switch xlsx.T {
	case "s":
		xlsxSI, err := strconv.Atoi(xlsx.V)
		if err != nil {
			return "", err
		}
		return d.Si[xlsxSI].T, nil
	default:
		return xlsx.V, nil
	}
}
