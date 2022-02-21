
NAME := hi

# Go sucks so bad you need a shell program to get the major version
GO_VERSION := $(shell go version | cut -d' ' -f3 | sed -e 's/^go//' | cut -d. -f1,2)

default: release

check:
	goreleaser check

snapshot: check
	goreleaser release --snapshot --rm-dist

install: release
	install hi $$GOPATH/bin/

r run-examples: run-hi

h run-hi:
	< cruft/dhcp.log ./hi.sh \
		. coal \
		'DHCP(\w+)' cyan \
		DHCP ocean \
		'10.0.0.\d+' lime \
		'10.0.0.' green \
		'10.255.255.\d+' umber \
		'10.255.255.' red \
		'[a-f0-9:]{17}' unique \
		'(?<=\()[^)]+(?=\))' yellow \
		eth0 lime \
		eth1 umber

valgrind: hi
	echo supz | valgrind \
		-v --track-origins=yes --leak-check=full --log-file=valgrind.log \
		./hi supz lime
	less -eS +G valgrind.log

complain:
	find ./ -type f -name \*.go | xargs golint

clean:
	git clean -dfx
	go mod tidy -v -go=$(GO_VERSION) -compat=$(GO_VERSION)

hi: $(wildcard cmd/hi/*.go) $(wildcard pkg/*/*.go) Makefile version-izer.sh
	./version-izer.sh -v ./cmd/hi

release stripped: hi
	strip $<
