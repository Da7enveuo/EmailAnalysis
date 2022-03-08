package Formatting

import (
//"golang.org/x/net/html"
//"io"
//"log"
)

func Format_HTML_Body(htmlBody string) {
	//tokenizer := html.NewTokenizer(htmlBody)
	//for {
	//get the next token type
	//	tokenType := tokenizer.Next()

	//if it's an error token, we either reached
	//the end of the file, or the HTML was malformed
	//	if tokenType == html.ErrorToken {
	//		err := tokenizer.Err()
	//		if err == io.EOF {
	//end of the file, break out of the loop
	//			break
	//		}
	//otherwise, there was an error tokenizing,
	//which likely means the HTML was malformed.
	//since this is a simple command-line utility,
	//we can just use log.Fatalf() to report the error
	//and exit the process with a non-zero status code
	//		log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
	//	} // else if tokenType == html.Link{
	//	htmlLinkParse(tokenizer value?)
	//}
	//want to check if token type is a link type
	//}
}
