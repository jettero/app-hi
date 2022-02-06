
NAME := hi
SRC_FILES := $(shell find ./ -type f -name \*.go)

# dummy or command targets

run-example: hi
	echo "This is an übertest with — multibyte unicode — characters." \
		| ./hi \
		. coal \
		'\S+test' purple \
		über red \
		multibyte yellow

complain:
	find ./ -type f -name \*.go | xargs golint

clean:
	git clean -dfx

# derived rules

$(NAME): $(SRC_FILES)
	go build -o $@
