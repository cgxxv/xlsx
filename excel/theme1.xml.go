package excel

type Theme struct {
	A              string         `xml:"a,attr"`
	R              string         `xml:"r,attr"`
	Name           string         `xml:"name,attr"`
	ThemeElements  ThemeElements  `xml:"themeElements"`
	ObjectDefaults ObjectDefaults `xml:"objectDefaults"`
}
type ThemeElements struct {
	ClrScheme  ClrScheme  `xml:"clrScheme"`
	FontScheme FontScheme `xml:"fontScheme"`
	FmtScheme  FmtScheme  `xml:"fmtScheme"`
}
type ObjectDefaults struct {
	SpDef SpDef `xml:"spDef"`
	LnDef LnDef `xml:"lnDef"`
	TxDef TxDef `xml:"txDef"`
}
type ClrScheme struct {
	Lt2      Lt2      `xml:"lt2"`
	Accent2  Accent2  `xml:"accent2"`
	Accent4  Accent4  `xml:"accent4"`
	Accent6  Accent6  `xml:"accent6"`
	Hlink    Hlink    `xml:"hlink"`
	Name     string   `xml:"name,attr"`
	Dk1      Dk1      `xml:"dk1"`
	Lt1      Lt1      `xml:"lt1"`
	Accent5  Accent5  `xml:"accent5"`
	FolHlink FolHlink `xml:"folHlink"`
	Dk2      Dk2      `xml:"dk2"`
	Accent1  Accent1  `xml:"accent1"`
	Accent3  Accent3  `xml:"accent3"`
}
type FontScheme struct {
	Name      string    `xml:"name,attr"`
	MajorFont MajorFont `xml:"majorFont"`
	MinorFont MinorFont `xml:"minorFont"`
}
type SpDef struct {
	Style    Style    `xml:"style"`
	SpPr     SpPr     `xml:"spPr"`
	BodyPr   BodyPr   `xml:"bodyPr"`
	LstStyle LstStyle `xml:"lstStyle"`
}
type Lt2 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type Lt1 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type Accent5 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type SrgbClr struct {
	Val string `xml:"val,attr"`
}
type TxDef struct {
	SpPr     SpPr     `xml:"spPr"`
	BodyPr   BodyPr   `xml:"bodyPr"`
	LstStyle LstStyle `xml:"lstStyle"`
	Style    Style    `xml:"style"`
}
type Accent2 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type Hlink struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type BodyPr struct {
	Vert             string `xml:"vert,attr"`
	LIns             string `xml:"lIns,attr"`
	RIns             string `xml:"rIns,attr"`
	BIns             string `xml:"bIns,attr"`
	TIns             string `xml:"tIns,attr"`
	SpcCol           string `xml:"spcCol,attr"`
	Rot              string `xml:"rot,attr"`
	SpcFirstLastPara string `xml:"spcFirstLastPara,attr"`
	VertOverflow     string `xml:"vertOverflow,attr"`
	HorzOverflow     string `xml:"horzOverflow,attr"`
	RtlCol           string `xml:"rtlCol,attr"`
	Anchor           string `xml:"anchor,attr"`
	Upright          string `xml:"upright,attr"`
	Wrap             string `xml:"wrap,attr"`
	NumCol           string `xml:"numCol,attr"`
}
type MajorFont struct {
	Latin Latin `xml:"latin"`
	Ea    Ea    `xml:"ea"`
	Cs    Cs    `xml:"cs"`
}
type MinorFont struct {
	Latin Latin `xml:"latin"`
	Ea    Ea    `xml:"ea"`
	Cs    Cs    `xml:"cs"`
}
type Style struct {
	LnRef     LnRef     `xml:"lnRef"`
	FillRef   FillRef   `xml:"fillRef"`
	EffectRef EffectRef `xml:"effectRef"`
	FontRef   FontRef   `xml:"fontRef"`
}
type SpPr struct {
	Ln        Ln        `xml:"ln"`
	SolidFill SolidFill `xml:"solidFill"`
}
type LnDef struct {
	BodyPr   BodyPr   `xml:"bodyPr"`
	LstStyle LstStyle `xml:"lstStyle"`
	Style    Style    `xml:"style"`
	SpPr     SpPr     `xml:"spPr"`
}
type Accent6 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type FolHlink struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type Accent3 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type LstStyle struct {
	DefPPr  DefPPr  `xml:"defPPr"`
	Lvl8pPr Lvl8pPr `xml:"lvl8pPr"`
	Lvl7pPr Lvl7pPr `xml:"lvl7pPr"`
	Lvl9pPr Lvl9pPr `xml:"lvl9pPr"`
	Lvl1pPr Lvl1pPr `xml:"lvl1pPr"`
	Lvl2pPr Lvl2pPr `xml:"lvl2pPr"`
	Lvl3pPr Lvl3pPr `xml:"lvl3pPr"`
	Lvl4pPr Lvl4pPr `xml:"lvl4pPr"`
	Lvl5pPr Lvl5pPr `xml:"lvl5pPr"`
	Lvl6pPr Lvl6pPr `xml:"lvl6pPr"`
}
type FmtScheme struct {
	FillStyleLst   FillStyleLst   `xml:"fillStyleLst"`
	LnStyleLst     LnStyleLst     `xml:"lnStyleLst"`
	BgFillStyleLst BgFillStyleLst `xml:"bgFillStyleLst"`
	Name           string         `xml:"name,attr"`
}
type Dk2 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type Accent4 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type Dk1 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type Accent1 struct {
	SrgbClr SrgbClr `xml:"srgbClr"`
}
type BgFillStyleLst struct {
	SolidFill SolidFill  `xml:"solidFill"`
	GradFill  []GradFill `xml:"gradFill"`
}
type GradFill struct {
	GsLst        GsLst  `xml:"gsLst"`
	Path         Path   `xml:"path"`
	RotWithShape string `xml:"rotWithShape,attr"`
	Lin          Lin    `xml:"lin"`
}
type Ea struct {
	Typeface string `xml:"typeface,attr"`
}
type Latin struct {
	Typeface string `xml:"typeface,attr"`
}
type SolidFill struct {
	SrgbClr   SrgbClr   `xml:"srgbClr"`
	SchemeClr SchemeClr `xml:"schemeClr"`
}
type Lvl9pPr struct {
	Algn         string `xml:"algn,attr"`
	DefRPr       DefRPr `xml:"defRPr"`
	MarL         string `xml:"marL,attr"`
	MarR         string `xml:"marR,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
	SpcBef       SpcBef `xml:"spcBef"`
	Indent       string `xml:"indent,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	Rtl          string `xml:"rtl,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	SpcAft       SpcAft `xml:"spcAft"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
}
type FillRef struct {
	Idx string `xml:"idx,attr"`
}
type SpcAft struct {
	SpcPts SpcPts `xml:"spcPts"`
}
type Lvl8pPr struct {
	MarR         string `xml:"marR,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
	SpcAft       SpcAft `xml:"spcAft"`
	MarL         string `xml:"marL,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
	Rtl          string `xml:"rtl,attr"`
	Algn         string `xml:"algn,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	Indent       string `xml:"indent,attr"`
	DefRPr       DefRPr `xml:"defRPr"`
}
type Lvl3pPr struct {
	MarL         string `xml:"marL,attr"`
	Algn         string `xml:"algn,attr"`
	Rtl          string `xml:"rtl,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	Indent       string `xml:"indent,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
	DefRPr       DefRPr `xml:"defRPr"`
	MarR         string `xml:"marR,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
	SpcAft       SpcAft `xml:"spcAft"`
}
type Lvl6pPr struct {
	MarR         string `xml:"marR,attr"`
	Indent       string `xml:"indent,attr"`
	Algn         string `xml:"algn,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
	SpcAft       SpcAft `xml:"spcAft"`
	DefRPr       DefRPr `xml:"defRPr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
	MarL         string `xml:"marL,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	Rtl          string `xml:"rtl,attr"`
}
type DefPPr struct {
	DefTabSz     string `xml:"defTabSz,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	DefRPr       DefRPr `xml:"defRPr"`
	Indent       string `xml:"indent,attr"`
	Algn         string `xml:"algn,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
	MarL         string `xml:"marL,attr"`
	MarR         string `xml:"marR,attr"`
	Rtl          string `xml:"rtl,attr"`
	SpcAft       SpcAft `xml:"spcAft"`
	HangingPunct string `xml:"hangingPunct,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
}
type FontRef struct {
	Idx string `xml:"idx,attr"`
}
type Cs struct {
	Typeface string `xml:"typeface,attr"`
}
type FillStyleLst struct {
	SolidFill SolidFill  `xml:"solidFill"`
	GradFill  []GradFill `xml:"gradFill"`
}
type Ln struct {
	W         string    `xml:"w,attr"`
	Cap       string    `xml:"cap,attr"`
	SolidFill SolidFill `xml:"solidFill"`
	PrstDash  PrstDash  `xml:"prstDash"`
	Miter     Miter     `xml:"miter"`
	Cmpd      string    `xml:"cmpd,attr"`
	Algn      string    `xml:"algn,attr"`
}
type Path struct {
	Path       string     `xml:"path,attr"`
	FillToRect FillToRect `xml:"fillToRect"`
}
type DefRPr struct {
	Cap        string    `xml:"cap,attr"`
	I          string    `xml:"i,attr"`
	Strike     string    `xml:"strike,attr"`
	Kumimoji   string    `xml:"kumimoji,attr"`
	B          string    `xml:"b,attr"`
	Baseline   string    `xml:"baseline,attr"`
	U          string    `xml:"u,attr"`
	NormalizeH string    `xml:"normalizeH,attr"`
	SolidFill  SolidFill `xml:"solidFill"`
	Spc        string    `xml:"spc,attr"`
	Sz         string    `xml:"sz,attr"`
	Latin      Latin     `xml:"latin"`
	Ea         Ea        `xml:"ea"`
	Sym        Sym       `xml:"sym"`
	Cs         Cs        `xml:"cs"`
}
type LnSpc struct {
	SpcPct SpcPct `xml:"spcPct"`
}
type EffectRef struct {
	Idx string `xml:"idx,attr"`
}
type SpcBef struct {
	SpcPts SpcPts `xml:"spcPts"`
}
type LnStyleLst struct {
	Ln []Ln `xml:"ln"`
}
type SpcPts struct {
	Val string `xml:"val,attr"`
}
type Lvl1pPr struct {
	LnSpc        LnSpc  `xml:"lnSpc"`
	SpcAft       SpcAft `xml:"spcAft"`
	MarL         string `xml:"marL,attr"`
	MarR         string `xml:"marR,attr"`
	Indent       string `xml:"indent,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
	DefRPr       DefRPr `xml:"defRPr"`
	Algn         string `xml:"algn,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	Rtl          string `xml:"rtl,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
}
type Lvl7pPr struct {
	HangingPunct string `xml:"hangingPunct,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
	MarL         string `xml:"marL,attr"`
	Indent       string `xml:"indent,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	SpcAft       SpcAft `xml:"spcAft"`
	DefRPr       DefRPr `xml:"defRPr"`
	MarR         string `xml:"marR,attr"`
	Rtl          string `xml:"rtl,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	Algn         string `xml:"algn,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
}
type Lvl2pPr struct {
	MarR         string `xml:"marR,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
	SpcAft       SpcAft `xml:"spcAft"`
	MarL         string `xml:"marL,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
	DefRPr       DefRPr `xml:"defRPr"`
	Indent       string `xml:"indent,attr"`
	Algn         string `xml:"algn,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	Rtl          string `xml:"rtl,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
}
type LnRef struct {
	Idx string `xml:"idx,attr"`
}
type Lvl4pPr struct {
	MarR         string `xml:"marR,attr"`
	Rtl          string `xml:"rtl,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
	Indent       string `xml:"indent,attr"`
	Algn         string `xml:"algn,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	SpcAft       SpcAft `xml:"spcAft"`
	DefRPr       DefRPr `xml:"defRPr"`
	MarL         string `xml:"marL,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
}
type GsLst struct {
	Gs []Gs `xml:"gs"`
}
type Lvl5pPr struct {
	MarL         string `xml:"marL,attr"`
	MarR         string `xml:"marR,attr"`
	DefTabSz     string `xml:"defTabSz,attr"`
	FontAlgn     string `xml:"fontAlgn,attr"`
	LnSpc        LnSpc  `xml:"lnSpc"`
	SpcAft       SpcAft `xml:"spcAft"`
	Indent       string `xml:"indent,attr"`
	Algn         string `xml:"algn,attr"`
	Rtl          string `xml:"rtl,attr"`
	LatinLnBrk   string `xml:"latinLnBrk,attr"`
	HangingPunct string `xml:"hangingPunct,attr"`
	SpcBef       SpcBef `xml:"spcBef"`
	DefRPr       DefRPr `xml:"defRPr"`
}
type Gs struct {
	Pos       string    `xml:"pos,attr"`
	SchemeClr SchemeClr `xml:"schemeClr"`
}
type SpcPct struct {
	Val string `xml:"val,attr"`
}
type Miter struct {
	Lim string `xml:"lim,attr"`
}
type Sym struct {
	Typeface string `xml:"typeface,attr"`
}
type PrstDash struct {
	Val string `xml:"val,attr"`
}
type FillToRect struct {
	L string `xml:"l,attr"`
	T string `xml:"t,attr"`
	R string `xml:"r,attr"`
	B string `xml:"b,attr"`
}
type SchemeClr struct {
	Val    string `xml:"val,attr"`
	Tint   Tint   `xml:"tint"`
	SatMod SatMod `xml:"satMod"`
	Shade  Shade  `xml:"shade"`
}
type SatMod struct {
	Val string `xml:"val,attr"`
}
type Shade struct {
	Val string `xml:"val,attr"`
}
type Lin struct {
	Ang    string `xml:"ang,attr"`
	Scaled string `xml:"scaled,attr"`
}
type Tint struct {
	Val string `xml:"val,attr"`
}
