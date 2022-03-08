package Formatting

import "strings"

func parsing(deliminated_headers []string, srsj *Sender_Recipient_Subject_Json, aj *Authentictation_Json, aij *Additional_Information_Json) {
	var key_value []string
	for key, value := range deliminated_headers {
		//This first string is getting the fields that are deliminated by ":"
		//Right now we are handling the from and to addresses, we need to get more fields and trace the path of the headers
		// Ugly parsing, can likely find some other way
		if strings.Contains(value, ":") == true {
			key_value = strings.Split(value, ":")
			// Checking to see if when we split by ":" if there are more than 2 entries, indicating a multiple split which will mess it up.
			// Want a way to make it so that it uses the first split as the key-value pair and then appends all of the remaining key_values together to be the value
			if len(key_value) > 2 {
				joined_strings := strings.Join(key_value[1:], ":")
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
}
