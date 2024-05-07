
all: build run

build:
	go build -o out

run:
	./out

clean:
	@echo "Cleaning up..."
	rm ./out

