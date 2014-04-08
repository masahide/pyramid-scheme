package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/masahide/pyramid-scheme/lib"
	"log"
	"net/http"
	"strconv"
)


type PyramidScheme struct {
	task    lib.Task
}


/**
 * postjob handler
 * post payload example: { "Pcode" : "hogeproject", "Name" : [ "host01", "host02", "host03", "host04" ] }
 */
// PostJob handler
func (this *PyramidScheme) PostJobHandler(w rest.ResponseWriter, req *rest.Request) {
	hostList := lib.HostList{}
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
	num, err := strconv.Atoi(req.PathParam("num"))
	if err != nil {
		num = 1
	}
	if job, err := this.task.NextHosts(id, num); err == nil {
		w.WriteJson(&job)
	} else {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateHost handler
func (this *PyramidScheme) PutUpdateHostHandler(w rest.ResponseWriter, req *rest.Request) {
	id, _ := strconv.Atoi(req.PathParam("id"))
	host := lib.Host{}
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
	flag.Usage = lib.Usage
	flag.Parse()

	if *lib.Version {
		fmt.Printf("%s\n", lib.ShowVersion())
		return
	}
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"POST", "/jobs", ps.PostJobHandler},
		&rest.Route{"GET", "/jobs/:id/hosts", ps.GetHostsHandler},
		&rest.Route{"PUT", "/jobs/:id/nexthosts/:num", ps.PutNextHostsHandler},
		&rest.Route{"PUT", "/jobs/:id/nexthost", ps.PutNextHostsHandler},
		&rest.Route{"PUT", "/jobs/:id/updatehost/", ps.PutUpdateHostHandler},
	)
	log.Fatal(http.ListenAndServe(":8000", &handler))
}
