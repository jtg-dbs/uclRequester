package getGic

import (
	"bufio"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// GetGicCert searches for GIC files and returns the content if there is exactly one file
func GetGicCert() (gic string) {
	var gicCertString string
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
	// gicCert, err := ioutil.ReadFile(gicCerts[0])
	// if err != nil {
	// 	getGicLog.Fatal(err)
	// }
	file, err := os.Open(gicCerts[0])
	if err != nil {
		getGicLog.Fatal(err)
	}
	defer file.Close()

	// fileStat, _ := file.Stat()
	// gicByte := make([]byte, fileStat.Size())
	// _, err = file.Read(gicByte)
	// if err != nil {
	// 	getGicLog.Fatal(err)
	// }
	//transform Unicode16 Inpot from https://stackoverflow.com/questions/15783830/how-to-read-utf16-text-file-to-string-in-golang
	scanner := bufio.NewScanner(transform.NewReader(file, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()))
	for scanner.Scan() {
		gicCertString = scanner.Text()
		getGicLog.Print("Test")
	}
	return gicCertString
}
