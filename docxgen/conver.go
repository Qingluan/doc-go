package docxgen

import (
	"fmt"
	"strings"
)

var (
	SEP_XML    = `<p w:rsidR="002423EF" w:rsidRDefault="002423EF"><pPr><spacing w:line="560" w:lineRule="exact" /><rPr><rFonts w:ascii="Times New Roman" w:eastAsia="方正仿宋简体" w:hAnsi="Times New Roman" /><sz w:val="32" /><szCs w:val="32" /></rPr></pPr></p>`
	R_XML_TEMP = `
	<r>
		<rPr>
			<rFonts w:ascii="{ASCII}" w:eastAsia="{EASTASIA}" w:hAnsi="{HANSI}" w:hint="eastAsia" />
			<sz w:val="{SIZE}" />
			<szCs w:val="{SIZE}"/>
		</rPr>
		<t>{TEXT}</t>
	</r>`
	PRP_TITLE = `<pPr>
	<spacing w:line="560" w:lineRule="exact" />
		<jc w:val="center" />
		<rPr>
			<rFonts w:ascii="Times New Roman" w:eastAsia="方正小标宋简体" w:hAnsi="Times New Roman" />
			<sz w:val="44" />
			<szCs w:val="32" />
		</rPr>
	</pPr>`
	PRP_H1 = `
	<pPr>
		<numPr>
			<ilvl w:val="0" />
			<numId w:val="1" />
		</numPr>
		<spacing w:line="560" w:lineRule="exact" />
		<ind w:firstLineChars="200" w:firstLine="720" />
		<rPr>
			<rFonts w:ascii="Times New Roman" w:eastAsia="黑体" w:hAnsi="Times New Roman" />
			<sz w:val="36" />
			<szCs w:val="36" />
			<lang w:val="zh-CN" />
		</rPr>
	</pPr>
	`
	PRP_H2 = `
	<pPr>
		<pStyle w:val="a5" />
		<widowControl />
		<numPr>
			<ilvl w:val="0" />
			<numId w:val="2" />
		</numPr>
		<shd w:val="clear" w:color="auto" w:fill="FFFFFF" />
		<spacing w:before="105" w:beforeAutospacing="0" w:after="105" w:afterAutospacing="0" w:line="560" w:lineRule="exact" />
		<ind w:firstLineChars="200" w:firstLine="720" />
		<rPr>
			<rFonts w:ascii="Times New Roman" w:eastAsia="楷体_GB2312" w:hAnsi="Times New Roman" />
			<sz w:val="36" />
			<szCs w:val="36" />
		</rPr>
	</pPr>
	<proofErr w:type="gramStart" />
	`
	PRP_NORM = `
	<pPr>
		<ind w:firstLineChars="200" w:firstLine="640" />
		<rPr>
			<rFonts w:ascii="Times New Roman" w:eastAsia="方正仿宋简体" w:hAnsi="Times New Roman" />
			<sz w:val="32" />
			<szCs w:val="32" />
		</rPr>
	</pPr>
	`
	END_XML = `<p w:rsidR="002423EF" w:rsidRDefault="002423EF">
		<pPr>
			<topLinePunct />
			<autoSpaceDE w:val="0" />
			<autoSpaceDN w:val="0" />
			<adjustRightInd w:val="0" />
			<spacing w:line="560" w:lineRule="exact" />
			<ind w:firstLineChars="200" w:firstLine="420" />
			<jc w:val="left" />
			<rPr>
				<rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" />
			</rPr>
		</pPr>
	</p>
	<sectPr w:rsidR="002423EF">
		<pgSz w:w="11906" w:h="16838" />
		<pgMar w:top="2098" w:right="1474" w:bottom="1984" w:left="1587" w:header="851" w:footer="992" w:gutter="0" />
		<cols w:space="0" />
		<docGrid w:type="lines" w:linePitch="312" />
	</sectPr>`
)

