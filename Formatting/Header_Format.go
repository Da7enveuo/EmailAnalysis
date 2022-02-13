package Formatting

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type Master_Json struct {
	Sender_recipient_subject Sender_Recipient_Subject_Json
	Authentication           Authentictation_Json
	Additional_headers       Additional_Information_Json
}

type Authentictation_Json struct {
	X_SONIC_DKIM_SIGN, Authentication_Results, ARC_Authentication_Results, ARC_Seal, ARC_Message_Signature, X_Google_DKIM_Signature string
	Received_Path, Received_SPF, DKIM_Signature                                                                                     []string
	Analysis_Score                                                                                                                  int64
}

// add pgp encryption
type Sender_Recipient_Subject_Json struct {
	Sender_name, Sender_email, Rec_name, Rec_email, Subject, Date string
	Analysis_Score                                                int64
}

// lots to add to this, may also want to break into different groups in addition.
type Additional_Information_Json struct {
	Mime_Type, Message_ID, X_Feedback_ID, X_Google_Smtp_Source, MIME_Version, Content_Disposition, Content_Transfer_Encoding, X_Mailer, X_Gm_Message_State, Message_Id, Content_Type, Content_Length string
}

// Need to make one of the analysis loops able to determine the amount of spaces to cut off the beginning and determine which groups with what.

