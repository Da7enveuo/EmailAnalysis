package WebServer

import ()

func InboxFlaggedEmailsGUI(w http.ResponseWriter, r *http.Request) {
	SetSecurityHeaders(w)
	http.ServeFile()
}
