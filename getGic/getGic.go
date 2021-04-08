package getGic

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GetGicCert searches for GIC files and returns the content if there is exactly one file
func GetGicCert() (gic string) {
	getGicLog := log.New(os.Stdout, "GetGicCertificate: ", log.Ldate|log.Ltime|log.Lshortfile)
	homedir, _ := os.UserHomeDir()
	dir := homedir + "/AppData/Local/Microsoft/MSIPC/GIC*"
	gicCerts, err := filepath.Glob(dir)
	if err != nil {
		getGicLog.Fatal(err)
	}
	if len(gicCerts) != 1 {
		getGicLog.Fatal("Found 0 or more than 1 Files. Number of Files found" + string(len(gicCerts)))
	}
	gicCert, err := ioutil.ReadFile(gicCerts[0])
	if err != nil {
		getGicLog.Fatal(err)
	}
	return string(gicCert)
}
