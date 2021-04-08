package getPl

import (
	"log"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/richardlehane/mscfb"
)

func GetPublishingLicense(path string) (MicrosoftRightsLabel string, ClienLicensorCertificate string) {
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

	return getMrlClc(pl, getPlLog)

}

//cleanXml In the PL there is are some type of text before the xml. Furhtermore the needed Elements get extracted
func getMrlClc(xmlByte []byte, getPlLog *log.Logger) (MicrosoftRightsLabel string, ClienLicensorCertificate string) {
	var mrl xmlquery.Node
	var clc xmlquery.Node
	// get ride of the text before the xml
	xmlStr := string(xmlByte)
	index := strings.Index(xmlStr, "<")
	xmlDoc := addRootElement(xmlStr[index:])
	publishingLicense, err := xmlquery.Parse(strings.NewReader(xmlDoc))
	if err != nil {
		getPlLog.Fatal(err)
	}
	// extract Microsoft Right Label and Client Licensor Certificate from the Pl
	xrmlArr := xmlquery.Find(publishingLicense, `//XrML`)
	for _, node := range xrmlArr {
		if xmlquery.FindOne(node, `//BODY[@type="Microsoft Rights Label"]`) != nil {
			mrl = *node
			getPlLog.Print("Microsoft Rights Label found")
		}
		if xmlquery.FindOne(node, `//BODY/DESCRIPTOR/OBJECT[@type="Client-Licensor-Certificate"]`) != nil {
			clc = *node
			getPlLog.Print("Client-Licensor-Certificate found")
		}
	}

	return mrl.OutputXML(true), clc.OutputXML(true)
}

// add a root element in the PL for a well formed XML
func addRootElement(xmlStr string) string {
	index := strings.Index(xmlStr, "?>") + 2
	wellFormedXml := xmlStr[:index] + "<root>" + xmlStr[index:] + "</root>"

	return wellFormedXml
}
