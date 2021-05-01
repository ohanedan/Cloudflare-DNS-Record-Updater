package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	t.Path = filepath.Join("data.json")
	_, err := os.Stat(t.Path)
	if os.IsNotExist(err) {
		fmt.Println("Creating data file.", t.Path)
		d := &Scheme{
			XAuthEmail:    "",
			XAuthKey:      "",
			Domain:        "",
			Record:        "",
			UseBearerAuth: false,
			BearerAuthKey: "",
			Proxied:       false,
		}
		json, err := json.MarshalIndent(d, "", " ")
		if err != nil {
			return false, err
		}
		err = ioutil.WriteFile(t.Path, json, 0777)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	bytes, err := ioutil.ReadFile(t.Path)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(bytes, &t.Data)
	if err != nil {
		return false, err
	}
	if t.Data.Domain == "" || t.Data.Record == "" {
		return false, nil
	}
	if t.Data.UseBearerAuth && t.Data.BearerAuthKey == "" {
		return false, nil
	}
	if !t.Data.UseBearerAuth && (t.Data.XAuthEmail == "" || t.Data.XAuthKey == "") {
		return false, nil
	}
	return true, nil
}
