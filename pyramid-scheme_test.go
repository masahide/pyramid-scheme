package main

import (
	"testing"
)

func TestPostJob(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var ps = PyramidScheme{}

	if jobId := ps.PostJob(&hostList); jobId != 0 {
		t.Errorf("PostJob(%v) = %+v, want %v", hostList, jobId, 0)
	}
	if jobId := ps.PostJob(&hostList); jobId != 1 {
		t.Errorf("PostJob(%v) = %+v, want %v", hostList, jobId, 1)
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
		t.Errorf("NextHosts(%v) = %+v, want %v", jobId, result, allHosts[0])
	}
	if len(result) != NextHostNum {
		t.Errorf("len(NextHosts(%v)) = %+v, want %v", jobId, len(result), NextHostNum)
	}
	hosts, _ := ps.GetHosts(jobId)
	i := 0
	for _, host := range hosts {
		if host.Status == Initalizing {
			i++
		}
	}
	if i != NextHostNum {
		t.Errorf("i = %v, want %v", i, NextHostNum)
		t.Errorf("hosts=%v", hosts)
	}

	result, _ = ps.NextHosts(jobId)
	result, _ = ps.NextHosts(jobId)
	if len(result) != 1 {
		t.Errorf("len(NextHosts(%v)) = %+v, want %v", jobId, len(result), 1)
	}
	hosts, _ = ps.GetHosts(jobId)
	i = 0
	for _, host := range hosts {
		if host.Status == Initalizing {
			i++
		}
	}
	if i != 7 {
		t.Errorf("i = %v, want %v", i, 7)
		t.Errorf("hosts=%v", hosts)
	}
	result, _ = ps.NextHosts(jobId)
	if len(result) != 0 {
		t.Errorf("len(NextHosts(%v)) = %+v, want %v", jobId, len(result), 0)
	}

}
