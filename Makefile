GO ?= go
GOFMT ?= gofmt "-s"
GO_VERSION=$(shell $(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
PACKAGES ?= $(shell $(GO) list ./...)
VETPACKAGES ?= $(shell $(GO) list ./... | grep -v /examples/)
GOFILES := $(shell find . -name "*.go")
TESTFOLDER := $(shell $(GO) list ./... | grep -v examples)
TESTTAGS ?= ""

.PHONY: test
test:
	echo "mode: count" > coverage.out
	for d in $(TESTFOLDER); do \
		$(GO) test -tags $(TESTTAGS) -v -covermode=count -coverprofile=profile.out $$d > tmp.out; \
		cat tmp.out; \
		if grep -q "^--- FAIL" tmp.out; then \
			echo 1; \
			rm tmp.out; \
			exit 1; \
		elif grep -q "build failed" tmp.out; then \
			echo 2\
			rm tmp.out; \
			exit 1; \
		elif grep -q "setup failed" tmp.out; then \
			echo 3\
			rm tmp.out; \
			exit 1; \
		fi; \
		if [ -f profile.out ]; then \
			echo 5\
			cat profile.out | grep -v "mode:" >> coverage.out; \
			rm profile.out; \
		fi; \
	done

