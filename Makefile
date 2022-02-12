
NAME := hi

PKG_FILES := $(wildcard pkg/*/*.go)
CMD_FILES := $(wildcard cmd/*/*.go)
CMD_NAMES := $(patsubst cmd/%,%, $(wildcard cmd/*))

run-examples: run-hi

h run-hi:
	< cruft/dhcp.log ./hi.go \
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

info list: .deps
	@ echo ./.deps:; sed 's/^/	/g' .deps; echo
	@ echo "PKG_FILES: $(PKG_FILES)"
	@ echo "CMD_FILES: $(CMD_FILES)"
	@ echo "CMD_NAMES: $(CMD_NAMES)"

complain:
	find ./ -type f -name \*.go | xargs golint

clean:
	git clean -dfx

$(CMD_NAMES): $(PKG_FILES) Makefile

.deps: Makefile
	@ for i in $(CMD_NAMES); do echo $$i: cmd/$$i/main.go; \
		echo "	go build -v ./cmd/$$i"; \
		echo "	strip $$i"; \
		done > $@

include .deps

