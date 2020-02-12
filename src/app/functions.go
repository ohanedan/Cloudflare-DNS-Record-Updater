package app

import (
	"io/ioutil"
	"net/http"
)

//GetCurrentIP returns current ip
func getCurrentIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ip := string(body)
	return ip, nil
}
