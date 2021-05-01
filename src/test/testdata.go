package test

import "cloudflare-dns-record-updater/data"

//GetTestData returns test data for tests
func GetTestData() (*data.Scheme, bool) {
	//if empty, tests won't work
	d := data.Scheme{
		XAuthEmail:    "",
		XAuthKey:      "",
		UseBearerAuth: true,
		BearerAuthKey: "",
		Domain:        "",
		Record:        "",
		Proxied:       false,
	}
	if d.Domain == "" || d.Record == "" {
		return nil, false
	}
	if d.UseBearerAuth && d.BearerAuthKey == "" {
		return nil, false
	}
	if !d.UseBearerAuth && (d.XAuthEmail == "" || d.XAuthKey == "") {
		return nil, false
	}
	return &d, true
}
