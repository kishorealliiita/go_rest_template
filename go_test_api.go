package main

import "testing"

func TestEndpoint(t *testing.T) {
	var e Endpoint
	//parse metadata file.
	e.getEndpoint()
	in := e.Version
	out := "1.0"
	if in != out {
		t.Errorf("Version(%v) = %v, want %v", in, out, out)
	}
}

func TestCommitHash(t *testing.T) {
	var e Endpoint
	//parse metadata file.
	e.getEndpoint()

	object := new(conf)
	//get commit id from git repo.
	getJson(e.Git_Repo, &object)
	commit_sha := object.Object.Sha
	in := commit_sha
	out := "b5ded28249898e61c102de923f82eba9e4b41628"
	if in != out {
		t.Errorf("Input(%v) = %v, want %v", in, out, out)
	}
}
