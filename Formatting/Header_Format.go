package Formatting

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// Planning to change the main in this when its integrated with the other code. Also planning on migrating to the same package as all the other code too.
func Format_Header(email_headers string) {
	// Open logging file and log timestamp and that header_analysis program has began
	//json object initialization
	srsj := Sender_Recipient_Subject_Json{}
	aj := Authentictation_Json{}
	aij := Additional_Information_Json{}
	Completely_Constructed_Header_Json := Master_Json{}
	// This is to split the string by ":" and get key-value pairs

	// This is the sync waitgroup variable that helps to make sure goroutines complete before proceeding in code
	var wg sync.WaitGroup
	// we will handle other types of values later.
	deliminated_headers := strings.Split(email_headers, "\n")
	// Parsing function for headers
	parsing(deliminated_headers, &srsj, &aj, &aij)

	// log that the go routines have begun in the header_analysis log
	wg.Add(2)
	// Note, we are adding all the authentication information to a large string called large_auth_string, we just need to make variables to the individual ones and channels
	// This function will take in the raw email header for recipients, senders, subject and will return them formatted properly to the channels.
	// Will then create json object once completed with everything
	go sender_recipient_sub(&srsj, &wg)
	// Authentication will take in the raw email header for authentication and ARC authentication, DKIM, SPF, DMARC and other authentication methods and
	// return through channels the formatted strings to add to the json object
	go authentication(&aj, &wg)
	// This waits for the go routines ahead to be finished before continuing
	wg.Wait()
	// Add log to header_analysis log that the functions have completed

	//wg.Add(1)
	// This function will do analysis on the sender domain, recipient, and subject. This may require api keys so it would end up being an optional function possibly unless can find opensource without api key and rate limit
	// I suppose we could do a ml or other kind of analysis on the subject, domain, and recipient to see how likely it appears that it is phishing/social-engineering/spam/legitimate.
	//go sender_recipient_sub_analysis(&srsj, &wg)
	//go authentication_analysis(&aj, &wg)
	//wg.Wait()

	// Assigning the Master Json fields to the constructed fields to create master json object
	Completely_Constructed_Header_Json.Sender_recipient_subject = srsj
	Completely_Constructed_Header_Json.Authentication = aj
	Completely_Constructed_Header_Json.Additional_headers = aij
	// Encoding into a json object
	Complete_Header_Json, err := json.MarshalIndent(Completely_Constructed_Header_Json, "", "    ")
	// Error handling
	if err != nil {
		// log if error
		fmt.Println(err)
	}
	// Transferring json object into base64 string
	base64_json := base64.StdEncoding.EncodeToString([]byte(string(Complete_Header_Json)))
	fmt.Println("Base64 Representation of Header Json Object:\n" + base64_json + "\n\n\n")
	//We would have the sql code to log the base64 json into the database right here.

	//This part is going to be a different function call that serves the api and gui endpoints. Just in for testing right now.
	from_base64_json, err := base64.StdEncoding.DecodeString(base64_json)
	if err != nil {
		//log if error
		fmt.Println(err)
	}
	fmt.Println("Header Json Object:\n" + string(from_base64_json))
}

func format_string_list_reverse(list []string) (return_list []string) {
	new_list := []string{}
	for _, value := range list {
		new_list = append(new_list, authentication_analysis_loop(value))
	}
	temp_list := make([]string, len(new_list))
	for i := 0; i < len(new_list); i++ {
		temp_list[len(new_list)-i-1] = strings.Replace(new_list[i], "\r", "", -1)
	}
	return_list = temp_list
	return return_list
}
