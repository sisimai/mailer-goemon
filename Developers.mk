# libsisimai.org/mailer-goemon/Developers.mk
#  ____                 _                                       _    
# |  _ \  _____   _____| | ___  _ __   ___ _ __ ___   _ __ ___ | | __
# | | | |/ _ \ \ / / _ \ |/ _ \| '_ \ / _ \ '__/ __| | '_ ` _ \| |/ /
# | |_| |  __/\ V /  __/ | (_) | |_) |  __/ |  \__ \_| | | | | |   < 
# |____/ \___| \_/ \___|_|\___/| .__/ \___|_|  |___(_)_| |_| |_|_|\_\
#                              |_|                                   
# -------------------------------------------------------------------------------------------------
SHELL := /bin/sh
HERE  := $(shell pwd)
FILE  := $(firstword $(MAKEFILE_LIST))
NAME  := mailer-goemon
MKDIR := mkdir -p
LS    := ls -1
RM    := rm -f
CP    := cp
GO    := go

GOROOT := $(shell echo $$GOROOT)
GOPATH := $(shell echo $$GOPATH)

LIBSISIMAI := libsisimai.org
SISIMAIDIR := address moji rfc1123 rfc5322 rfc791 smtp/*/
COVERAGETO := coverage.txt
EXECUTABLE := bin/maigo
BUILDFLAGS := -ldflags="-s -w" -trimpath
GOLANGLINT := golangci-lint
GO_SYSNAME := $(shell echo $$GOOS   || $(GO) env GOOS  )
GO_CPUARCH := $(shell echo $$GOARCH || $(GO) env GOARCH)
LISTENADDR := 127.0.0.1:5322
K          := neko

# -------------------------------------------------------------------------------------------------
.PHONY: clean
$(EXECUTABLE):
	CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o $@ $@.go

$(EXECUTABLE).$(GO_SYSNAME)-$(GO_CPUARCH):
	GOOS=$(GO_SYSNAME) GOARCH=$(GO_CPUARCH) CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o $@ $(EXECUTABLE).go

build:
	$(RM) $(EXECUTABLE)
	$(MAKE) -f $(FILE) $(EXECUTABLE)

cross-build:
	# https://go.dev/doc/install/source#environment
	test -n "$(GO_SYSNAME)"
	test -n "$(GO_CPUARCH)"
	$(RM) $(EXECUTABLE).$(GO_SYSNAME)-$(GO_CPUARCH)
	$(MAKE) -f $(FILE) $(EXECUTABLE).$(GO_SYSNAME)-$(GO_CPUARCH)

test:
	@ $(GO) test ./...

count-test-cases:
	@ $(GO) test -v ./... | grep 'The number of ' | awk '{ cx += $$7 } END { print cx }'

loc:
	@ find $(SISIMAIDIR) -type f -name '*.go' -not -name '*_test.go' | \
		xargs grep -vE '(^$$|^//|/[*]|[*]/|^ |^--)' | grep -vE "\t+//" | wc -l

coverage:
	@ $(GO) test -v ./ $(addprefix ./, $(SISIMAIDIR)) -coverprofile=$(COVERAGETO)

profile:
	test -f bin/cpu-prof.go && CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o cpu-maigo ./bin/cpu-prof.go
	test -f ./cpu-maigo     && ./cpu-maigo ./$(PROFILESET)
	go tool pprof --top ./cpu.pprof > usage-of-cpu-x

	test -f bin/mem-prof.go && CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o mem-maigo ./bin/mem-prof.go
	test -f ./mem-maigo     && ./mem-maigo ./$(PROFILESET)
	go tool pprof --top ./mem.pprof > usage-of-mem-x

	ls -laF ./usage-of-*

benchmark:
	test -f bin/benchmark.go && CGO_ENABLED=0 $(GO) build $(BUILDFLAGS) -o min-maigo ./bin/benchmark.go
	uptime
	while true; do zsh -c 'time ./min-maigo $(PROFILESET)'; sleep 10; done

lint:
	test -x `which $(GOLANGLINT)`
	NO_COLOR=2 $(GOLANGLINT) run $(SISIMAIDIR)

find:
	find . -type f -name '*.go' -not -name '*_test.go' -not -path '*/bin/*' -not -path '*/sbin/*' \
		-not -path '*/stash/*' -not -path '*/tmp/*' -exec grep '$(K)' {} +

init:
	test -e ./go.mod || $(GO) mod init $(LIBSISIMAI)/$(NAME)

update-go-mod:
	@ $(GO) mod tidy

start-godoc-server:
	open http://$(LISTENADDR)
	godoc -http=$(LISTENADDR)

clean:
	$(RM)    ./_reason-table.txt
	$(RM)    ./$(EXECUTABLE)
	$(RM)    ./$(COVERAGETO)

