package excel

type StyleSheet struct {
	TableStyles  TableStyles  `xml:"tableStyles"`
	Colors       Colors       `xml:"colors"`
	NumFmts      NumFmts      `xml:"numFmts"`
	Fonts        Fonts        `xml:"fonts"`
	Fills        Fills        `xml:"fills"`
	CellXfs      CellXfs      `xml:"cellXfs"`
	CellStyles   CellStyles   `xml:"cellStyles"`
	Dxfs         Dxfs         `xml:"dxfs"`
	Xmlns        string       `xml:"xmlns,attr"`
	Borders      Borders      `xml:"borders"`
	CellStyleXfs CellStyleXfs `xml:"cellStyleXfs"`
}
type TableStyles struct {
	Count string `xml:"count,attr"`
}
type Fonts struct {
	Count string `xml:"count,attr"`
	Font  []Font `xml:"font"`
}
type Fills struct {
	Count string `xml:"count,attr"`
	Fill  []Fill `xml:"fill"`
}
type CellXfs struct {
	Xf    []Xf   `xml:"xf"`
	Count string `xml:"count,attr"`
}
type CellStyles struct {
	Count     string    `xml:"count,attr"`
	CellStyle CellStyle `xml:"cellStyle"`
}
type Dxfs struct {
	Count string `xml:"count,attr"`
}
type Borders struct {
	Count  string   `xml:"count,attr"`
	Border []Border `xml:"border"`
}
type Colors struct {
	IndexedColors IndexedColors `xml:"indexedColors"`
}
type NumFmts struct {
	Count  string `xml:"count,attr"`
	NumFmt NumFmt `xml:"numFmt"`
}
type CellStyleXfs struct {
	Count string `xml:"count,attr"`
	Xf    Xf     `xml:"xf"`
}
type Xf struct {
	ApplyFill         string    `xml:"applyFill,attr"`
	ApplyBorder       string    `xml:"applyBorder,attr"`
	ApplyAlignment    string    `xml:"applyAlignment,attr"`
	Alignment         Alignment `xml:"alignment"`
	ApplyFont         string    `xml:"applyFont,attr"`
	FontId            string    `xml:"fontId,attr"`
	ApplyNumberFormat string    `xml:"applyNumberFormat,attr"`
	ApplyProtection   string    `xml:"applyProtection,attr"`
	NumFmtId          string    `xml:"numFmtId,attr"`
	FillId            string    `xml:"fillId,attr"`
	BorderId          string    `xml:"borderId,attr"`
}
type Font struct {
	Sz    Sz    `xml:"sz"`
	Color Color `xml:"color"`
	Name  Name  `xml:"name"`
}
type Fill struct {
	PatternFill PatternFill `xml:"patternFill"`
}
type Border struct {
	Left   Left   `xml:"left"`
	Right  Right  `xml:"right"`
	Top    Top    `xml:"top"`
	Bottom Bottom `xml:"bottom"`
}
type IndexedColors struct {
	RgbColor []RgbColor `xml:"rgbColor"`
}
type NumFmt struct {
	NumFmtId   string `xml:"numFmtId,attr"`
	FormatCode string `xml:"formatCode,attr"`
}
type RgbColor struct {
	Rgb string `xml:"rgb,attr"`
}
type Alignment struct {
	Vertical string `xml:"vertical,attr"`
}
type Color struct {
	Indexed string `xml:"indexed,attr"`
}
type Top struct {
	Style string `xml:"style,attr"`
	Color Color  `xml:"color"`
}
type Left struct {
	Style string `xml:"style,attr"`
	Color Color  `xml:"color"`
}
type Bottom struct {
	Style string `xml:"style,attr"`
	Color Color  `xml:"color"`
}
type CellStyle struct {
	Name      string `xml:"name,attr"`
	XfId      string `xml:"xfId,attr"`
	BuiltinId string `xml:"builtinId,attr"`
}
type Sz struct {
	Val string `xml:"val,attr"`
}
type PatternFill struct {
	PatternType string  `xml:"patternType,attr"`
	BgColor     BgColor `xml:"bgColor"`
	FgColor     FgColor `xml:"fgColor"`
}
type BgColor struct {
	Auto string `xml:"auto,attr"`
}
type Right struct {
	Style string `xml:"style,attr"`
	Color Color  `xml:"color"`
}
type Name struct {
	Val string `xml:"val,attr"`
}
type FgColor struct {
	Indexed string `xml:"indexed,attr"`
}
