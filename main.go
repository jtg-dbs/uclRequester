package main

import (
	"flag"
	"log"
	"os"
	"uclRequester/getGic"
	"uclRequester/getPl"
)

func main() {
	mainLog := log.New(os.Stdout, "main: ", log.Ldate|log.Ltime|log.Lshortfile)
	pathPtr := flag.String("path", "", "path to the protected Document")
	flag.Parse()
	path := *pathPtr
	mrl, clc := getPl.GetPublishingLicense(path)
	mainLog.Println(mrl[:50])
	mainLog.Println(clc[:50])
	gic := getGic.GetGicCert()
	mainLog.Print(gic[:50])
}
