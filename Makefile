NAME = particuland

all:
	go build -o ${NAME} .

run:
	go run .