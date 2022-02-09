
NAME := hi

PKG_FILES := $(wildcard pkg/*/*.go)
CMD_FILES := $(wildcard cmd/*/*.go)
CMD_NAMES := $(patsubst %.go,%, $(notdir $(CMD_FILES)))

run-example: hi
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

q: quick
	@echo
	@./quick sky on coal
	@./quick nc_file
	@./quick white on blue
	@./quick alert

list:
	@ echo "PKG_FILES: $(PKG_FILES)"
	@ echo "CMD_FILES: $(CMD_FILES)"
	@ echo "CMD_NAMES: $(CMD_NAMES)"

complain:
	find ./ -type f -name \*.go | xargs golint

clean:
	git clean -dfx

$(CMD_NAMES): $(PKG_FILES) Makefile

.deps: Makefile
	@ echo "building rules for targets: $(CMD_NAMES)"
	@ for i in $(CMD_NAMES); do echo $$i: cmd/$$i/$$i.go; echo "	go build -o $$i \$$<"; done > $@

include .deps

