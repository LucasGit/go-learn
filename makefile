
format:
	go build -o format  format.go

main:
	go build -o main  main.go

all: main

test: all
	cd fn;./unittest.sh;cd ..
	
	./main

clean:
	rm -rf main
	rm -rf format
