.PHONY: install build

install:
	go get github.com/robertkrimen/otto
	go get github.com/hokaccha/go-prettyjson

build:
	gox -output "out/{{.Dir}}_{{.OS}}_{{.Arch}}"
