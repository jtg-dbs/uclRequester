package getPl

import (
	"encoding/xml"
	"log"
	"os"
	"strings"

	"github.com/richardlehane/mscfb"
)
// GetPublishingLicense retrieves the path to a protected document and response with the PL of this
func GetPublishingLicense(path string) (string, string) {
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
func cleanXml(xmlByte []byte, getPlLog *log.Logger) (string, string) {
	// get ride of the text before the xml
	xmlStr := string(xmlByte)
	index := strings.Index(xmlStr, "<")
	xmlDoc := xmlStr[index:]
	clcXmlDoc := xmlStr[index:]
	

	return string(mrLabelPl), string(clcPl)
}
