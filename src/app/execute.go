package app

import (
	"cloudflare-dns-record-updater/cloudflare"
	"cloudflare-dns-record-updater/data"
	"errors"
	"fmt"
)

//Execute function for starting app
func Execute(data *data.Scheme) error {
	currentIP, err := getCurrentIP()
	if err != nil {
		return err
	}

	cf := cloudflare.NewCloudflare(data.XAuthEmail, data.XAuthKey)
	zoneInfo, err := cf.GetZoneInfo(data.Domain)
	if err != nil {
		return err
	}
	zoneID := zoneInfo.ID

	recordInfo, err := cf.GetDNSRecord(zoneID, data.Record, data.Domain)
	if err != nil {
		return err
	}
	fmt.Println("Old IP:", recordInfo.Content)
	fmt.Println("Current IP:", currentIP)
	if recordInfo.Content == currentIP {
		fmt.Println("IP doesn't have to be changed")
		return nil
	}
	recordInfo.Content = currentIP
	recordInfo.Proxied = data.Proxied
	res, err := cf.SetDNSRecord(zoneID, recordInfo)
	if err != nil {
		return err
	}
	if !res.Success {
		return errors.New(res.Errors[0].Message)
	}
	fmt.Println("IP changed")
	return nil
}
