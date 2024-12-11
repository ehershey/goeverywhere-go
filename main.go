package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// generated by protoc-gen-connect-go

	connect "connectrpc.com/connect"
	"ernie.org/goe/cmd"
	"ernie.org/goe/proto"
	"ernie.org/goe/proto/protoconnect"
	"github.com/alecthomas/kingpin/v2"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// grpc-connect " "google.golang.org/grpc"

	"github.com/gorilla/mux"
)

const autoupdate_version = 357

const GRACEFUL_SHUTDOWN_TIMEOUT_SECS = 10
const WRITE_TIMEOUT_SECS = 10
const READ_TIMEOUT_SECS = 10
const IDLE_TIMEOUT_SECS = 20

var routes []string

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
	versionCommand := app.Command("version", "Display program version.")
	short := versionCommand.Flag("short", "Use short format").Default("false").Bool()
	handleJobs := serveCommand.Flag("handle-jobs", "Run background jobs (default)").Default("true").Bool()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case browseCommand.FullCommand():
		if err = cmd.Browse(*nodeId); err != nil {
			log.Fatal(err)
		}
	case serveCommand.FullCommand():
		serve(*handleJobs)
	case versionCommand.FullCommand():
		if *short {
			println(shortVersion())
		} else {
			println(version())
		}
	default:
		serve(*handleJobs)
	}

}
func serve(handleJobs bool) {
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	GRPCPath, GRPCHandler := protoconnect.NewGOEServiceHandler(newServer())
	log.Printf("GRPCPath: %v\n", GRPCPath)

	r := mux.NewRouter()
	r.Handle("/GOEService/{method}", sentryHandler.Handle(GRPCHandler))
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
	r.PathPrefix("/").HandlerFunc(catchAllHandler)

	// store routes array for index requests to /
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("pathTemplate: %v\n", pathTemplate)

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
	log.Printf("Version: %d\n", autoupdate_version)
	log.Println("config:")
	log.Printf("HTTPListenAddr: %v\n", config.HTTPListenAddr)
	log.Printf("GRPCListenAddr: %v\n", config.GRPCListenAddr)
	log.Printf("DB_Name: %v\n", config.DB_Name)
	log.Printf("MongoDB_Uri: %v\n", config.MongoDB_Uri)
	log.Printf("Strava_MongoDB_Uri: %v\n", config.Strava_MongoDB_Uri)
	log.Printf("Strava_DB_Name: %v\n", config.Strava_DB_Name)

	srv := &http.Server{
		Addr:         config.HTTPListenAddr,
		Handler:      h2c.NewHandler(r, &http2.Server{}),
		WriteTimeout: time.Second * WRITE_TIMEOUT_SECS,
		ReadTimeout:  time.Second * READ_TIMEOUT_SECS,
		IdleTimeout:  time.Second * IDLE_TIMEOUT_SECS,
	}
	//h2chandler := h2c.NewHandler(r, srv)
	shuttingDownGracefully := false
	go func() {
		log.Printf("Starting HTTP listener on: %v\n", config.HTTPListenAddr)
		err = srv.ListenAndServe()
		if err != nil && !shuttingDownGracefully {
			log.Fatalf("failed to start HTTP listener: %v", err)
		}
	}()

	grpc_mux := http.NewServeMux()
	//path, handler := protoconnect.NewGOEServiceHandler(newServer())
	//log.Printf("path: %v\n", path)
	grpc_mux.Handle(GRPCPath, GRPCHandler)

	go func() {
		log.Printf("Starting GRPC listener on: %v\n", config.GRPCListenAddr)
		err = http.ListenAndServe(
			config.GRPCListenAddr, // Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(grpc_mux, &http2.Server{}),
		)
		if err != nil {
			log.Fatalf("failed to start GRPC listener: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C), or SIGTERM
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c

	shuttingDownGracefully = true

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), GRACEFUL_SHUTDOWN_TIMEOUT_SECS*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("failed to shutdown HTTP server: %v", err)
	}
	log.Println("shutting down")
	os.Exit(0)
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
	fmt.Fprint(w, shortVersion())
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	fmt.Printf("Not Found: %v\n", r.URL)
	fmt.Fprintf(w, "Not Found\n")
}

func version() string {
	return fmt.Sprintf("main: %d\nBuild time: %v\nGit commit: %v\nGo Version: %v", autoupdate_version, BuildTime, CommitHash, GoVersion)
}
func shortVersion() string {
	return fmt.Sprintf("%d.%v.%v", autoupdate_version, BuildTime, CommitHash)
}

// build flags
var BuildTime = "Unspecified"

var CommitHash = "Unspecified"
var GoVersion = "Unspecified"

func newServer() *gOEServiceServer {
	s := &gOEServiceServer{myContext: context.Background()}
	return s
}

type gOEServiceServer struct {
	//proto.UnimplementedGOEServiceServer
	myContext context.Context
}

func (s *gOEServiceServer) GetPoints(
	ctx context.Context,
	req *connect.Request[proto.GetPointsRequest],
	stream *connect.ServerStream[proto.GetPointsResponse]) error {
	log.Println("GetPoints request headers: ", req.Header())
	return fmt.Errorf("Unimplemented")
}

func (s *gOEServiceServer) GetBookmarks(
	ctx context.Context,
	req *connect.Request[proto.GetBookmarksRequest],
	stream *connect.ServerStream[proto.GetBookmarksResponse]) error {
	log.Println("GetBookmarks request headers: ", req.Header())
	return fmt.Errorf("Unimplemented")
}

func (s *gOEServiceServer) SaveBookmark(
	ctx context.Context,
	req *connect.Request[proto.SaveBookmarkRequest]) (*connect.Response[proto.SaveBookmarkResponse], error) {
	log.Println("SaveBookmark request headers: ", req.Header())
	return nil, fmt.Errorf("Unimplemented")
}

func (s *gOEServiceServer) GetPolylines(
	ctx context.Context,
	req *connect.Request[proto.GetPolylinesRequest],
	stream *connect.ServerStream[proto.GetPolylinesResponse]) error {
	log.Println("GetPolylines request headers: ", req.Header())

	for line := range getPolylines(ctx, req.Msg) {

		if err := stream.Send(line); err != nil {
			wrappedErr := fmt.Errorf("Error sending res to stream: %w", err)
			return wrappedErr
		}
		// to test streaming
		//time.Sleep(200 * time.Millisecond)
	}
	log.Println("GetPolylines finished streaming results")
	return nil
}

func (s *gOEServiceServer) GetStats(
	ctx context.Context,
	req *connect.Request[proto.GetStatsRequest],
) (*connect.Response[proto.GetStatsResponse], error) {
	log.Println("GetStats request headers: ", req.Header())

	response, err := getStats(ctx, req.Msg)

	res := connect.NewResponse(response)

	if err != nil {
		wrappedErr := fmt.Errorf("Error calling getStats() in grpc method: %w", err)
		return res, wrappedErr
	}

	return res, nil
}

func (s *gOEServiceServer) SavePosition(
	ctx context.Context,
	req *connect.Request[proto.SavePositionRequest],
) (*connect.Response[proto.SavePositionResponse], error) {
	log.Println("SavePosition request headers: ", req.Header())

	entry_source := req.Header().Get("X-Forwarded-Host")

	// TODO: sanitize this header - use allow list and verify multiple values still works
	//
	response, err := savePosition(ctx, entry_source, req.Msg)

	if err != nil {
		wrappedErr := fmt.Errorf("Error calling savePosition() in grpc method: %w", err)
		return nil, wrappedErr
	}

	res := connect.NewResponse(response.raw)

	return res, nil
}
