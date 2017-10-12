package excel

type Types struct {
	Xmlns    string     `xml:"xmlns,attr"`
	Default  []Default  `xml:"Default"`
	Override []Override `xml:"Override"`
}
type Default struct {
	Extension   string `xml:"Extension,attr"`
	ContentType string `xml:"ContentType,attr"`
}
type Override struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}
