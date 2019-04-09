package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//this struct is used to read metadata json file.
type Endpoint struct {
	Version     string `json:"version"`
	Description string `json:"description"`
	Git_Repo    string `json:"git_url"`
}

//method to read json file.
func (e *Endpoint) getEndpoint() *Endpoint {

	jsonFile, err := ioutil.ReadFile("metadata.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonFile, e)
	if err != nil {
		panic(err)
	}

	return e
}

//these structs are used to get commit sha from github url.
type Object struct {
	Sha  string `json:"sha"`
	Type string `json:"type"`
}

type conf struct {
	Ref    string `json:"ref"`
	url    string `json:"url"`
	Object struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
	} `json:"object"`
}

//method to json output from http GET.
func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
