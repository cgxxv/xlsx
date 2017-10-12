package excel

type Workbook struct {
	R         string    `xml:"r,attr"`
	Xmlns     string    `xml:"xmlns,attr"`
	BookViews BookViews `xml:"bookViews"`
	Sheets    Sheets    `xml:"sheets"`
}
type BookViews struct {
	WorkbookView WorkbookView `xml:"workbookView"`
}
type Sheets struct {
	Sheet []Sheet `xml:"sheet"`
}
type WorkbookView struct {
	XWindow      string `xml:"xWindow,attr"`
	YWindow      string `xml:"yWindow,attr"`
	WindowWidth  string `xml:"windowWidth,attr"`
	WindowHeight string `xml:"windowHeight,attr"`
}
type Sheet struct {
	Name    string `xml:"name,attr"`
	SheetId string `xml:"sheetId,attr"`
	Id      string `xml:"id,attr"`
}
