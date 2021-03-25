package getPl

import (
	"log"
	"os"

	"github.com/richardlehane/mscfb"
)

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
	var pl string
	for entry, err := protectedDoc.Next(); err == nil; entry, err = protectedDoc.Next() {
		if entry.Name == "Primary" {
			getPlLog.Println(entry.Name)
			getPlLog.Println(entry.Size)
			buf := make([]byte, entry.Size)
			i, err := entry.Read(buf)
			if err != nil {
				getPlLog.Fatalln(err)
			}
			if i > 0 {
				pl = string(buf[:i])
			} else {
				getPlLog.Fatalln("Read of the PL failed")
			}

		}
	}

	return pl
}
