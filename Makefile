.PHONY: goe release test deploy
goe: *.go
	go build

release: goe.linux

goe.linux: test.success
	GOOS=linux GOARCH=amd64 go build -o goe.linux

test: *.go db.created
	go test -v && touch test.success

test.success: test

db.created:
	scripts/setup_empty_db.sh


deploy: goe.linux
	scp goe.linux eahdroplet4:
	ssh eahdroplet4.ernie.org cp -pri /usr/local/bin/goe goe.`date +%s`
	ssh -t eahdroplet4.ernie.org 'sudo sh -c "mv /usr/local/bin/goe /usr/local/bin/goe.last && mv goe.linux /usr/local/bin/goe && chcon --reference /usr/local/bin/goe.last /usr/local/bin/goe && systemctl restart goe"'
