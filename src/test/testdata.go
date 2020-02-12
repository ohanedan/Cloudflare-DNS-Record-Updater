package test

import "cloudflare-dns-record-updater/data"

//GetTestData returns test data for tests
func GetTestData() (*data.Scheme, bool) {
	//if empty, tests won't work
	d := data.Scheme{
		XAuthEmail: "",
		XAuthKey:   "",
		Domain:     "",
		Record:     "",
		Proxied:    false,
	}
	if d.XAuthEmail == "" || d.Domain == "" || d.XAuthKey == "" || d.Record == "" {
		return nil, false
	}
	return &d, true
}
