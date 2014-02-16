package main

import (
	"testing"
)

func TestGetHostsAndPostJob(t *testing.T) {
	const host1 = "a"
	const host2 = "b"
	var hostlist = HostList{"p_code", []string{host1,host2}}
	var ps = PyramidScheme{}
	
	if job := ps.PostJob(&hostlist); job.Hosts[0].Name != host1 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostlist, job, host1)
	}
	if job := ps.PostJob(&hostlist); job.Hosts[1].Name != host2 {
		t.Errorf("GetHosts(%v) = %+v, want %v", hostlist, job, host2)
	}
}

