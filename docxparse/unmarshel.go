package docxparse

import (
	"fmt"

	"github.com/muktihari/xmltokenizer"
)

func (p *P) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	if se.SelfClosing {
		return nil
	}

	for {
		token, err := tok.Token()
		if err != nil {
			return err
		}
		if token.IsEndElementOf(se) { // Reach desired EndElement
			return nil
		}
		if token.IsEndElement { // Ignore child's EndElements
			continue
		}
		switch string(token.Name.Local) {

		case "r":
			r := new(R)
			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = r.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				return err
			}
			// fmt.Println("+r")
			p.R = append(p.R, r)
		case "pPr":
			ppr := new(PPr)
			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = ppr.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				return err
			}
			p.Ppr = ppr
		}
	}
}

func (ppr *PPr) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	if se.SelfClosing {
		return nil
	}

	for {
		token, err := tok.Token()
		if err != nil {
			return err
		}
		if token.IsEndElementOf(se) { // Reach desired EndElement
			return nil
		}
		if token.IsEndElement { // Ignore child's EndElements
			continue
		}
		switch string(token.Name.Local) {
		case "spacing":
			spacing := new(Spacing)
			se := xmltokenizer.GetToken().Copy(token)
			err = spacing.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				return err
			}
			ppr.Spacing = spacing
		}
	}

}

func (spacing *Spacing) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	for i := range se.Attrs {
		if string(se.Attrs[i].Name.Local) == "line" {
			spacing.Line = string(se.Attrs[i].Value)
		}
		if string(se.Attrs[i].Name.Local) == "lineRule" {
			spacing.LineRule = string(se.Attrs[i].Value)
		}
	}
	if se.SelfClosing {
		return nil
	}
	return nil
}

func (r *R) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	if se.SelfClosing {
		return nil
	}

	for {
		token, err := tok.Token()
		if err != nil {
			return err
		}
		if token.IsEndElementOf(se) { // Reach desired EndElement
			return nil
		}
		if token.IsEndElement { // Ignore child's EndElements
			continue
		}
		switch string(token.Name.Local) {
		case "t":
			r.Text = string(token.Data)
			// fmt.Println("+t", r.Text)
		case "rPr":
			attr := new(Attr)
			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = attr.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				return err
			}

			r.Attr = attr
			// fmt.Println("+attr", r.Attr)
		}
	}
}

func (attr *Attr) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	if se.SelfClosing {
		return nil
	}

	for {
		token, err := tok.Token()
		if err != nil {
			return err
		}
		if token.IsEndElementOf(se) { // Reach desired EndElement
			return nil
		}
		if token.IsEndElement { // Ignore child's EndElements
			continue
		}
		switch string(token.Name.Local) {
		case "rFonts":
			font := new(RFonts)

			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = font.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				fmt.Println("eee:", err)
				return err
			}

			attr.Fonts = font
		case "sz":
			size := new(Sz)
			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = size.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				return err
			}

			attr.Size = size
			// fmt.Println("+size", size)
		}
	}
}

func (rfonts *RFonts) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	// fmt.Println("rfont:", se.Attrs)
	for i := range se.Attrs {
		attr := &se.Attrs[i]
		// fmt.Println("rfont:", string(attr.Name.Local))
		switch string(attr.Name.Local) {
		case "ascii":
			rfonts.Ascii = string(attr.Value)
		case "eastAsia":
			rfonts.EastAsia = string(attr.Value)
		case "hAnsi":
			rfonts.HAnsi = string(attr.Value)
		}
	}
	// fmt.Println("+fonts", rfonts.EastAsia, rfonts.HAnsi, rfonts.Ascii)
	if se.SelfClosing {
		return nil
	}

	return nil
}

func (sz *Sz) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	for i := range se.Attrs {
		attr := &se.Attrs[i]
		switch string(attr.Name.Local) {
		case "val":
			sz.Val = string(attr.Value)
		}
	}
	if se.SelfClosing {
		return nil
	}

	// fmt.Println("+size : ", sz.Val)
	return nil
}

func (tbl *Tbl) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {
	if se.SelfClosing {
		return nil
	}

	for {
		token, err := tok.Token()
		if err != nil {
			return err
		}
		if token.IsEndElementOf(se) { // Reach desired EndElement
			return nil
		}
		if token.IsEndElement { // Ignore child's EndElements
			continue
		}
		switch string(token.Name.Local) {
		case "tr":
			tr := new(Tr)

			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = tr.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err == nil {

				tbl.Rows = append(tbl.Rows, tr)
			}
			// fmt.Println("+tr")

		}
	}

}

func (tr *Tr) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {

	if se.SelfClosing {
		return nil
	}

	for {
		token, err := tok.Token()
		if err != nil {
			return err
		}
		if token.IsEndElementOf(se) { // Reach desired EndElement
			return nil
		}
		if token.IsEndElement { // Ignore child's EndElements
			continue
		}
		switch string(token.Name.Local) {
		case "tc":
			tc := new(Tc)

			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = tc.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err == nil {
				// fmt.Println("tc:", tc.String())
				tr.Cells = append(tr.Cells, tc)
			}

		}
	}
}

func (tc *Tc) UnmarshalToken(tok *xmltokenizer.Tokenizer, se *xmltokenizer.Token) error {

	if se.SelfClosing {
		return nil
	}

	for {
		token, err := tok.Token()
		if err != nil {
			// fmt.Println("err 1:", err)
			return err
		}
		if token.IsEndElementOf(se) { // Reach desired EndElement
			return nil
		}
		if token.IsEndElement { // Ignore child's EndElements
			continue
		}
		switch string(token.Name.Local) {

		case "p":
			p := new(P)

			// Reuse Token object in the sync.Pool since we only use it temporarily.
			se := xmltokenizer.GetToken().Copy(token)
			err = p.UnmarshalToken(tok, se)
			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				fmt.Println("p err", err)
				return err
			}
			// fmt.Println("+r")
			tc.P = p
		}
	}
}
