package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/masahide/pyramid-scheme/lib"
	"github.com/masahide/pyramid-scheme/version"
	"log"
	"net/http"
	"os"
	"strconv"
)

const name = "pyramid-scheme"

type PyramidScheme struct {
	task    task.Task
	version bool
}

func showVersion() string {
	return fmt.Sprintf("%s version: %v-%v", name, version.VERSION, version.GITCOMMIT)
}

func usage() {
	fmt.Printf("%s\n", showVersion())
	fmt.Fprintf(os.Stderr, "usage: %s [flags ...]\n", name)
	flag.PrintDefaults()
	os.Exit(2)
}

// PostJob handler
func (this *PyramidScheme) PostJobHandler(w rest.ResponseWriter, req *rest.Request) {
	hostList := task.HostList{}
	err := req.DecodeJsonPayload(&hostList)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(this.task.PostJob(&hostList))
}

// GetHosts handler
func (this *PyramidScheme) GetHostsHandler(w rest.ResponseWriter, req *rest.Request) {
	id, _ := strconv.Atoi(req.PathParam("id"))
	if job, err := this.task.GetHosts(id); err == nil {
		w.WriteJson(&job)
	} else {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NextHosts handler
func (this *PyramidScheme) PutNextHostsHandler(w rest.ResponseWriter, req *rest.Request) {
	id, _ := strconv.Atoi(req.PathParam("id"))
	if job, err := this.task.NextHosts(id); err == nil {
		w.WriteJson(&job)
	} else {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateHost handler
func (this *PyramidScheme) PutUpdateHostHandler(w rest.ResponseWriter, req *rest.Request) {
	id, _ := strconv.Atoi(req.PathParam("id"))
	host := task.Host{}
	err := req.DecodeJsonPayload(&host)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := this.task.UpdateHost(id, host); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	ps := PyramidScheme{}
	/*
		flag.StringVar(&co.fileName, "f", "config.yml", "config file")
		flag.IntVar(&co.sleepTime, "t", 30, "sleep time(Sec)")
	*/
	flag.BoolVar(&ps.version, "v", false, "show version")
	flag.Usage = usage
	flag.Parse()

	if ps.version {
		fmt.Printf("%s\n", showVersion())
		return
	}
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"POST", "/jobs", ps.PostJobHandler},
		&rest.Route{"GET", "/jobs/:id/hosts", ps.GetHostsHandler},
		&rest.Route{"PUT", "/jobs/:id/nexthosts", ps.PutNextHostsHandler},
		&rest.Route{"PUT", "/jobs/:id/updatehost/", ps.PutUpdateHostHandler},
	)
	log.Fatal(http.ListenAndServe(":8000", &handler))
}
