CMD=timer

GOOS=darwin
GOARCH=amd64

all: $(CMD)

$(CMD): $(CMD).go
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ $<
	chmod a+x $(CMD)
