package WebServer

import ()

func SettingsGUI(w http.ResponseWriter, r *http.Request) {
	SetSecurityHeaders(w)
}
