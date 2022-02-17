
NAME := hi

default: release

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

hi: $(wildcard cmd/hi/*.go) $(wildcard pkg/*/*.go) Makefile version-izer.sh
	./version-izer.sh -v ./cmd/hi

release stripped: hi
	strip $<
