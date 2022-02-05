
NAME := hi

test: hi
	echo this is a test | ./hi test lime

$(NAME):
	go build -o $@
