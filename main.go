package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ernie.org/goe/cmd"
	"github.com/alecthomas/kingpin"
	"github.com/gorilla/mux"
)

const autoupdate_version = 96

var routes []string

// var MongoDB_Uri = kingpin.Flag("mongodb_uri", "MongoDB URI").String()

func main() {
	// parse config to check for errors before doing anything else
	//
	_, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	// app := kingpin.New(os.Args[0], "GO Everywhere backend")
	// command := app.Command("command", "").Default()
	// flag := command.Flag("flag", "").Bool()
	// subcommand := command.Command("subcommand", "").Default()
	// arg := subcommand.Arg("arg", "").Required().String()

	kingpin.Version(version())
	browseCommand := kingpin.Command("browse", "Open a browser browsing the given node id.")
	nodeId := browseCommand.Arg("nodeId", "Node ID to browse to").Required().Int()

	serveCommand := kingpin.Command("serve", "Run backend code.").Default()
	switch kingpin.Parse() {
	case browseCommand.FullCommand():
		cmd.Browse(*nodeId)
	case serveCommand.FullCommand():
		serve()
	default:
		serve()
	}

	log.Fatal("after parse")
}
func serve() {

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.Handle("/nodes", GetNodesHandlerWithTiming)
	r.Handle("/refresh_nodes", RefreshNodesHandlerWithTiming)
	r.Handle("/ignore_nodes", IgnoreNodesHandlerWithTiming)
	r.HandleFunc("/points", GetNodesHandler)
	r.HandleFunc("/bookmarks", GetNodesHandler)
	r.HandleFunc("/echo", echo)
	r.HandleFunc("/kv", KeyValueHandler)
	r.HandleFunc("/version", VersionHandler)

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
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

	go HandleJobs()
	config, err := GetConfig()
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
