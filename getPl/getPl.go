package getPl

import (
	"encoding/xml"
	"log"
	"os"
	"strings"

	"github.com/richardlehane/mscfb"
)

// created with https://www.onlinetool.io/xmltogo/
type XrML struct {
	XMLName xml.Name `xml:"XrML"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Purpose string   `xml:"purpose,attr"`
	BODY    struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		Version    string `xml:"version,attr"`
		ISSUEDTIME string `xml:"ISSUEDTIME"`
		DESCRIPTOR struct {
			Text   string `xml:",chardata"`
			OBJECT struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				ID   struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ID"`
				NAME string `xml:"NAME"`
			} `xml:"OBJECT"`
		} `xml:"DESCRIPTOR"`
		ISSUER struct {
			Text   string `xml:",chardata"`
			OBJECT struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				ID   struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ID"`
				NAME    string `xml:"NAME"`
				ADDRESS struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ADDRESS"`
			} `xml:"OBJECT"`
			PUBLICKEY struct {
				Text      string `xml:",chardata"`
				ALGORITHM string `xml:"ALGORITHM"`
				PARAMETER []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					VALUE struct {
						Text     string `xml:",chardata"`
						Encoding string `xml:"encoding,attr"`
						Size     string `xml:"size,attr"`
					} `xml:"VALUE"`
				} `xml:"PARAMETER"`
			} `xml:"PUBLICKEY"`
			SECURITYLEVEL []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value,attr"`
			} `xml:"SECURITYLEVEL"`
		} `xml:"ISSUER"`
		DISTRIBUTIONPOINT []struct {
			Text   string `xml:",chardata"`
			OBJECT struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				ID   struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ID"`
				NAME    string `xml:"NAME"`
				ADDRESS struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ADDRESS"`
			} `xml:"OBJECT"`
		} `xml:"DISTRIBUTIONPOINT"`
		ISSUEDPRINCIPALS struct {
			Text      string `xml:",chardata"`
			PRINCIPAL struct {
				Text       string `xml:",chardata"`
				InternalID string `xml:"internal-id,attr"`
				OBJECT     struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
					ID   struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"ID"`
					NAME    string `xml:"NAME"`
					ADDRESS struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"ADDRESS"`
				} `xml:"OBJECT"`
				PUBLICKEY struct {
					Text      string `xml:",chardata"`
					ALGORITHM string `xml:"ALGORITHM"`
					PARAMETER []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						VALUE struct {
							Text     string `xml:",chardata"`
							Encoding string `xml:"encoding,attr"`
							Size     string `xml:"size,attr"`
						} `xml:"VALUE"`
					} `xml:"PARAMETER"`
				} `xml:"PUBLICKEY"`
				SECURITYLEVEL []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Value string `xml:"value,attr"`
				} `xml:"SECURITYLEVEL"`
				ENABLINGBITS struct {
					Text  string `xml:",chardata"`
					Type  string `xml:"type,attr"`
					VALUE struct {
						Text     string `xml:",chardata"`
						Encoding string `xml:"encoding,attr"`
						Size     string `xml:"size,attr"`
					} `xml:"VALUE"`
				} `xml:"ENABLINGBITS"`
			} `xml:"PRINCIPAL"`
		} `xml:"ISSUEDPRINCIPALS"`
		WORK struct {
			Text   string `xml:",chardata"`
			OBJECT struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				ID   struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"ID"`
			} `xml:"OBJECT"`
			METADATA struct {
				Text  string `xml:",chardata"`
				OWNER struct {
					Text   string `xml:",chardata"`
					OBJECT struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"ID"`
						NAME string `xml:"NAME"`
					} `xml:"OBJECT"`
				} `xml:"OWNER"`
			} `xml:"METADATA"`
			RIGHTSGROUP struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"name,attr"`
				RIGHTSLIST struct {
					Text  string `xml:",chardata"`
					RIGHT struct {
						Text          string `xml:",chardata"`
						Name          string `xml:"name,attr"`
						CONDITIONLIST struct {
							Text string `xml:",chardata"`
							TIME struct {
								Text      string `xml:",chardata"`
								RANGETIME struct {
									Text  string `xml:",chardata"`
									FROM  string `xml:"FROM"`
									UNTIL string `xml:"UNTIL"`
								} `xml:"RANGETIME"`
							} `xml:"TIME"`
							ACCESS struct {
								Text      string `xml:",chardata"`
								PRINCIPAL struct {
									Text         string `xml:",chardata"`
									InternalID   string `xml:"internal-id,attr"`
									ENABLINGBITS struct {
										Text  string `xml:",chardata"`
										Type  string `xml:"type,attr"`
										VALUE struct {
											Text     string `xml:",chardata"`
											Encoding string `xml:"encoding,attr"`
											Size     string `xml:"size,attr"`
										} `xml:"VALUE"`
									} `xml:"ENABLINGBITS"`
								} `xml:"PRINCIPAL"`
							} `xml:"ACCESS"`
						} `xml:"CONDITIONLIST"`
						ACCESS struct {
							Text      string `xml:",chardata"`
							PRINCIPAL struct {
								Text         string `xml:",chardata"`
								InternalID   string `xml:"internal-id,attr"`
								ENABLINGBITS struct {
									Text  string `xml:",chardata"`
									Type  string `xml:"type,attr"`
									VALUE struct {
										Text     string `xml:",chardata"`
										Encoding string `xml:"encoding,attr"`
										Size     string `xml:"size,attr"`
									} `xml:"VALUE"`
								} `xml:"ENABLINGBITS"`
							} `xml:"PRINCIPAL"`
						} `xml:"ACCESS"`
					} `xml:"RIGHT"`
				} `xml:"RIGHTSLIST"`
			} `xml:"RIGHTSGROUP"`
		} `xml:"WORK"`
		AUTHENTICATEDDATA []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name string `xml:"name,attr"`
		} `xml:"AUTHENTICATEDDATA"`
	} `xml:"BODY"`
	SIGNATURE struct {
		Text      string `xml:",chardata"`
		ALGORITHM string `xml:"ALGORITHM"`
		DIGEST    struct {
			Text      string `xml:",chardata"`
			ALGORITHM string `xml:"ALGORITHM"`
			PARAMETER struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				VALUE struct {
					Text     string `xml:",chardata"`
					Encoding string `xml:"encoding,attr"`
				} `xml:"VALUE"`
			} `xml:"PARAMETER"`
			VALUE struct {
				Text     string `xml:",chardata"`
				Encoding string `xml:"encoding,attr"`
				Size     string `xml:"size,attr"`
			} `xml:"VALUE"`
		} `xml:"DIGEST"`
		VALUE struct {
			Text     string `xml:",chardata"`
			Encoding string `xml:"encoding,attr"`
			Size     string `xml:"size,attr"`
		} `xml:"VALUE"`
	} `xml:"SIGNATURE"`
}

// GetPublishingLicense retrieves the path to a protected document and response with the PL of this
func GetPublishingLicense(path string) string {
	getPlLog := log.New(os.Stdout, "GetPublishingLicense: ", log.Ldate|log.Ltime|log.Lshortfile)
	file, err := os.Open(path)
	// defer let this function proced when GetPublishingLicense is finished
	defer file.Close()
	if err != nil {
		getPlLog.Fatalln(err)
	}

	protectedDoc, err := mscfb.New(file)
	if err != nil {
		getPlLog.Fatalln(err)
	}
	// iterate through the file to finde the PL
	var pl []byte
	for entry, err := protectedDoc.Next(); err == nil; entry, err = protectedDoc.Next() {
		if entry.Name == "Primary" {
			buf := make([]byte, entry.Size)
			i, err := entry.Read(buf)
			if err != nil {
				getPlLog.Fatalln(err)
			}
			if i > 0 {
				pl = buf[:i]
			} else {
				getPlLog.Fatalln("Read of the PL failed")
			}

		}
	}

	return cleanXml(pl, getPlLog)
}

//cleanXml In the PL there is are some type of text before the xml.´
func cleanXml(xmlByte []byte, getPlLog *log.Logger) string {
	// get ride of the text before the xml
	xmlStr := string(xmlByte)
	index := strings.Index(xmlStr, "<")
	xmlDoc := xmlStr[index:]
	// encode the string as an xml
	var xrml XrML
	err := xml.Unmarshal([]byte(xmlDoc), &xrml)
	if err != nil {
		getPlLog.Fatal(err)
	}
	xmlPl, err := xml.Marshal(&xrml)
	if err != nil {
		getPlLog.Fatal(err)
	}
	return string(xmlPl)
}
