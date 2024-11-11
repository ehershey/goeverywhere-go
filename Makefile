ifeq ($(filter grouped-target,$(value .FEATURES)),)
$(error Make too old - try gmake?)
endif

.PHONY: release test deploy run clean

BUILD_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")
COMMIT_HASH = $(shell git log -1 --pretty=format:%h)
GO_VERSION = $(shell go version | cut -f3- -d\ )
GOPATH = $(shell go env GOPATH)

GEN_FILES = proto/stats.pb.go proto/polylines.pb.go proto/points.pb.go proto/save_position.pb.go proto/nodes.pb.go proto/goe_service.pb.go proto/protoconnect/goe_service.connect.go
GEN_PRETAG_FILES = proto_pretag/stats.pb.go proto_pretag/polylines.pb.go proto_pretag/points.pb.go proto_pretag/save_position.pb.go proto_pretag/nodes.pb.go proto_pretag/goe_service.pb.go
PROTO_FILES = $(wildcard $(PROTO_PATH)*.proto)
PROTO_PATH = ../goeverywhere/grpcify/

FLAGS = -X \"main.BuildTime=$(BUILD_TIME)\"
FLAGS += -X \"main.CommitHash=$(COMMIT_HASH)\"
FLAGS += -X \"main.GoVersion=$(GO_VERSION)\"

goe: *.go */*.go $(GEN_FILES)
	go build -ldflags "$(FLAGS)"

run: goe
	./goe $(GOE_RUN_FLAGS)

release: goe.linux.arm64 goe.linux.amd64

goe.linux.arm64: test.success goe
	GOOS=linux GOARCH=arm64 go build -o goe.linux.arm64 -ldflags "$(FLAGS)"

goe.linux.amd64: test.success goe
	GOOS=linux GOARCH=amd64 go build -o goe.linux.amd64 -ldflags "$(FLAGS)"

test: *.go db.created
	go test -v && scripts/verify_no_extra_output.sh && touch test.success

test.success: test

db.created: scripts/loadData.js scripts/setup_empty_db.sh
	scripts/setup_empty_db.sh

deploy: goe.linux.arm64 goe.linux.amd64 goe
	echo latest version:
	./goe --version
	echo deployed version:
	ssh oci1 if test -e /usr/local/bin/goe \; then /usr/local/bin/goe --version \; fi
	ssh eahdroplet4 if test -e /usr/local/bin/goe \; then /usr/local/bin/goe --version \; fi
	scp goe.linux.arm64 goe.service oci1:
	scp goe.linux.amd64 goe.service eahdroplet4:
	# if no previous version, copy new version into place so remaining commands will work
	ssh -t oci1 if test -e /usr/local/bin/goe \; then cp -pri /usr/local/bin/goe goe.`date +%s` \; else sudo cp -pri goe.linux.arm64 /usr/local/bin/goe \; fi
	ssh -t eahdroplet4 if test -e /usr/local/bin/goe \; then cp -pri /usr/local/bin/goe goe.`date +%s` \; else sudo cp -pri goe.linux.amd64 /usr/local/bin/goe \; fi
	# if no previous version, copy new version into place so remaining commands will work
	ssh -t oci1 if test -e /etc/systemd/system/goe.service \; then cp -pri /etc/systemd/system/goe.service goe.service.`date +%s` \; else sudo cp -pri goe.service /etc/systemd/system \; fi
	ssh -t eahdroplet4 if test -e /etc/systemd/system/goe.service \; then cp -pri /etc/systemd/system/goe.service goe.service.`date +%s` \; else sudo cp -pri goe.service /etc/systemd/system \; fi
	ssh -t oci1 'sudo sh -c "mv /usr/local/bin/goe /usr/local/bin/goe.last && mv goe.linux.arm64 /usr/local/bin/goe && chcon --reference /usr/local/bin/goe.last /usr/local/bin/goe && systemctl restart goe"'
	ssh -t eahdroplet4 'sudo sh -c "mv /usr/local/bin/goe /usr/local/bin/goe.last && mv goe.linux.amd64 /usr/local/bin/goe && chcon --reference /usr/local/bin/goe.last /usr/local/bin/goe && systemctl restart goe"'

install: goe
	sudo install ./goe /usr/local/bin/goe

$(GEN_PRETAG_FILES) &: $(PROTO_FILES) /opt/homebrew/bin/protoc-gen-go
	protoc -I$(wildcard $(GOPATH)/pkg/mod/github.com/srikrsna/protoc-gen-gotag*) --go_out=./proto_pretag --go_opt=paths=source_relative $(PROTO_FILES) --proto_path $(PROTO_PATH)

$(GEN_FILES) &: $(GEN_PRETAG_FILES) $(GOPATH)/bin/protoc-gen-connect-go $(GOPATH)/bin/protoc-gen-gotag
	cp -pr $(GEN_PRETAG_FILES) ./proto/
	protoc -I$(wildcard $(GOPATH)/pkg/mod/github.com/srikrsna/protoc-gen-gotag*) --gotag_out=outdir=./proto:. --gotag_opt=paths=source_relative --connect-go_out=./proto --connect-go_opt=paths=source_relative $(PROTO_FILES) --proto_path $(PROTO_PATH)

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
