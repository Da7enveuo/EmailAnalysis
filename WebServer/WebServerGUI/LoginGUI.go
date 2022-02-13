package WebServer

import (
    ""
)

func LoginGUI(w http.ResponseWriter, r *http.Request) {
	SetSecurityHeaders(w)
}


