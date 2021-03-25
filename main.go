package main

import (
	"flag"
	"log"
	"os"
	"uclRequester/getPl"
)

func main() {
	mainLog := log.New(os.Stdout, "main: ", log.Ldate|log.Ltime|log.Lshortfile)
	pathPtr := flag.String("path", "", "path to the protected Document")
	flag.Parse()
	path := *pathPtr
	pl := getPl.GetPublishingLicense(path)
	mainLog.Println("PublishingLicense")
	mainLog.Println(pl)
}
