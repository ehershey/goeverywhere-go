ifeq ($(filter grouped-target,$(value .FEATURES)),)
$(error Make too old - try gmake?)
endif

.PHONY: release test deploy run clean showvars gen

BUILD_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")
COMMIT_HASH = $(shell git log -1 --pretty=format:%h)
GO_VERSION = $(shell go version | cut -f3- -d\ )
GOPATH = $(shell go env GOPATH)
MODULE_VERSION = $(shell git describe --tags)

PROTO_PATH = ../goeverywhere/grpcify/
PROTO_PRETAG_PATH = ./proto_pretag/
PROTO_GEN_PATH = ./proto/

PROTO_FILES = $(wildcard $(PROTO_PATH)*.proto)

GEN_FILES = $(patsubst $(PROTO_PATH)%.proto,$(PROTO_GEN_PATH)%.pb.go,$(PROTO_FILES))
#GEN_FILES = proto/stats.pb.go proto/polylines.pb.go proto/points.pb.go proto/bookmarks.pb.go proto/save_position.pb.go proto/nodes.pb.go proto/goe_service.pb.go proto/protoconnect/goe_service.connect.go
GEN_PRETAG_FILES = $(patsubst $(PROTO_PATH)%.proto,$(PROTO_PRETAG_PATH)%.pb.go,$(PROTO_FILES))
#GEN_PRETAG_FILES = proto_pretag/stats.pb.go proto_pretag/polylines.pb.go proto_pretag/points.pb.go proto_pretag/bookmarks.pb.go proto_pretag/save_position.pb.go proto_pretag/nodes.pb.go proto_pretag/goe_service.pb.go

LDFLAGS = -X \"main.BuildTime=$(BUILD_TIME)\"
LDFLAGS += -X \"main.CommitHash=$(COMMIT_HASH)\"
LDFLAGS += -X \"main.GoVersion=$(GO_VERSION)\"
LDFLAGS += -X \"main.ModuleVersion=$(MODULE_VERSION)\"
# from go tool link  -help:
# -s    disable symbol table
# -w    disable DWARF generation
LDFLAGS += -s -w
GOFLAGS = -tags BuildArgsIncluded

goe: *.go */*.go $(GEN_FILES) Makefile
	go build -ldflags "$(LDFLAGS)" $(GOFLAGS)

run: goe
	./goe $(GOE_RUN_FLAGS)

release: goe.linux.arm64

goe.linux.arm64: test.success goe
	GOOS=linux GOARCH=arm64 go build -o goe.linux.arm64 -ldflags "$(LDFLAGS)" $(GOFLAGS)

test: *.go db.created
	go test -v ./... $(GOFLAGS) && scripts/verify_no_extra_output.sh && touch test.success

vet:
	go vet $(GOFLAGS) ./...

test.success: test

db.created: scripts/loadData.js scripts/setup_empty_db.sh
	scripts/setup_empty_db.sh

deploy: goe.linux.arm64 goe
	echo latest version:
	./goe --version
	echo Checking deployed version and backing up if present:
	ssh goe@oci1 "if test -e ./goe ; then ./goe --version ; cp -pr ./goe ./goe.`date +%s` ; else echo none; fi"
	echo SCPing new build:
	scp goe.linux.arm64 goe@oci1:goe-new
	echo Copying new build into place:
	ssh -t goe@oci1 "mv -f ./goe ./goe.last; mv ./goe-new ./goe && chcon unconfined_u:object_r:bin_t:s0 ./goe && sudo /bin/systemctl restart goe.service"

install: goe
	sudo install ./goe /usr/local/bin/goe

$(GEN_PRETAG_FILES) &: $(PROTO_FILES) /opt/homebrew/bin/protoc-gen-go
	protoc -I$(wildcard $(GOPATH)/pkg/mod/github.com/srikrsna/protoc-gen-gotag*) --go_out=./proto_pretag --go_opt=paths=source_relative $(PROTO_FILES) --proto_path $(PROTO_PATH)

gen: $(GEN_FILES)

$(GEN_FILES) &: $(GEN_PRETAG_FILES) $(GOPATH)/bin/protoc-gen-connect-go $(GOPATH)/bin/protoc-gen-gotag
	cp -pr $(GEN_PRETAG_FILES) ./proto/
	# protoc -I$(wildcard $(GOPATH)/pkg/mod/github.com/srikrsna/protoc-gen-gotag*) --gotag_out=outdir=./proto:. --gotag_opt=paths=source_relative --connect-go_out=./proto --connect-go_opt=paths=source_relative $(PROTO_FILES) --proto_path $(PROTO_PATH)
	protoc --plugin=../protoc-gen-gotag/protoc-gen-gotag -I$(wildcard $(GOPATH)/pkg/mod/github.com/srikrsna/protoc-gen-gotag*) --gotag_out=outdir=./proto:. --gotag_opt=paths=source_relative --connect-go_out=./proto --connect-go_opt=paths=source_relative $(PROTO_FILES) --proto_path $(PROTO_PATH)

$(GOPATH)/bin/protoc-gen-connect-go:
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

$(GOPATH)/bin/protoc-gen-gotag:
	go install github.com/srikrsna/protoc-gen-gotag@latest

/opt/homebrew/bin/protoc-gen-go:
	brew install protoc-gen-go

/opt/homebrew/bin/protoc-gen-go-grpc:
	brew install protoc-gen-go-grpc

clean:
	rm goe

showvars:
	$(foreach v, $(.VARIABLES), $(info $(v) = $($(v))))

