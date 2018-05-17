GOCMD=go
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD)fmt

setup:
	go get -u github.com/kardianos/govendor

vendoring:
	govendor remove +external
	govendor add +external
	govendor update +external

test:
	@echo "Running iss-notifier tests"
	$(GOTEST) -v ./...

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w *.go