package main

import (
	"errors"
	"fmt"
	"github.com/ant0ine/go-json-rest"
	"log"
	"net/http"
	"strconv"
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

const (
	NextHostNum = 3 //
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

type PyramidScheme struct {
	jobs []Job
}

func (this *PyramidScheme) GetHosts(id int) ([]Host, error) {
	if len(this.jobs) <= id {
		return nil, errors.New(fmt.Sprintf("ID that does not exist. max id is %+v", len(this.jobs)))
	}
	return this.jobs[id].Hosts, nil
}

func (this *PyramidScheme) NextHosts(id int) ([]Host, error) {
	if len(this.jobs) <= id {
		return nil, errors.New(fmt.Sprintf("ID that does not exist. max id is %+v", len(this.jobs)))
	}
	hosts := []Host{}
	for index, host := range this.jobs[id].Hosts {
		if host.Status == Queued {
			this.jobs[id].Hosts[index].Status = Initalizing
			hosts = append(hosts, host)
			if len(hosts) >= NextHostNum {
				break
			}
		}
	}
	return hosts, nil
}

func (this *PyramidScheme) UpdateHost(jobId int, host Host) error {
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

func (this *PyramidScheme) PostJob(hostList *HostList) int {
	hosts := []Host{}
	for _, name := range hostList.Name {
		hosts = append(hosts, Host{name, Queued, 0, ""})
	}
	job := Job{hostList.Pcode, hosts}
	this.jobs = append(this.jobs, job)
	return len(this.jobs) - 1
}

// PostJob handler
func (this *PyramidScheme) PostJobHandler(w *rest.ResponseWriter, r *rest.Request) {
	hostList := HostList{}
	err := r.DecodeJsonPayload(&hostList)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(this.PostJob(&hostList))
}

// GetHosts handler
func (this *PyramidScheme) GetHostsHandler(w *rest.ResponseWriter, req *rest.Request) {
	id, _ := strconv.Atoi(req.PathParam("id"))
	if job, err := this.GetHosts(id); err == nil {
		w.WriteJson(&job)
	} else {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NextHosts handler
func (this *PyramidScheme) PutNextHostsHandler(w *rest.ResponseWriter, req *rest.Request) {
	id, _ := strconv.Atoi(req.PathParam("id"))
	if job, err := this.NextHosts(id); err == nil {
		w.WriteJson(&job)
	} else {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	ps := PyramidScheme{}
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		rest.Route{"POST", "/jobs", ps.PostJobHandler},
		rest.Route{"GET", "/jobs/:id/hosts", ps.GetHostsHandler},
		rest.Route{"PUT", "/jobs/:id/nexthosts", ps.PutNextHostsHandler},
	)
	log.Fatal(http.ListenAndServe(":8000", &handler))
}
