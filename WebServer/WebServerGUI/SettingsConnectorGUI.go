package WebServer

import ()

func SettingsConnectorGUI(w http.ResponseWriter, r *http.Request) {
	SetSecurityHeaders(w)
	http.ServeFile("path to settings file")
}
