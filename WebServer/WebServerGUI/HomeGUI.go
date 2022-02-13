package WebServer

import ()

func HomeGUI(w http.ResponseWriter, r *http.Request) {
	// Check if User is authorized, if not then redirect to Login
	SetSecurityHeaders(w)
}
