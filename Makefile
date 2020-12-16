GOOS=linux
GOARCH=amd64
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

.PHONY: build


build:
	$(GOGET)
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o s3-query -ldflags "-w -s"  -a .