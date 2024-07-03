package docxparse

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/muktihari/xmltokenizer"
)

func ReadDocx(filename string) (paragraphs []string) {
	xmlName, err := GetDocxRawText(filename)

	if err != nil {
		log.Println(err)
		return
	}

	/*
		text example is :
			<p>
				<r>
					<t>Hello</t>
				</r>
			</p>



		table example is :
			<tbl>
					<tblPr>
					<tblW w="0" type="auto"/>
					<tblLook val="04E0" firstRow="1" lastRow="0" firstColumn="1" lastColumn="0" noHBand="0" noVBand="0"/>
				</tblPr>
				<tblGrid>
					<gridCol w="413"/>
					<gridCol w="413"/>
					<gridCol w="413"/>
				</tblGrid>
				<tr>
					<tc>
						<p>
							<r>
								<t>Hello</t>
							</r>
						</p>
					</tc>
					<tc>
						<p>
							<r>
								<t>Hello</t>
							</r>
						</p>
					</tc>
					<tc>
						<p>
							<r>
								<t>Hello</t>
							</r>
						</p>
					</tc>
				</tr>
			</tbl>
	*/

	// parse xml
	// f := bytes.NewBuffer([]byte(rawXml))
	// br := bufio.NewReaderSize(f, 1024*1024)
	// fmt.Println(xmlName)
	f, err := os.OpenFile(xmlName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Read xml file err:", err)
		return
	}
	// f := bytes.NewReader([]byte(rawXml))
	// parser := xmlparser.NewXMLParser(f, "book")
	// .SkipElements([]string{
	// 	"spacing", "szCs",
	// })
	// fmt.Println("Start:")
	tok := xmltokenizer.New(f)
	// loop:
	for {
		token, err := tok.Token() // Token is only valid until next tok.Token() invocation (short-lived object).
		if err == io.EOF {
			break
		}
		if err != nil {
			// fmt.Println("Err:", err)
			panic(err)
		}
		switch string(token.Name.Local) { // This do not allocate 🥳👍
		case "p":
			// Reuse Token object in the sync.Pool since we only use it temporarily.
			paragraph := new(P)
			se := xmltokenizer.GetToken().Copy(token) // se: StartElement, we should copy it since token is a short-lived object.
			err = paragraph.UnmarshalToken(tok, se)

			xmltokenizer.PutToken(se) // Put back to sync.Pool.
			if err != nil {
				continue
			}

			paragraphs = append(paragraphs, paragraph.Text())

		case "tbl":
			tablePar := new(Tbl)
			se := xmltokenizer.GetToken().Copy(token) // se: StartElement, we should copy it since token is a short-lived object.
			err = tablePar.UnmarshalToken(tok, se)
			if err == nil {
				// panic(err)
				paragraphs = append(paragraphs, tablePar.String())
			}

		}
	}

	// time.Sleep(1 * time.Second)

	return
}
