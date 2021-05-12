package main

import (
	"flag"
	"log"
	"os"
	"uclRequester/createEul"
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
	gic, gic_name := getGic.GetGicCert()
	mainLog.Print("Soap")
	eul := createSoap.UclSoapRequest(mrl, clc, gic)
	createEul.CreateEulFile(eul, gic_name)
}
