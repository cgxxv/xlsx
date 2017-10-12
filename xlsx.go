package xlsx

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"path"
	"strconv"
	"strings"

	"github.com/stackerzzq/xj2go"
	"github.com/stackerzzq/xlsx/excel"
)

type File struct {
	checked       map[string]bool
	Path          string
	SharedStrings *excel.Sst
	Sheet         map[string]*excel.Worksheet
	SheetCount    int
	Styles        *excel.StyleSheet
	WorkBook      *excel.Workbook
	WorkBookRels  *excel.Relationships
	XLSX          map[string][]byte
}

func OpenFile(filename string) (*File, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}

	fl, sc, err := readZip(zr)
	if err != nil {
		return nil, err
	}

	workSheets := 0
	for k, v := range fl {
		if path.Base(k)[:1] == "." {
			continue
		}

		if len(k) > 18 && k[0:19] == "xl/worksheets/sheet" {
			if workSheets == 0 {
				xj2go.BytesToGo("sheet.xml", "excel", &v)
			}
			workSheets++
			continue
		}

		filename := path.Base(k)
		xj2go.BytesToGo(filename, "excel", &v)
	}

	return &File{
		checked:    make(map[string]bool),
		Sheet:      make(map[string]*excel.Worksheet),
		SheetCount: sc,
		XLSX:       fl,
	}, err
}

func (f *File) GetSheetIndex(name string) (int, error) {
	content := f.workbookReader()
	rels := f.workbookRelsReader()
	for _, v := range content.Sheets.Sheet {
		if v.Name == name {
			for _, rel := range rels.Relationship {
				if v.Id == rel.Id {
					rID, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(rel.Target, "worksheets/sheet"), ".xml"))
					if err != nil {
						return 0, err
					}
					return rID, nil
				}
			}
		}
	}
	return 0, nil
}

func (f *File) GetRows(sheet string) ([][]string, error) {
	rows := [][]string{}
	d := f.sharedStringsReader()
	var r excel.Row
	var cr int
	name := "xl/worksheets/" + strings.ToLower(sheet) + ".xml"
	decoder := xml.NewDecoder(bytes.NewReader(f.readXML(name)))
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch element := t.(type) {
		case xml.StartElement:
			if element.Name.Local == "row" {
				r = excel.Row{}
				decoder.DecodeElement(&r, &element)
				rows = append(rows, []string{})
				for _, colCell := range r.C {
					val, err := f.getValueFrom(d, &colCell)
					if err != nil {
						return nil, err
					}
					cr = len(rows) - 1
					rows[cr] = append(rows[cr], val)
				}
			}
		default:
		}
	}

	return rows, nil
}
