package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func EndpointRes(w http.ResponseWriter, r *http.Request) {
	var e Endpoint
	//parse metadata file.
	e.getEndpoint()

	object := new(conf)
	//get commit id from git repo.
	getJson(e.Git_Repo, &object)
	commit_sha := object.Object.Sha

	//construct response.
	response := map[string]string{"version": e.Version, "description": e.Description, "lastcommitsha": commit_sha}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
