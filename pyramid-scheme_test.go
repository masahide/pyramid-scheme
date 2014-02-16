package main

import (
	"testing"
)

func TestPostJob(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var ps = PyramidScheme{}

	if jobId := ps.PostJob(&hostList); jobId != 0 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostList, jobId, 0)
	}
	if jobId := ps.PostJob(&hostList); jobId != 1 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostList, jobId, 1)
	}
}

func TestGetHosts(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var ps = PyramidScheme{}
	jobId := ps.PostJob(&hostList)
	result, _ := ps.GetHosts(jobId)
	if result[0].Name != allHosts[0] {
		t.Errorf("GetHosts(%v) = %+v, want %v", jobId, result, allHosts[0])
	}
}

func TestNextHosts(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var ps = PyramidScheme{}
	jobId := ps.PostJob(&hostList)
	result, _ := ps.NextHosts(jobId)

	if result[0].Name != allHosts[0] {
		t.Errorf("GetHosts(%v) = %+v, want %v", jobId, result, allHosts[0])
	}
}
