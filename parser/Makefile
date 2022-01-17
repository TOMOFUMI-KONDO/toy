.SUFFIXES: .peg .go

.peg.go:
	peg -noast -output $@ $<

all: parser.go

clean:
	rm -f parser.go
