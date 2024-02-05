package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"ernie.org/goe/cmd"
	"github.com/alecthomas/kingpin/v2"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"

	"github.com/gorilla/mux"
)

const autoupdate_version = 123

var routes []string

// var MongoDB_Uri = kingpin.Flag("mongodb_uri", "MongoDB URI").String()

func main() {
	// parse config to check for errors before doing anything else
	//
	_, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Debug:              false,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	app := kingpin.New(os.Args[0], "GO Everywhere backend")

	app.Version(version())
	app.HelpFlag.Short('h')
	browseCommand := app.Command("browse", "Open a browser browsing the given node id.")
	nodeId := browseCommand.Arg("nodeId", "Node ID to browse to").Required().Int()

	serveCommand := app.Command("serve", "Run backend code.").Default()
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case browseCommand.FullCommand():
		cmd.Browse(*nodeId)
	case serveCommand.FullCommand():
		serve()
	default:
		serve()
	}

}
func serve() {
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	r := mux.NewRouter()
	r.HandleFunc("/", sentryHandler.HandleFunc(index))
	r.Handle("/nodes", sentryHandler.Handle(GetNodesHandlerWithTiming))
	r.Handle("/refresh_nodes", sentryHandler.Handle(RefreshNodesHandlerWithTiming))
	r.Handle("/ignore_nodes", sentryHandler.Handle(IgnoreNodesHandlerWithTiming))
	r.HandleFunc("/points", sentryHandler.HandleFunc(GetNodesHandler))
	r.HandleFunc("/bookmarks", sentryHandler.HandleFunc(GetNodesHandler))
	r.HandleFunc("/echo", sentryHandler.HandleFunc(echo))
	r.HandleFunc("/kv", sentryHandler.HandleFunc(KeyValueHandler))
	r.HandleFunc("/version", sentryHandler.HandleFunc(VersionHandler))

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
