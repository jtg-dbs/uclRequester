package getGic

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	uni "unicode"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// GetGicCert searches for GIC files and returns the content if there is exactly one file
func GetGicCert() (gic string, gic_name string) {
	getGicLog := log.New(os.Stdout, "GetGicCertificate: ", log.Ldate|log.Ltime|log.Lshortfile)
	homedir, _ := os.UserHomeDir()
	dir := homedir + "/AppData/Local/Microsoft/MSIPC/GIC*"
	gicCerts, err := filepath.Glob(dir)
	if err != nil {
		getGicLog.Fatal(err)
	}
	if len(gicCerts) != 1 {
		getGicLog.Fatal("Found 0 or more than 1 Files. Number of Files found" + string(len(gicCerts)))
		// ToDo: Was wenn es merh asl ein GIC gib und beide zu probieren
	}

	file, err := os.Open(gicCerts[0])
	if err != nil {
		getGicLog.Fatal(err)
	}
	defer file.Close()

	//transform Unicode16 Import from https://forum.golangbridge.org/t/reading-a-utf-16-text-file/1496
	scanner := transform.NewReader(file, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder())
	gicCertString, err := ioutil.ReadAll(scanner)
	if err != nil {
		getGicLog.Fatal(err)
	}

	// Cleaning the XML https://blog.zikes.me/post/cleaning-xml-files-before-unmarshaling-in-go/
	printOnly := func(r rune) rune {
		if uni.IsPrint(r) {
			return r
		}
		return -1
	}
	gicCertString = []byte(strings.Map(printOnly, string(gicCertString)))
	return string(gicCertString), gicCerts[0]
}
