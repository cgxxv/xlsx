package excel

type Sst struct {
	UniqueCount string `xml:"uniqueCount,attr"`
	Xmlns       string `xml:"xmlns,attr"`
	Si          []Si   `xml:"si"`
}
type Si struct {
	T string `xml:"t"`
}
