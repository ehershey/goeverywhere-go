.PHONY: release test deploy run gen

BUILD_TIME = $(shell date +"%Y-%m-%d %H:%M:%S")
COMMIT_HASH = $(shell git log -1 --pretty=format:%h)
GO_VERSION = $(shell go version | cut -f3- -d\ )

FLAGS = -X \"main.BuildTime=$(BUILD_TIME)\"
FLAGS += -X \"main.CommitHash=$(COMMIT_HASH)\"
FLAGS += -X \"main.GoVersion=$(GO_VERSION)\"

goe: *.go */*.go proto/stats.pb.go proto/stats_grpc.pb.go
	go build -ldflags "$(FLAGS)"

run: goe
	./goe

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

gen: ../goeverywhere/grpcify/stats.proto /opt/homebrew/bin/protoc-gen-go-grpc /opt/homebrew/bin/protoc-gen-go
	protoc --go_out=./proto --go_opt=paths=source_relative     --go-grpc_out=./proto --go-grpc_opt=paths=source_relative     ../goeverywhere/grpcify/stats.proto  --proto_path ../goeverywhere/grpcify/

/opt/homebrew/bin/protoc-gen-go:
	brew install protoc-gen-go

/opt/homebrew/bin/protoc-gen-go-grpc:
	brew install protoc-gen-go-grpc
