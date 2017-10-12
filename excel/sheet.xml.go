package excel

type Worksheet struct {
	R             string        `xml:"r,attr"`
	Cols          Cols          `xml:"cols"`
	SheetData     SheetData     `xml:"sheetData"`
	HeaderFooter  HeaderFooter  `xml:"headerFooter"`
	Xmlns         string        `xml:"xmlns,attr"`
	Dimension     Dimension     `xml:"dimension"`
	SheetViews    SheetViews    `xml:"sheetViews"`
	SheetFormatPr SheetFormatPr `xml:"sheetFormatPr"`
	PageMargins   PageMargins   `xml:"pageMargins"`
	PageSetup     PageSetup     `xml:"pageSetup"`
}
type SheetData struct {
	Row []Row `xml:"row"`
}
type PageMargins struct {
	Right  string `xml:"right,attr"`
	Top    string `xml:"top,attr"`
	Bottom string `xml:"bottom,attr"`
	Header string `xml:"header,attr"`
	Footer string `xml:"footer,attr"`
	Left   string `xml:"left,attr"`
}
type Row struct {
	Ht           string `xml:"ht,attr"`
	CustomHeight string `xml:"customHeight,attr"`
	C            []C    `xml:"c"`
	R            string `xml:"r,attr"`
}
type Cols struct {
	Col []Col `xml:"col"`
}
type HeaderFooter struct {
	OddFooter string `xml:"oddFooter"`
}
type Dimension struct {
	Ref string `xml:"ref,attr"`
}
type SheetViews struct {
	SheetView SheetView `xml:"sheetView"`
}
type SheetFormatPr struct {
	DefaultColWidth  string `xml:"defaultColWidth,attr"`
	DefaultRowHeight string `xml:"defaultRowHeight,attr"`
	CustomHeight     string `xml:"customHeight,attr"`
	OutlineLevelRow  string `xml:"outlineLevelRow,attr"`
	OutlineLevelCol  string `xml:"outlineLevelCol,attr"`
}
type PageSetup struct {
	FitToWidth         string `xml:"fitToWidth,attr"`
	Scale              string `xml:"scale,attr"`
	UseFirstPageNumber string `xml:"useFirstPageNumber,attr"`
	Orientation        string `xml:"orientation,attr"`
	PageOrder          string `xml:"pageOrder,attr"`
	FirstPageNumber    string `xml:"firstPageNumber,attr"`
	FitToHeight        string `xml:"fitToHeight,attr"`
}
type SheetView struct {
	WorkbookViewId   string `xml:"workbookViewId,attr"`
	ShowGridLines    string `xml:"showGridLines,attr"`
	DefaultGridColor string `xml:"defaultGridColor,attr"`
}
type C struct {
	S string `xml:"s,attr"`
	V string `xml:"v"`
	R string `xml:"r,attr"`
	T string `xml:"t,attr"`
}
type Col struct {
	Min         string `xml:"min,attr"`
	Max         string `xml:"max,attr"`
	Width       string `xml:"width,attr"`
	Style       string `xml:"style,attr"`
	CustomWidth string `xml:"customWidth,attr"`
}
