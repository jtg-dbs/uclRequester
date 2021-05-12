package createSoap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/dpotapov/go-spnego"
	"github.com/gonutz/w32/v2"
)

// UclSoapRequest creates the SOAP Request for the Ucl endpoint. Authentication is done by this function
func UclSoapRequest(mrl string, clc string, gic string) (eul string) {
	uclSoapRequestLog := log.New(os.Stdout, "UclSoapRequest: ", log.Ldate|log.Ltime|log.Lshortfile)
	winVersion := getWinVersion(uclSoapRequestLog)
	headers := map[string]string{
		"User-Agent":      "MSIPC;version=1.0.4349.3;AppName=WINWORD.EXE;AppVersion=16.0.13801.20294;AppArch=x86;OSName=Windows;OSVersion=" + winVersion + ";OSArch=amd64",
		"Content-Type":    "text/xml; charset=utf-8",
		"Accept-Encoding": "gzip, deflate",
		"SOAPAction":      `"http://microsoft.com/DRM/LicensingService/AcquireLicense"`,
		"Accept":          "*/*",
		"Connection":      "Keep-Alive",
		// both ids are necessary but can be recycled every time, could be a problem for Detection with a SIEM
		"rmsocorrelationid": "{9DE381BF-2552-4713-8CF4-A43B548F2A05}",
		"rmsoscenarioid":    "{F966E37B-14B2-4E85-8E3B-3870C0B51242}",
	}
	soapClient := &http.Client{
		Transport: &spnego.Transport{},
	}
	xmlBody := newEnvelope()
	body, err := xml.Marshal(xmlBody)
	if err != nil {
		uclSoapRequestLog.Print(err)
	}
	bodyFilled := insertContent(mrl, clc, gic, string(body))
	url := getUrl(gic, uclSoapRequestLog) + "/licensing/License.asmx"
	uclSoapRequestLog.Print(url)
	soapRequest, err := http.NewRequest("POST", url, bytes.NewReader([]byte(bodyFilled))) // XML Document has to be in body

	requestHeader := make(http.Header)
	for key, value := range headers {
		requestHeader[key] = []string{value}
	}
	soapRequest.Header = requestHeader
	if err != nil {
		uclSoapRequestLog.Fatal(err)
	}
	resp, err := soapClient.Do(soapRequest)
	if err != nil {
		uclSoapRequestLog.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		uclSoapRequestLog.Println(resp.Status)
		uclSoapRequestLog.Fatalln(resp.Body)
	}

	eul_byte, err := io.ReadAll(resp.Body)
	if err != nil {
		uclSoapRequestLog.Fatal(err)
	}
	return string(eul_byte)
}

func insertContent(mrl string, clc string, gic string, bodyEmpty string) (bodyFilled string) {
	IssuanceLicense := `<IssuanceLicense soapenc:arrayType="Certificate[2]">`
	LicenseeCerts := `<LicenseeCerts soapenc:arrayType="Certificate[1]">`
	xmlHead := `<?xml version="1.0" encoding="utf-8"?>`
	indexArray1 := strings.Index(bodyEmpty, IssuanceLicense) + len(IssuanceLicense)
	body := bodyEmpty[:indexArray1] + "<Certificate>" + mrl + "</Certificate>" + "<Certificate>" + clc + "</Certificate>" + bodyEmpty[indexArray1:]
	indexArray2 := strings.Index(body, LicenseeCerts) + len(LicenseeCerts)
	bodyFilled = body[:indexArray2] + "<Certificate>" + gic + "</Certificate>" + body[indexArray2:]
	bodyFilled = xmlHead + bodyFilled
	return bodyFilled
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

type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soapenc string   `xml:"xmlns:soapenc,attr"`
	Header  struct {
		Text        string `xml:",chardata"`
		VersionData struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			MinimumVersion string `xml:"MinimumVersion"`
			MaximumVersion string `xml:"MaximumVersion"`
		} `xml:"VersionData"`
	} `xml:"soap:Header"`
	Body struct {
		Text           string `xml:",chardata"`
		AcquireLicense struct {
			Text          string `xml:",chardata"`
			Xmlns         string `xml:"xmlns,attr"`
			RequestParams struct {
				Text                 string `xml:",chardata"`
				Xmlns                string `xml:"xmlns,attr"`
				AcquireLicenseParams struct {
					Text            string `xml:",chardata"`
					IssuanceLicense struct {
						Text        string   `xml:",chardata"`
						ArrayType   string   `xml:"soapenc:arrayType,attr"`
						Certificate []string `xml:"Certificate"`
					} `xml:"IssuanceLicense"`
					LicenseeCerts struct {
						Text        string   `xml:",chardata"`
						ArrayType   string   `xml:"soapenc:arrayType,attr"`
						Certificate []string `xml:"Certificate"`
					} `xml:"LicenseeCerts"`
				} `xml:"AcquireLicenseParams"`
			} `xml:"RequestParams"`
		} `xml:"AcquireLicense"`
	} `xml:"soap:Body"`
}

func newEnvelope() Envelope {
	envelope := Envelope{}
	envelope.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	envelope.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	envelope.Xsd = "http://www.w3.org/2001/XMLSchema"
	envelope.Soapenc = "http://schemas.xmlsoap.org/soap/encoding/"
	envelope.Header.VersionData.Xmlns = "http://microsoft.com/DRM/LicensingService"
	envelope.Header.VersionData.MaximumVersion = "1.0.0.0"
	envelope.Header.VersionData.MinimumVersion = "1.0.0.0"
	envelope.Body.AcquireLicense.Xmlns = "http://microsoft.com/DRM/LicensingService"
	envelope.Body.AcquireLicense.RequestParams.Xmlns = "http://microsoft.com/DRM/LicensingService"
	envelope.Body.AcquireLicense.RequestParams.AcquireLicenseParams.IssuanceLicense.ArrayType = "Certificate[2]"
	envelope.Body.AcquireLicense.RequestParams.AcquireLicenseParams.LicenseeCerts.ArrayType = "Certificate[1]"
	return envelope
}
