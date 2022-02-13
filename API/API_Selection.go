package API

type API_Select struct {
	Service string
	Account string
	Perms   Permissions
}

type Permissions struct {
	Signature string
}

type (P *Permissions)Permissions_Checker interface {
	Perm_Check() bool
}

func (P *Permissions)Perm_Check(){
	// Check that the Permissions Signature is valid before proceeding.
}
// This will take in the request from the frontend that sets the api Service the user wants to run. It will then take in keys to run it properly. 
func Api_Selection(r *http.Request) {
	// will take in json from frontend server on what options the user wants.
	apis := API_Select{r.Body}
	perms := Permission{apis.Perms}
	perms.Perm_Check()
}
