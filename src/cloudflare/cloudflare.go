package cloudflare

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//NewCloudflare creates Cloudflare instance
func NewCloudflare(XAuthEmail string, XAuthKey string) *Cloudflare {
	client := &http.Client{}
	headers := make(map[string]string)
	headers["X-Auth-Email"] = XAuthEmail
	headers["X-Auth-Key"] = XAuthKey
	headers["Content-Type"] = "application/json"
	d := &Cloudflare{
		Client:  client,
		Headers: headers,
		URL:     "https://api.cloudflare.com/client/v4/",
	}
	return d
}

//GetDNSRecord returns info of dns given as parameter
func (cf *Cloudflare) GetDNSRecord(zoneID string, dns string, domain string) (*DNSRecordScheme, error) {
	fullDNS := dns + "." + domain
	dnsResult := dnsRecordResultScheme{}
	res, err := cf.sendGetRequest("zones/"+zoneID+"/dns_records?name="+fullDNS, &dnsResult)

	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errors.New(res.Errors[0].Message)
	}
	if len(dnsResult.Result) == 0 {
		return nil, errors.New("No dns found")
	}
	return dnsResult.Result[0], nil
}

//SetDNSRecord sets dns record using DNSRecordScheme
func (cf *Cloudflare) SetDNSRecord(zoneID string, scheme *DNSRecordScheme) (*Response, error) {
	res, err := cf.sendPutRequest("zones/"+zoneID+"/dns_records/"+scheme.ID, scheme)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errors.New(res.Errors[0].Message)
	}

	return res, nil
}

//GetZoneInfo gets info of zone given as parameter
func (cf *Cloudflare) GetZoneInfo(domain string) (*ZoneDataScheme, error) {
	zoneResult := zoneDataResultScheme{}
	res, err := cf.sendGetRequest("zones?name="+domain, &zoneResult)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errors.New(res.Errors[0].Message)
	}
	if len(zoneResult.Result) == 0 {
		return nil, errors.New("No zone found")
	}
	return zoneResult.Result[0], nil
}

func (cf *Cloudflare) sendGetRequest(url string, resultScheme interface{}) (*Response, error) {
	reqURL := cf.URL + url
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	for header, val := range cf.Headers {
		req.Header.Add(header, val)
	}
	resp, err := cf.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if resultScheme != nil {
		err = json.Unmarshal(body, resultScheme)
		if err != nil {
			return nil, err
		}
	}
	return &response, nil
}

func (cf *Cloudflare) sendPutRequest(url string, scheme interface{}) (*Response, error) {
	jsonData, err := json.Marshal(scheme)
	if err != nil {
		return nil, err
	}

	reqURL := cf.URL + url
	req, err := http.NewRequest("PUT", reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	for header, val := range cf.Headers {
		req.Header.Add(header, val)
	}
	resp, err := cf.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
