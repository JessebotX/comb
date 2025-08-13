GOEXE = go
GOFMTEXE = gofmt
GOFMTARGS = -s -w -l
GOLINTEXE = staticcheck
GOLINTARGS =

all: build

build:
	$(GOEXE) build ./example/comb

fmt:
	$(GOFMTEXE) $(GOFMTARGS) *.go
	$(GOFMTEXE) $(GOFMTARGS) example/comb/*.go

vet: lint
check: lint
lint:
	$(GOLINTEXE) $(GOLINTARGS) ./...

clean:
	rm -f comb
