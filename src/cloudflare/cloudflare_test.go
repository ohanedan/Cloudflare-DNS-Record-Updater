package cloudflare

import (
	"cloudflare-dns-record-updater/test"
	"testing"
)

func TestCloudflare_sendGetRequest(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	res, err := cf.sendGetRequest("zones?name="+cfg.Domain, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Response: %v", res)
	if !res.Success {
		t.Errorf("Request not return success.")
		t.FailNow()
	}
}

func TestCloudflare_sendGetRequest_WithResponse(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	zoneResult := zoneDataResultScheme{}
	res, err := cf.sendGetRequest("zones?name="+cfg.Domain, &zoneResult)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Response: %+v", zoneResult.Result[0])
	if !res.Success {
		t.Errorf("Request not return success.")
		t.FailNow()
	}
}

func TestCloudflare_sendGetRequest_WrongRequest(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	res, err := cf.sendGetRequest("whatisit", nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Response: %v", res)
	t.Logf("Response Error[0]: %v", res.Errors[0])
	if res.Success {
		t.Errorf("Request return success.")
		t.FailNow()
	}
}

func TestCloudflare_GetZoneInfo(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	info, err := cf.GetZoneInfo(cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("zone[0] ID: %+v", info)
}

func TestCloudflare_GetZoneInfo_WrongDomain(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	_, err := cf.GetZoneInfo("boyledomainmiolur")
	if err == nil {
		t.FailNow()
	}
}

func TestCloudflare_GetDNSRecord(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	info, err := cf.GetZoneInfo(cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	zoneID := info.ID
	record, err := cf.GetDNSRecord(zoneID, cfg.Record, cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("zoneID: %v, record: %+v", zoneID, record)
}

func TestCloudflare_orGetDNSRecd_NotRegisteredDNS(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	info, err := cf.GetZoneInfo(cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	zoneID := info.ID
	_, err = cf.GetDNSRecord(zoneID, cfg.Record+"x", cfg.Domain)
	if err == nil {
		t.FailNow()
	}
}

func TestCloudflare_sendPutRequest(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	zoneInfo, err := cf.GetZoneInfo(cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	zoneID := zoneInfo.ID

	recordInfo, err := cf.GetDNSRecord(zoneID, cfg.Record, cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	oldProxyStatus := recordInfo.Proxied
	recordInfo.Proxied = !recordInfo.Proxied
	res, err := cf.sendPutRequest("zones/"+zoneID+"/dns_records/"+recordInfo.ID, recordInfo)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !res.Success {
		t.Errorf("Request not return success")
		t.FailNow()
	}
	recordInfo, err = cf.GetDNSRecord(zoneID, cfg.Record, cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if recordInfo.Proxied == oldProxyStatus {
		t.Error("proxy status not changed")
		t.FailNow()
	}
}

func TestCloudflare_SetDNSRecord(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	cf := NewCloudflare(cfg.XAuthEmail, cfg.XAuthKey)
	zoneInfo, err := cf.GetZoneInfo(cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	zoneID := zoneInfo.ID

	recordInfo, err := cf.GetDNSRecord(zoneID, cfg.Record, cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	oldProxyStatus := recordInfo.Proxied
	recordInfo.Proxied = !recordInfo.Proxied
	res, err := cf.SetDNSRecord(zoneID, recordInfo)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if !res.Success {
		t.Error("proxy status not changed")
		t.FailNow()
	}
	recordInfo, err = cf.GetDNSRecord(zoneID, cfg.Record, cfg.Domain)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if recordInfo.Proxied == oldProxyStatus {
		t.Error("proxy status not changed")
		t.FailNow()
	}
}
