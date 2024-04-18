.PHONY: release test deploy run gen
goe: *.go */*.go stats.pb.go stats_grpc.pb.go
	go build

run: goe
	./goe

release: goe.linux

goe.linux: test.success goe
	GOOS=linux GOARCH=amd64 go build -o goe.linux

test: *.go db.created
	go test -v && scripts/verify_no_extra_output.sh && touch test.success

test.success: test

db.created: scripts/loadData.js scripts/setup_empty_db.sh
	scripts/setup_empty_db.sh

deploy: goe.linux goe
	echo latest version:
	./goe --version
	echo deployed version:
	ssh eahdroplet4 /usr/local/bin/goe --version
	scp goe.linux eahdroplet4:
	ssh eahdroplet4.ernie.org cp -pri /usr/local/bin/goe goe.`date +%s`
	ssh -t eahdroplet4.ernie.org 'sudo sh -c "mv /usr/local/bin/goe /usr/local/bin/goe.last && mv goe.linux /usr/local/bin/goe && chcon --reference /usr/local/bin/goe.last /usr/local/bin/goe && systemctl restart goe"'

install: goe
	sudo install ./goe /usr/local/bin/goe

gen: ../goeverywhere/grpcify/stats.proto
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     ../goeverywhere/grpcify/stats.proto  --proto_path ../goeverywhere/grpcify/
