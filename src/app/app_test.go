package app

import (
	"cloudflare-dns-record-updater/test"
	"regexp"
	"testing"
)

func Test_getCurrentIP(t *testing.T) {
	ip, err := getCurrentIP()
	if err != nil {
		t.Error(err)
	}
	regex := "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	matched, err := regexp.MatchString(regex, ip)
	if err != nil {
		t.Error(err)
	}
	if !matched {
		t.Errorf("Regex not matched. IP: %v", ip)
	}
	t.Logf("ip: %v", ip)
}

func Test_Execute(t *testing.T) {
	cfg, ok := test.GetTestData()
	if !ok {
		t.Log("TestConfig missing.")
		t.FailNow()
	}
	err := Execute(cfg)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