func (p *DocxParagraph) Xml() string {
	/*
		<r>
			<rPr>
				<rFonts w:ascii="{ascii}" w:eastAsia="{eastAsia}" w:hAnsi="{hAnsi}" />
				<sz w:val="{size}" />
				<szCs w:val="{size}"/>
			</rPr>
			<t>
				{text}
			</t>
		</r>
	*/
	text := strings.TrimSpace(p.Text)
	// if text == "" {
	// 	return ""
	// }
	root := `<p w:rsidR="002423EF" w:rsidRDefault="0015648D">`
	switch p.Type {
	case TYPE_TITLE:
		root += PRP_TITLE
	case TYPE_H1:
		root += PRP_H1
	case TYPE_H2:
		root += PRP_H2
	case TYPE_SEP:
		return SEP_XML
	default:
		root += PRP_NORM
	}

	if p.Type == TYPE_H1 {
		// startString := text[:3]
		// leftString := text[3 : len(text)-3]
		// endString := text[len(text)-3:]
		// root += Wrap(startString, p.Font.Ascii, p.Font.EastAsia, p.Font.HAnsi, fmt.Sprint(p.Size))
		root += `<proofErr w:type="gramStart" />`
		root += Wrap(text, p.Font.Ascii, p.Font.EastAsia, p.Font.HAnsi, fmt.Sprint(p.Size))
		// root += Wrap(leftString, p.Font.Ascii, p.Font.EastAsia, p.Font.HAnsi, fmt.Sprint(p.Size))
		// root += Wrap(endString, p.Font.Ascii, p.Font.EastAsia, p.Font.HAnsi, fmt.Sprint(p.Size))
		root += `<proofErr w:type="gramEnd" />`
		root += "</p>"
		// fmt.Println(startString)
		// fmt.Println(leftString)
		// fmt.Println(endString)
	} else {
		root += Wrap(text, p.Font.Ascii, p.Font.EastAsia, p.Font.HAnsi, fmt.Sprint(p.Size))
		root += "</p>"

	}

	// fmt.Println(text)
	return root
}

func Wrap(text, ascii, eastAsia, hansi string, size string) string {
	o := strings.ReplaceAll(R_XML_TEMP, "{ASCII}", ascii)
	o = strings.ReplaceAll(o, "{EASTASIA}", eastAsia)
	o = strings.ReplaceAll(o, "{HANSI}", hansi)
	o = strings.ReplaceAll(o, "{SIZE}", size)
	o = strings.ReplaceAll(o, "{TEXT}", text)
	return o
}

func (c *DocxCreator) Xml() string {
	root := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<document xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup" xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk" xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml" xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" mc:Ignorable="w14 wp14">
	<body>`

	for _, p := range c.Paragraphs {
		root += p.Xml()
	}
	root += END_XML
	root += "</body></document>"
	return root
}

func (dc *DocxCreator) FromMarkdown(markdown string) *DocxCreator {
	lines := strings.Split(markdown, "\n")
	dc.Paragraphs = []*DocxParagraph{}
	for _, line := range lines {
		if strings.HasPrefix(line, "### ") {
			// fmt.Println("\t\tadd h3")
			dc.AddParagraph(line[4:])
			dc.LastPragraph().AsH2()
		} else if strings.HasPrefix(line, "## ") {
			// fmt.Println("\tadd h2")
			// dc.AddSepParagraph()
			// dc.AddParagraph("")
			dc.AddParagraph(line[3:])
			dc.LastPragraph().AsH1()
		} else if strings.HasPrefix(line, "# ") {
			// fmt.Println("add h1")
			dc.AddParagraph(line[2:])
			dc.LastPragraph().AsTitle()
			dc.AddSepParagraph()
		} else if line == "" {
			dc.AddSepParagraph()
			// fmt.Println("---- add sep")
		} else {
			// fmt.Println("\t\t\tadd normal")
			dc.AddParagraph(line)
		}
	}
	return dc
}
