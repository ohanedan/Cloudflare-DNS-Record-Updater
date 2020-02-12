package cloudflare

import (
	"net/http"
)

//Cloudflare contains API functions.
type Cloudflare struct {
	Client  *http.Client
	Headers map[string]string
	URL     string
}

//Response is the API response scheme
type Response struct {
	Success bool     `json:"success"`
	Errors  []*Error `json:"errors"`
}

//Error is the API response scheme for errors
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type zoneDataResultScheme struct {
	Result []*ZoneDataScheme `json:"result"`
}

//ZoneDataScheme is the API response scheme for zones
type ZoneDataScheme struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	CreatedOn           string   `json:"created_on"`
	ModifiedOn          string   `json:"modified_on"`
	ActivatedOn         string   `json:"activated_on"`
	OriginalNameservers []string `json:"original_name_servers"`
	Nameservers         []string `json:"name_servers"`
	Status              string   `json:"status"`
}

type dnsRecordResultScheme struct {
	Result []*DNSRecordScheme `json:"result"`
}

//DNSRecordScheme is the API response scheme for dns records
type DNSRecordScheme struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Proxied bool   `json:"proxied"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	TTL     int    `json:"ttl"`
}