// Planning to change the main in this when its integrated with the other code. Also planning on migrating to the same package as all the other code too.
func Format_Header(email_headers string) {
	// Open logging file and log timestamp and that header_analysis program has began
	//json object initialization
	srsj := Sender_Recipient_Subject_Json{}
	aj := Authentictation_Json{}
	aij := Additional_Information_Json{}
	Completely_Constructed_Header_Json := Master_Json{}
	// This is to split the string by ":" and get key-value pairs
	var key_value []string
	// This is the sync waitgroup variable that helps to make sure goroutines complete before proceeding in code
	var wg sync.WaitGroup
	// we will handle other types of values later.
	deliminated_headers := strings.Split(email_headers, "\n")
	for key, value := range deliminated_headers {
		//This first string is getting the fields that are deliminated by ":"
		//Right now we are handling the from and to addresses, we need to get more fields and trace the path of the headers
		// Ugly parsing, can likely find some other way
		if strings.Contains(value, ":") == true {
			key_value = strings.Split(value, ":")
			// Checking to see if when we split by ":" if there are more than 2 entries, indicating a multiple split which will mess it up.
			// Want a way to make it so that it uses the first split as the key-value pair and then appends all of the remaining key_values together to be the value
			if len(key_value) > 2 {
				joined_strings := strings.Join(key_value[1:len(key_value)], ":")
				key_value = []string{key_value[0], joined_strings}
			}
			if key_value[0] == "From" {
				srsj.Sender_name = strings.Replace(key_value[1], "\r", "", -1)
				//from_address_string = string(key_value[1])
			} else if key_value[0] == "To" {
				srsj.Rec_name = strings.Replace(key_value[1], "\r", "", -1)
				//recipients_string = key_value[1]
			} else if key_value[0] == "Subject" {
				first_character := key_value[1][0:1]
				if first_character == " " {
					srsj.Subject = strings.Replace(key_value[1][1:], "\r", "", -1)
				} else {
					srsj.Subject = strings.Replace(key_value[1], "\r", "", -1)
				}
			} else if key_value[0] == "Date" {
				srsj.Date = strings.Replace(key_value[1], "\r", "", -1)
			} else if strings.ToLower(key_value[0]) == "authentication-results" || strings.ToLower(key_value[0]) == "arc-authentication-results" || strings.ToLower(key_value[0]) == "arc-message-signature" || strings.ToLower(key_value[0]) == "arc-seal" || strings.ToLower(key_value[0]) == "x_sonic_dkim_sign" || strings.ToLower(key_value[0]) == "x-google-dkim-signature" {
				// Need to make this so that it is a list of the results and then we can check if its a single or multiple and handle each differently. This allows us to get all information and analyze it.
				switch strings.ToLower(key_value[0]) {
				case "authentication-results":
					aj.Authentication_Results = auth_sub_loop_for_analysis(deliminated_headers, key, key_value)
				case "arc-authentication-results":
					aj.ARC_Authentication_Results = auth_sub_loop_for_analysis(deliminated_headers, key, key_value)
				case "arc-message-signature":
					aj.ARC_Message_Signature = auth_sub_loop_for_analysis(deliminated_headers, key, key_value)
				case "arc-seal":
					aj.ARC_Seal = auth_sub_loop_for_analysis(deliminated_headers, key, key_value)
				case "x-google-dkim-signature":
					aj.X_Google_DKIM_Signature = auth_sub_loop_for_analysis(deliminated_headers, key, key_value)
				}
			} else if strings.ToLower(key_value[0]) == "x-sonic-dkim-sign" || strings.ToLower(key_value[0]) == "x-google-smtp-source" || strings.ToLower(key_value[0]) == "mime-version" || strings.ToLower(key_value[0]) == "content-disposition" || strings.ToLower(key_value[0]) == "content-transfer-encoding" || strings.ToLower(key_value[0]) == "x-mailer" || strings.ToLower(key_value[0]) == "x-gm-message-state" || strings.ToLower(key_value[0]) == "message-id" || strings.ToLower(key_value[0]) == "content-type" || strings.ToLower(key_value[0]) == "content-length" {
				switch strings.ToLower(key_value[0]) {
				case "x-sonic-dkim-sign":
					aj.X_SONIC_DKIM_SIGN = key_value[1]
				case "x-google-smtp-source":
					aij.X_Google_Smtp_Source = key_value[1]
				case "mime-version":
					aij.MIME_Version = key_value[1]
				case "content-disposition":
					aij.Content_Disposition = key_value[1]
				case "content-transfer-encoding":
					aij.Content_Transfer_Encoding = key_value[1]
				case "x-mailer":
					aij.X_Mailer = key_value[1]
				case "x-gm-message-state":
					aij.X_Gm_Message_State = key_value[1]
				case "message-id":
					aij.Message_Id = key_value[1]
				case "content-type":
					aij.Content_Type = key_value[1]
				case "content-length":
					aij.Content_Length = key_value[1]
				}
			} else if strings.ToLower(key_value[0]) == "received" || strings.ToLower(key_value[0]) == "x-received" || strings.ToLower(key_value[0]) == "received-spf" || strings.ToLower(key_value[0]) == "dkim-signature" {
				switch strings.ToLower(key_value[0]) {
				case "received", "x-received":
					aj.Received_Path = append(aj.Received_Path, auth_sub_loop_for_analysis(deliminated_headers, key, key_value))
				case "received-spf":
					aj.Received_SPF = append(aj.Received_SPF, auth_sub_loop_for_analysis(deliminated_headers, key, key_value))
				case "dkim-signature":
					aj.DKIM_Signature = append(aj.DKIM_Signature, auth_sub_loop_for_analysis(deliminated_headers, key, key_value))
				}
			} else {
				key_value = []string{""}
			}
		}
	}
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

// May want to split this to auth_sub_loop_for_analysis into another file and just import it instead
func authentication_analysis_loop(analysis string) (final_string string) {
	final_string = ""
	if analysis != "" {
		f_c := analysis[0:1]
		if f_c == " " {
			final_string = analysis[1:]
		} else {
			final_string = analysis
		}
		l_c := final_string[len(final_string)-1 : len(final_string)]
		if l_c == " " {
			final_string = final_string[:len(final_string)-1]
		}
	}
	return strings.Replace(final_string, "\r", "", -1)
}

// this is for formatting the individual elements of the authentication struct
func authentication(aj *Authentictation_Json, wg *sync.WaitGroup) {
	defer wg.Done()
	aj.X_SONIC_DKIM_SIGN = authentication_analysis_loop(aj.X_SONIC_DKIM_SIGN)
	aj.Authentication_Results = authentication_analysis_loop(aj.Authentication_Results)
	aj.ARC_Authentication_Results = authentication_analysis_loop(aj.ARC_Authentication_Results)
	aj.ARC_Seal = authentication_analysis_loop(aj.ARC_Seal)
	aj.ARC_Message_Signature = authentication_analysis_loop(aj.ARC_Message_Signature)
	aj.X_Google_DKIM_Signature = authentication_analysis_loop(aj.X_Google_DKIM_Signature)
	// for the []string arrays
	aj.Received_Path = format_string_list_reverse(aj.Received_Path)
	aj.Received_SPF = format_string_list_reverse(aj.Received_SPF)
	aj.DKIM_Signature = format_string_list_reverse(aj.DKIM_Signature)

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

func auth_sub_loop_for_analysis(deliminated_headers []string, key int, key_value []string) (return_string string) {
	return_string = key_value[1]
	Auth_line_counter := 1
	for true {
		if strings.Contains(string(deliminated_headers[key+Auth_line_counter]), "       ") == true {
			return_string = return_string + strings.Replace(string(deliminated_headers[key+Auth_line_counter]), "       ", "", -1)
			Auth_line_counter++
		} else {
			return return_string
		}
	}
	return strings.Replace(return_string, "\r", "", -1)
}

// May want to split here down into another file as well and import it.
func sender_recipient_sub(srsj *Sender_Recipient_Subject_Json, wg *sync.WaitGroup) {
	defer wg.Done()
	//recipients_string with srsj.Rec_name
	if strings.Contains(srsj.Rec_name, "<") == true {
		rec := strings.Split(srsj.Rec_name, "<")
		first_character_email := rec[1][0:1]
		if first_character_email == " " {
			srsj.Rec_email = strings.Replace(strings.Trim(strings.Replace(rec[1][1:], ">", "", -1), "\""), "\r", "", -1)
		} else {
			srsj.Rec_email = strings.Replace(strings.Trim(strings.Replace(rec[1], ">", "", -1), "\""), "\r", "", -1)
		}
		if strings.Contains(srsj.Rec_email, " ") == true {
			srsj.Rec_email = strings.Replace(strings.Replace(string(strings.Split(string(srsj.Rec_email), " ")[0]), ",", "", -1), "\r", "", -1)
		}
		first_letter := rec[0][0:1]
		if first_letter == " " {
			srsj.Rec_name = strings.Replace(strings.Replace(rec[0][1:], "\"", "", -1), "\r", "", -1)
		} else {
			srsj.Rec_name = strings.Replace(strings.Replace(rec[0], "\"", "", -1), "\r", "", -1)
		}
	} else {
		if len(srsj.Rec_name) == 0 || len(srsj.Rec_name) == 1 {
			srsj.Rec_name = "Not Found"
		} else {
			first_character_email := string(srsj.Rec_name[0:1])
			if first_character_email == " " {
				srsj.Rec_email = strings.Replace(strings.Replace(srsj.Rec_name[1:], "\"", "", -1), "\r", "", -1)
			} else {
				srsj.Rec_email = strings.Replace(strings.Replace(srsj.Rec_name, "\"", "", -1), "\r", "", -1)
			}
			if strings.Contains(srsj.Rec_email, " ") == true {
				srsj.Rec_email = strings.Replace(strings.Replace(string(strings.Split(string(srsj.Rec_email), " ")[0]), ",", "", -1), "\r", "", -1)
			}
			srsj.Rec_name = "Not Found"
		}
	}
	// Formatting the sender string to get the email address
	if strings.Contains(srsj.Sender_name, "<") == true {
		send := strings.Split(srsj.Sender_name, "<")
		first_character_email := send[1][0:1]
		if first_character_email == " " {
			srsj.Sender_email = strings.Replace(strings.Replace(strings.Replace(send[1][1:], ">", "", -1), "\"", "", -1), "\r", "", -1)
		} else {
			srsj.Sender_email = strings.Replace(strings.Replace(strings.Replace(send[1], ">", "", -1), "\"", "", -1), "\r", "", -1)
		}
		first_character := send[0][0:1]
		if first_character == " " {
			srsj.Sender_name = strings.Replace(strings.Replace(send[0][1:], "\"", "", -1), "\r", "", -1)
		} else {
			srsj.Sender_name = strings.Replace(strings.Replace(send[0], "\"", "", -1), "\r", "", -1)
		}
	} else {
		srsj.Sender_email = strings.Replace(strings.Replace(srsj.Sender_name, "\"", "", -1), "\r", "", -1)
		srsj.Sender_name = "Not Found"
	}
	// we will do nothing with this and just assign it outright
}
