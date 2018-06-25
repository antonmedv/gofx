package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hokaccha/go-prettyjson"
	"github.com/mattn/go-colorable"
	"github.com/robertkrimen/otto"
)

var vm = otto.New()

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fatal(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		usage()
		os.Exit(1)
	}

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fatal(err)
	}
	var input interface{}
	if err := json.Unmarshal(buf, &input); err != nil {
		fatal(err)
	}

	if err := vm.Set("json", input); err != nil {
		fatal(err)
	}
	value, err := vm.Get("json")
	if err != nil {
		fatal(err)
	}
	for _, code := range os.Args[1:] {
		value, err = reduce(value, code)
		if err != nil {
			fatal(err)
		}
	}

	if value.IsUndefined() {
		fmt.Fprintln(os.Stderr, "undefined")
		return
	}

	i, err := value.Export()
	if err != nil {
		fatal(err)
	}
	s, err := prettyjson.Marshal(i)
	if err != nil {
		fatal(err)
	}
	fmt.Fprintln(colorable.NewColorableStdout(), string(s))
}

func reduce(value otto.Value, code string) (otto.Value, error) {
	if err := vm.Set("json", value); err != nil {
		fatal(err)
	}
	switch code {
	case "?":
		code = "Object.keys(json)"
	default:
		code = fmt.Sprintf(`(function () {return %v;}).call(json)`, code)
	}
	result, err := vm.Run(code)
	if err != nil {
		return otto.UndefinedValue(), err
	}
	return result, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `
  Usage
    $ xx [code ...]

  Examples
    $ echo '{"key": "value"}' | xx 'this.key'
    "value"

    $ echo '[1,2,3]' | xx 'this.map(function (x) { return x * 2; })'
    [2, 4, 6]

    $ echo '{"items": ["one", "two"]}' | xx 'this.items' 'this[1]'
    "two"

`)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
