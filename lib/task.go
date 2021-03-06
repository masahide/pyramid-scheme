package lib

import (
	"errors"
	"fmt"
)

// Job status
const (
	Queued      = iota // 0
	Initalizing        // 1
	Running            // 2
	Finished           // 3
	Failed             // 4
	Cancelled          // 5
)


type Host struct {
	Name       string
	Status     int
	ReturnCode int
	Message    string
}

type Job struct {
	Pcode string
	Hosts []Host
}

type HostList struct {
	Pcode string
	Name  []string
}

type Task struct {
	jobs []Job
}

func (this *Task) GetHosts(id int) ([]Host, error) {
	if len(this.jobs) <= id {
		return nil, errors.New(fmt.Sprintf("ID that does not exist. max id is %+v", len(this.jobs)))
	}
	return this.jobs[id].Hosts, nil
}

func (this *Task) NextHosts(id int, num int) ([]Host, error) {
	if len(this.jobs) <= id {
		return nil, errors.New(fmt.Sprintf("ID that does not exist. max id is %+v", len(this.jobs)))
	}
	hosts := []Host{}
	for index, host := range this.jobs[id].Hosts {
		if host.Status == Queued {
			this.jobs[id].Hosts[index].Status = Initalizing
			hosts = append(hosts, host)
			if len(hosts) >= num {
				break
			}
		}
	}
	return hosts, nil
}

func (this *Task) UpdateHost(jobId int, host Host) error {
	if len(this.jobs) <= jobId {
		return errors.New(fmt.Sprintf("ID that does not exist. max jobId is %+v", len(this.jobs)))
	}
	for hostId, h := range this.jobs[jobId].Hosts {
		if h.Name == host.Name {
			this.jobs[jobId].Hosts[hostId] = host
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Host can not be found is %+v", host.Name))

}

func (this *Task) PostJob(hostList *HostList) int {
	hosts := []Host{}
	for _, name := range hostList.Name {
		hosts = append(hosts, Host{name, Queued, 0, ""})
	}
	job := Job{hostList.Pcode, hosts}
	this.jobs = append(this.jobs, job)
	return len(this.jobs) - 1 //jobId
}
