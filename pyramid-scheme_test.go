package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"testing"
)

func TestShowVersion(t *testing.T) {
	if showVersion() != "pyramid-scheme version: -" {
		t.Errorf("showVersion() = [%s], want %v", showVersion(), "pyramid-scheme version: -")
	}
}

func TestHandler(t *testing.T) {

	ps := PyramidScheme{}
	handler := rest.ResourceHandler{
		EnableStatusService: true,
	}
	handler.SetRoutes(
		&rest.Route{"POST", "/jobs", ps.PostJobHandler},
	)

	recorded := test.RunRequest(t, &handler, test.MakeSimpleRequest("GET", "http://1.2.3.4/jobs", nil))
	recorded.CodeIs(405)
	recorded.ContentTypeIsJson()

	payload := map[string]interface{}{}

	err := recorded.DecodeJsonPayload(&payload)

	if err != nil {
		t.Fatal(err)
	}

}
