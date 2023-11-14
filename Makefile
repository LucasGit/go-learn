

format:
	go build -o format  format.go

funct:
	go build -o funct  funct.go

all: format funct


test: all
	./format
	./funct

clean:
	rm -rf format
	rm -rf funct
