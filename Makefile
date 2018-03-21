deps:
	go get -v github.com/whyrusleeping/gx-go
	gx install
test:
	go fmt
	go test