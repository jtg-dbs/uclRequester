package createSoap

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/dpotapov/go-spnego"
	"github.com/gonutz/w32/v2"
)

// UclSoapRequest creates the SOAP Request for the Ucl endpoint. Authentication is done by this function
func UclSoapRequest(mrl string, clc string, gic string) {
	uclSoapRequestLog := log.New(os.Stdout, "UclSoapRequest: ", log.Ldate|log.Ltime|log.Lshortfile)
	winVersion := getWinVersion(uclSoapRequestLog)
	headers := map[string]string{
		"User-Agent":      "MSIPC;version=1.0.4349.3;AppName=WINWORD.EXE;AppVersion=16.0.13801.20294;AppArch=x86;OSName=Windows;OSVersion=" + winVersion + ";OSArch=amd64",
		"Content-Type":    "text/xml; charset=utf-8",
		"Accept-Encoding": "gzip, deflate",
		"SOAPAction":      `"http://microsoft.com/DRM/LicensingService/AcquireLicense"`,
	}
	soapClient := &http.Client{
		Transport: &spnego.Transport{},
	}

	url := getUrl(gic, uclSoapRequestLog) + "/licensing/License.amx"
	uclSoapRequestLog.Print(url)
	soapRequest, err := http.NewRequest("POST", url, nil) // XML Document has to be in body
	for key, value := range headers {
		soapRequest.Header.Set(key, value)
	}
	if err != nil {
		uclSoapRequestLog.Fatal(err)
	}
	resp, err := soapClient.Do(soapRequest)
	uclSoapRequestLog.Print(err)
	uclSoapRequestLog.Print(resp)
}

//extracts the ad rms server url from the gic
func getUrl(gic string, log *log.Logger) (url string) {
	//add root element for valid xml
	gicWellFormed := "<root>" + gic + "</root>"
	gicXml, err := xmlquery.Parse(strings.NewReader(gicWellFormed))
	if err != nil {
		log.Fatal(err)
	}
	urlNode := xmlquery.FindOne(gicXml, `//XrML/BODY/ISSUER/OBJECT[@type="MS-DRM-Server"]/ADDRESS`)
	return urlNode.InnerText()
}

// extracts the Windows Version from the registry
func getWinVersion(log *log.Logger) (version string) {
	versionWin := w32.RtlGetVersion()
	return fmt.Sprint(versionWin.MajorVersion) + "." + fmt.Sprint(versionWin.MinorVersion) + "." + fmt.Sprint(versionWin.BuildNumber)
}
