package data

//Data contains informations of data and functions
type Data struct {
	Data *Scheme
	Path string
}

//Scheme is the scheme of the data type
type Scheme struct {
	XAuthEmail    string `json:"x_auth_email"`
	XAuthKey      string `json:"x_auth_key"`
	Domain        string `json:"domain"`
	Record        string `json:"record"`
	Proxied       bool   `json:"proxied"`
	UseBearerAuth bool   `json:"use_bearer_auth"`
	BearerAuthKey string `json:"bearer_auth_key"`
}
