package Formatting

import (
	"strings"
	"sync"
)

func authentication_analysis_loop(analysis string) (final_string string) {
	final_string = ""
	if analysis != "" {
		f_c := analysis[0:1]
		if f_c == " " {
			final_string = analysis[1:]
		} else {
			final_string = analysis
		}
		l_c := final_string[len(final_string)-1:]
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
