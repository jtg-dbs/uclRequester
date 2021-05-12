package getPl

import (
	"log"
	"os"
	"strings"

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
	var mrl string
	var clc string
	// get ride of the text before the xml
	xmlStr := string(xmlByte)
	index := strings.Index(xmlStr, "<")
	plObj := strings.Split(xmlStr[index:], "<XrML")

	for _, data := range plObj {
		if strings.Contains(data, "Microsoft Rights Label") {
			mrl = "<XrML" + data
			getPlLog.Print("Fond: MRL")
		}
		if strings.Contains(data, "Client-Licensor-Certificate") {
			clc = "<XrML" + data
			getPlLog.Print("Found CLC")
		}

	}

	if clc == "" {
		getPlLog.Print("No CLC found")
	}
	return mrl, clc
}

// add a root element in the PL for a well formed XML
// func addRootElement(xmlStr string) string {
// 	index := strings.Index(xmlStr, "?>") + 2
// 	wellFormedXml := xmlStr[:index] + "<root>" + xmlStr[index:] + "</root>"

// 	return wellFormedXml
// }
