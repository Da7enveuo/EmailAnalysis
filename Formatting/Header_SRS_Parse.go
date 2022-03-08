package Formatting

import (
	"strings"
	"sync"
)

func sender_recipient_sub(srsj *Sender_Recipient_Subject_Json, wg *sync.WaitGroup) {
	defer wg.Done()
	//recipients_string with srsj.Rec_name
	if strings.Contains(srsj.Rec_name, "<") == true {
		srsj.Rec_email, srsj.Rec_name = strip_email_addr(srsj.Rec_email, srsj.Rec_name)
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
		srsj.Sender_email, srsj.Sender_name = strip_email_addr(srsj.Sender_email, srsj.Sender_name)
	} else {
		srsj.Sender_email = strings.Replace(strings.Replace(srsj.Sender_name, "\"", "", -1), "\r", "", -1)
		srsj.Sender_name = "Not Found"
	}
	// we will do nothing with this and just assign it outright
}

func strip_email_addr(email_addr string, name string) (string, string) {
	rec := strings.Split(email_addr, "<")
	first_character_email := rec[1][0:1]
	if first_character_email == " " {
		email_addr = strings.Replace(strings.Trim(strings.Replace(rec[1][1:], ">", "", -1), "\""), "\r", "", -1)
	} else {
		email_addr = strings.Replace(strings.Trim(strings.Replace(rec[1], ">", "", -1), "\""), "\r", "", -1)
	}
	if strings.Contains(email_addr, " ") == true {
		email_addr = strings.Replace(strings.Replace(string(strings.Split(string(email_addr), " ")[0]), ",", "", -1), "\r", "", -1)
	}
	first_letter := rec[0][0:1]
	if first_letter == " " {
		name = strings.Replace(strings.Replace(rec[0][1:], "\"", "", -1), "\r", "", -1)
	} else {
		name = strings.Replace(strings.Replace(rec[0], "\"", "", -1), "\r", "", -1)
	}
	return email_addr, name
}
