
NAME := hi

PKG_FILES := $(wildcard pkg/*/*.go)
CMD_FILES := $(wildcard cmd/*/*.go)
CMD_NAMES := $(patsubst cmd/%,%, $(wildcard cmd/*))

run-examples: run-hi

h run-hi: hi
	echo "This is an übertest with — multibyte unicode — characters." \
		| ./hi \
		'[ico]' coal \
		'\S+test' purple \
		über red \
		multibyte yellow \
	    tiby nc_file \
		be umber \
		ib mc_curs \
		es violet

info list:
	@ sed -e 's/^/.deps: /' .deps
	@ echo "PKG_FILES: $(PKG_FILES)"
	@ echo "CMD_FILES: $(CMD_FILES)"
	@ echo "CMD_NAMES: $(CMD_NAMES)"

complain:
	find ./ -type f -name \*.go | xargs golint

clean:
	git clean -dfx

$(CMD_NAMES): $(PKG_FILES) Makefile

.deps: Makefile
	@ for i in $(CMD_NAMES); do echo $$i: cmd/$$i/main.go; echo "	go build ./cmd/$$i"; \
		done > $@

include .deps

