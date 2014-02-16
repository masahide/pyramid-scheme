package main

import (
	"fmt"
	//"io/ioutil"
	//"os"
	"errors"
	"github.com/ant0ine/go-json-rest"
	"log"
	"net/http"
	"strconv"
)

type Host struct {
	Name    string
	Status  int
	Result  int
	Message string
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

func (this *PyramidScheme) GetHosts(id int) (*[]Host, error) {
	if len(this.jobs) > id {
		return &this.jobs[id].Hosts, nil
	}
	return nil, errors.New(fmt.Sprintf("ID that does not exist. max id is %+v", len(this.jobs)))
}

func (this *PyramidScheme) GetHostsRest(w *rest.ResponseWriter, req *rest.Request) {
	id, _ := strconv.Atoi(req.PathParam("id"))
	if job, err := this.GetHosts(id); err == nil {
		w.WriteJson(job)
	} else {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (this *PyramidScheme) PostJob(hostlist *HostList) *Job {
	hosts := []Host{}
	for _, name := range hostlist.Name {
		hosts = append(hosts, Host{name, 0, 0, ""})
	}
	job := Job{hostlist.Pcode, hosts}
	this.jobs = append(this.jobs, job)
	return &job
}

func (this *PyramidScheme) PostJobRest(w *rest.ResponseWriter, r *rest.Request) {
	hostlist := HostList{}
	err := r.DecodeJsonPayload(&hostlist)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var job = this.PostJob(&hostlist)
	w.WriteJson(job)
}

func main() {
	var ps = PyramidScheme{}
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		rest.Route{"POST", "/jobs", ps.PostJobRest},
		rest.Route{"GET", "/jobs/:id/hosts", ps.GetHostsRest},
	)
	log.Fatal(http.ListenAndServe(":8000", &handler))
}
