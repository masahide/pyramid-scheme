package main

import (
	"testing"
)

func TestPostJob(t *testing.T) {
	hosts := []string{"a","b","c","d","e","f","g"}
	var hostlist = HostList{"p_code", hosts}
	var ps = PyramidScheme{}

	if jobId := ps.PostJob(&hostlist); jobId != 0 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostlist, jobId, 0)
	}
	if jobId := ps.PostJob(&hostlist); jobId != 1 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostlist, jobId, 1)
	}
}

func TestGetHosts(t *testing.T) {
	hosts := []string{"a","b","c","d","e","f","g"}
	var hostlist = HostList{"p_code", hosts}
	var ps = PyramidScheme{}
	jobId := ps.PostJob(&hostlist)
	result,_ := ps.GetHosts(jobId)
	if result[0].Name != hosts[0]  {
		t.Errorf("GetHosts(%v) = %+v, want %v", jobId, result, hosts[0])
	}
}

/*
func TestNextHosts(t *testing.T) {
	const hosts = []string{"a","b","c","d","e","f","g"}
	hostlist := HostList{"p_code", hosts}
	ps := PyramidScheme{}
	jobId := ps.PostJob(&hostlist)

	if ; job.Hosts[0].Name != host1 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostlist, job, host1)
	}
	if job := ps.PostJob(&hostlist); job.Hosts[1].Name != host2 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostlist, job, host2)
	}
}
*/
