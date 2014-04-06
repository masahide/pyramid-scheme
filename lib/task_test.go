package task

import (
	"testing"
)

func TestPostJob(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var task = Task{}

	if jobId := task.PostJob(&hostList); jobId != 0 {
		t.Errorf("PostJob(%v) = %+v, want %v", hostList, jobId, 0)
	}
	if jobId := task.PostJob(&hostList); jobId != 1 {
		t.Errorf("PostJob(%v) = %+v, want %v", hostList, jobId, 1)
	}
}

func TestGetHosts(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var task = Task{}
	jobId := task.PostJob(&hostList)
	result, _ := task.GetHosts(jobId)
	if result[0].Name != allHosts[0] {
		t.Errorf("GetHosts(%v) = %+v, want %v", jobId, result, allHosts[0])
	}
}

func TestNextHosts(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var task = Task{}
	jobId := task.PostJob(&hostList)
	result, _ := task.NextHosts(jobId)

	if result[0].Name != allHosts[0] {
		t.Errorf("NextHosts(%v) = %+v, want %v", jobId, result, allHosts[0])
	}
	if len(result) != NextHostNum {
		t.Errorf("len(NextHosts(%v)) = %+v, want %v", jobId, len(result), NextHostNum)
	}
	hosts, _ := task.GetHosts(jobId)
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

	result, _ = task.NextHosts(jobId)
	result, _ = task.NextHosts(jobId)
	if len(result) != 1 {
		t.Errorf("len(NextHosts(%v)) = %+v, want %v", jobId, len(result), 1)
	}
	hosts, _ = task.GetHosts(jobId)
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
	result, _ = task.NextHosts(jobId)
	if len(result) != 0 {
		t.Errorf("len(NextHosts(%v)) = %+v, want %v", jobId, len(result), 0)
	}

}

func TestUpdateHosts(t *testing.T) {
	allHosts := []string{"a", "b", "c", "d", "e", "f", "g"}
	var hostList = HostList{"p_code", allHosts}
	var task = Task{}
	jobId := task.PostJob(&hostList)
	result, _ := task.NextHosts(jobId)
	result[0].Status = Finished
	result[0].ReturnCode = 3
	result[0].Message = "hoge"

	if err := task.UpdateHost(jobId, result[0]); err != nil {
		t.Errorf("UpdateHost(%v,%+v) = %+v, want %v", jobId, result[0], err, 0)
		t.Errorf("allHosts=%v", allHosts)
	}
	hosts, _ := task.GetHosts(jobId)
	if hosts[0].ReturnCode != 3 {
		t.Errorf("GetHosts(%v) = %+v, want %v", jobId, hosts[0].ReturnCode, 3)
		t.Errorf("hosts=%v", hosts)
	}
}
