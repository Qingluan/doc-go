package docxparse

import (
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type R struct {
	Text string `xml:"t,omitempty"`
	Attr *Attr  `xml:"rPr"`
}

type PPr struct {
	Spacing *Spacing `xml:"w:spacing,omitempty"`
}

type Spacing struct {
	Line     string `xml:"w:line,attr"`
	LineRule string `xml:"w:lineRule,attr"`
}

type P struct {
	Ppr *PPr `xml:"pPr"`
	R   []*R `xml:"r"`
}

func (p *P) IsTitle() bool {
	font := p.GetFont()
	if font != nil {
		if p.GetSize() == 44 && font.EastAsia == "方正小标宋简体" {
			return true
		}
	}
	return false
}
func (p *P) IsH1() bool {
	font := p.GetFont()
	if font != nil {
		if p.GetSize() == 36 && font.EastAsia == "黑体" {
			return true
		}
	}
	return false
}

func (p *P) IsH2() bool {
	font := p.GetFont()
	if font != nil {
		if p.GetSize() == 36 && font.EastAsia == "楷体_GB2312" {
			return true
		}
	}
	return false
}

func (p *P) GetSize() int {
	if len(p.R) > 0 {
		r := p.R[0]
		if r.Attr != nil {
			attr := r.Attr
			size := attr.Size
			if size != nil {
				ss, err := strconv.Atoi(size.Val)
				//fmt.Println(size)
				if err == nil {
					return ss
				}
			}

		}
	}
	return 0
}

func (p *P) GetFont() *RFonts {
	if len(p.R) > 0 {
		r := p.R[0]
		if r.Attr != nil {
			attr := r.Attr
			font := attr.Fonts
			return font
		}
	}
	return nil
}

func (p *P) Text() string {
	paras := ""
	if p.Ppr != nil {
		if p.Ppr.Spacing != nil {
			paras += "\n"
		}
	}

	if p.IsTitle() {
		paras += "# "
	} else if p.IsH1() {
		paras += "## "
	} else if p.IsH2() {
		paras += "### "
	}
	if strings.HasSuffix(paras, "\n") || paras == "" {
		paras += "  "
	}
	for _, r := range p.R {
		paras += r.Text
	}
	return paras
}

type Attr struct {
	Fonts *RFonts `xml:"rFonts"`
	Size  *Sz     `xml:"sz,omitempty"`
}

type Sz struct {
	Val string `xml:"w:val,attr"`
}

type RFonts struct {
	Ascii    string `xml:"w:ascii,attr"`
	HAnsi    string `xml:"w:hAnsi,attr"`
	EastAsia string `xml:"w:eastAsia,attr"`
}

type Tbl struct {
	Rows []*Tr `xml:"tr"`
}

func (tbl *Tbl) String() string {
	t := table.NewWriter()
	for _, tr := range tbl.Rows {
		ro := table.Row{}
		for _, cell := range tr.Strings() {
			ro = append(ro, cell)
		}
		t.AppendRow(ro)
	}
	return t.Render()
}

type Tr struct {
	Cells []*Tc `xml:"tc"`
}

func (tr *Tr) Strings() []string {
	var strings []string
	for _, cell := range tr.Cells {
		strings = append(strings, cell.String())
	}
	return strings
}

type Tc struct {
	P *P `xml:"p"`
}

func (tc *Tc) String() string {
	if tc.P != nil {
		return tc.P.Text()
	} else {
		return ""
	}

}
