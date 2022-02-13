package WebServer

import (
	"log"
	"net/http"
)

func Server() {

	mux := http.NewServeMux()

	// GUI Endpoints
	mux.HandleFunc("/", HomeGUI)
	mux.HandleFunc("/login", LoginGUI)
	mux.HandleFunc("/settings", SettingsGUI)
	mux.HandleFunc("/settings/connector", SettingsConnectorGUI)
	mux.HandleFunc("/settings/user-config", SettingsUserConfigGUI)
	mux.HandleFunc("/inbox/FlaggedEmails", InboxFlaggedEmailsGUI)

	// API Endpoints, may want to run the api on a different thread and port
	mux.HandleFunc("/api/settings/connector", SettingsConnectorAPI)
	mux.HandleFunc("/api/login", LoginAPI)
	mux.HandleFunc("/api/settings", SettingsAPI)
	mux.HandleFunc("/api/settings/connector", SettingsConnectorAPI)
	mux.HandleFunc("/api/settings/user-config", SettingsUserConfigAPI)
	mux.HandleFunc("/api/inbox/FlaggedEmails", InboxFlaggedEmailsAPI)

	//http.ListenAndServe uses the default server structure.
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
