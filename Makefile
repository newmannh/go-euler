ifeq ($(shell uname -s), Linux)
# Do a build without dependencies for docker image.  Much slower.
GO_BUILD = CGO_ENABLED=0 go build -a -installsuffix cgo
else
GO_BUILD = go build -a
endif

.PHONY: all
all : show test build
.PHONY: build
build :
	$(GO_BUILD) -o bin/server .

# Use the 'go vet' tool to examine Go source code and report suspicious constructs for each package in project
.PHONY: vet
vet :
	find . ! -path "./vendor/*" -a -name "*.go" -print0 | xargs -0 -n1 go vet

.PHONY: gofmt
gofmt:
	find . ! -path "./vendor/*" -a -name "*.go" -print0 | xargs -0 -n1 gofmt -w

.PHONY: show
show :
	@echo "========================================"
	@echo "GOPATH=${GOPATH}"
	@echo "GOBIN=${GOBIN}"
	@echo "SRCDIRS=${SRCDIRS}"
	@echo "TESTDIRS=${TESTDIRS}"
	@go version
	@echo "========================================"

# List all available argets/tasks
.PHONY: no_targets__ list
no_targets__:
list:
	sh -c "$(MAKE) -p no_targets__ | awk -F':' '/^[a-zA-Z0-9][^\$$#\/\\t=]*:([^=]|$$)/ {split(\$$1,A,/ /);for(i in A)print A[i]}' | grep -v '__\$$' | sort"

.PHONY: clean
clean :
	rm -f *.out
	rm -rf bin
	find . -iname *.test | xargs -I LALA rm LALA
