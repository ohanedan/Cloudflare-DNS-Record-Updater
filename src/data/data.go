package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

//NewData function gets or creates new Data and initilize it.
func NewData() (*Data, bool, error) {
	d := &Data{}
	filled, err := d.initialize()
	if err != nil {
		return nil, filled, err
	}
	return d, filled, nil
}

/*
//SaveData saves current data to data file.
func (t *Data) SaveData() error {
	json, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(t.Path, json, 0777)
	if err != nil {
		return err
	}
	return nil
}*/

func (t *Data) initialize() (bool, error) {
	usr, err := user.Current()
	if err != nil {
		return false, err
	}
	t.Path = filepath.Join(usr.HomeDir, ".cloudflare-dns-record-updater")
	_, err = os.Stat(t.Path)
	if os.IsNotExist(err) {
		fmt.Println("Creating data file. Path:", t.Path)
		d := &Scheme{
			XAuthEmail: "",
			XAuthKey:   "",
			Domain:     "",
			Record:     "",
			Proxied:    false,
		}
		json, err := json.MarshalIndent(d, "", " ")
		if err != nil {
			return false, err
		}
		err = ioutil.WriteFile(t.Path, json, 0777)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	bytes, err := ioutil.ReadFile(t.Path)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(bytes, &t.Data)
	if err != nil {
		return false, err
	}
	if t.Data.XAuthEmail == "" || t.Data.Domain == "" ||
		t.Data.XAuthKey == "" || t.Data.Record == "" {
		return false, nil
	}
	return true, nil
}
