package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/alecthomas/kingpin"
	"github.com/gorilla/mux"
)

const autoupdate_version = 63

var routes []string

// var MongoDB_Uri = kingpin.Flag("mongodb_uri", "MongoDB URI").String()

// var app = kingpin.New(os.Args[0], "GO Everywhere backend")

func main() {
	// app.GetFlag("help").Short('h')
	// app.GetFlag("version").Short('v')
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	kingpin.Version(version())
	kingpin.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.Handle("/nodes", GetNodesHandlerWithTiming)
	r.Handle("/refresh_nodes", RefreshNodesHandlerWithTiming)
	r.HandleFunc("/points", GetNodesHandler)
	r.HandleFunc("/bookmarks", GetNodesHandler)
	r.HandleFunc("/echo", echo)
	r.HandleFunc("/kv", KeyValueHandler)
	r.HandleFunc("/version", VersionHandler)

	err = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			log.Fatal(err)
		}

		routes = append(routes, pathTemplate)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("config:")
	log.Println("ListenPort: ", config.ListenPort)
	log.Println("DB_Name: ", config.DB_Name)
	log.Println("MongoDB_Uri: ", config.MongoDB_Uri)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.ListenPort), r)) // Run Server

}

func echo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "vars: %v\n", vars)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "routes: %v\n", routes)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, version())
}

func version() string {
	return fmt.Sprintf("%d", autoupdate_version)
}
