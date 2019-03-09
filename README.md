# gofx [![Build Status](https://travis-ci.org/antonmedv/gofx.svg?branch=master)](https://travis-ci.org/antonmedv/gofx)
 
[fx](https://github.com/antonmedv/fx)-like command-line JSON processing tool. This is implementation of some functionality of fx tool.

## Features

* Don't need to learn new syntax
* Written in Go
* Formatting and highlighting

## Differences

* Only ES5 (no arrow functions, no spread)
* Small binary size
* Can't use npm packages

## Install

```
$ go get github.com/antonmedv/gofx
```

Or download precompiled binary from [releases](https://github.com/antonmedv/gofx/releases) page.

## Usage

Pipe into `gofx` any JSON and JS code for reducing it.

```
$ gofx [code ...]
```

Pretty print JSON without passing any arguments:
```
$ echo '{"key":"value"}' | gofx
{
    "key": "value"
}
```

### This Binding

You can get access to JSON by `this` keyword:
```
$ echo '{"foo": [{"bar": "value"}]}' | gofx 'this.foo[0].bar'
value
```

### Dot

It is possible to omit `this` keyword:

```
$ echo '{"foo": [{"bar": "value"}]}' | gofx .foo[0].bar
value
```

### Chain

You can pass any number of code blocks for reducing JSON:
```
$ echo '{"foo": [{"bar": "value"}]}' | gofx 'this.foo' 'this[0]' 'this.bar'
value
```

### Formatting

If you need something different then JSON (for example arguments for xargs) do not return anything from reducer.
`undefined` value printed into stderr by default.
```
$ echo '[]' | gofx 'void 0'
undefined
```

```
$ echo '[1,2,3]' | gofx 'this.forEach(function (x) { console.log(x) })' 2>/dev/null | xargs echo
1 2 3
```

### Modifying

To modify object use command separated by comma `,` and return `this` as the end.

```
$ echo '{"a": 2}' | gofx 'this["b"] = Math.pow(this.a, 10), this'
{
  "a": 2,
  "b": 1024
}
``` 

### Object keys

Get all object keys:
```
$ echo '{"foo": 1, "bar": 2}' | gofx 'Object.keys(this)'
[
  "foo",
  "bar"
]
```

By the way, gofx has shortcut for `Object.keys(this)`. Previous example can be rewritten as:

```
$ echo '{"foo": 1, "bar": 2}' | gofx ?
``` 

## Related

* [fx](https://github.com/antonmedv/fx) – original `fx` package
* [ymlx](https://github.com/matthewadams/ymlx) - `fx`-like YAML cli processor

## License

MIT
