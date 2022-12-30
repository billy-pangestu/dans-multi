package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	getAll           = "http://dev3.dansmultipro.co.id/api/recruitment/positions.json?"
	getByID          = "http://dev3.dansmultipro.co.id/api/recruitment/positions/"
	getContact       = "/v1/contacts"
	PostaddTempalate = "/v1/configs/templates"
)

func GetAll(page int, location, description, fullTime string) (res interface{}, err error) {

	var addOnVar string
	if description != "" {
		addOnVar += "description=" + description
	}
	if location != "" {
		if addOnVar != "" {
			addOnVar += "&"
		}
		addOnVar += "location=" + location
	}
	if fullTime != "" {
		if addOnVar != "" {
			addOnVar += "&"
		}
		addOnVar += "full_time=" + fullTime
	}

	urlStr := getAll + addOnVar

	fmt.Println(urlStr)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return res, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(body), &res)

	return res, err
}

func GetByID(id string) (res interface{}, err error) {

	urlStr := getByID + id

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return res, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(body), &res)

	return res, err
}
