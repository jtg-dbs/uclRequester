package main

import (
	"flag"
	"log"
	"os"
	"uclRequester/createSoap"
	"uclRequester/getGic"
	"uclRequester/getPl"
)

func main() {
	mainLog := log.New(os.Stdout, "main: ", log.Ldate|log.Ltime|log.Lshortfile)
	pathPtr := flag.String("path", "", "path to the protected Document")
	flag.Parse()
	path := *pathPtr
	mrl, clc := getPl.GetPublishingLicense(path)
	gic := getGic.GetGicCert()
	mainLog.Print("Soap")
	createSoap.UclSoapRequest(mrl, clc, gic)
	//ToDo: Kerberos oder NTLM Authentifizierung einbauen
	//ToDO: UserAgent emulieren
}
