package Formatting

import (
	"log"
	"net/http"
	"os"
)

// Want to run static and dynamic analysis to determine the maliciousness of the attachment.
func Format_Attachments(attachment_b64 string) {
	go static_analysis(attachment_b64)
	go dynamic_analysis(attachment_b64)
}

// Will have to move to a sandbox
func dynamic_analysis(attachment_b64 string) {
	var filetype string
	var err error
	filetype, err = determine_file_type(attachment_b64)
	if err != nil {
		log.Println(err)
	}
	log.Println(filetype)
}

func determine_file_type(attachment_b64 string) (string, error) {
	// Must download then open the file on the sandbox
	f, err := os.Open("golangcode.pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func static_analysis(attachment_b64 string) {
	// Going to just run strings command and run through open threat intelligence if customer has option configured in settings

}
