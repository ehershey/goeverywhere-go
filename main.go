package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"ernie.org/goe/cmd"
	"ernie.org/goe/proto"
	"github.com/alecthomas/kingpin/v2"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	grpc "google.golang.org/grpc"

	"github.com/gorilla/mux"
)

const autoupdate_version = 170

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
	handleJobs := serveCommand.Flag("handle-jobs", "Run background jobs (default)").Bool()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case browseCommand.FullCommand():
		cmd.Browse(*nodeId)
	case serveCommand.FullCommand():
		serve(*handleJobs)
	default:
		serve(*handleJobs)
	}

}
func serve(handleJobs bool) {
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	r := mux.NewRouter()
	r.HandleFunc("/", sentryHandler.HandleFunc(index))
	r.Handle("/nodes", sentryHandler.Handle(GetNodesHandlerWithTiming))
	r.Handle("/stats", sentryHandler.Handle(GetStatsHandlerWithTiming))
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

	if handleJobs {
		log.Println("Starting job handler background goroutine")
		go HandleJobs()
	} else {
		log.Println("Skipping job handler background goroutine")
	}
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("config:")
	log.Printf("HTTPListenAddr: %v\n", config.HTTPListenAddr)
	log.Printf("GRPCListenAddr: %v\n", config.GRPCListenAddr)
	log.Printf("DB_Name: %v\n", config.DB_Name)
	log.Printf("MongoDB_Uri: %v\n", config.MongoDB_Uri)

	go func() {
		log.Printf("Starting HTTP listener on: %v", config.HTTPListenAddr) // Run Server
		log.Fatal(http.ListenAndServe(config.HTTPListenAddr, r))
	}()

	log.Printf("Starting GRPC listener on: %v", config.GRPCListenAddr)
	lis, err := net.Listen("tcp", config.GRPCListenAddr)
	if err != nil {
		log.Fatalf("failed to start GRPC listener: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterGOEServer(grpcServer, newServer())
	grpcServer.Serve(lis)
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

func newServer() *gOEServer {
	s := &gOEServer{myContext: context.Background()}
	return s
}

type gOEServer struct {
	proto.UnimplementedGOEServer
	myContext context.Context
}

func (s *gOEServer) GetStats(ctx context.Context, request *proto.StatsRequest) (*proto.StatsResponse, error) {

	response, err := getStats(ctx, request)

	if err != nil {
		wrappedErr := fmt.Errorf("Error calling getStats() in grpc method: %w", err)
		return response, wrappedErr
	}

	return response, nil
}
