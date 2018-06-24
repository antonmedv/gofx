<img src="https://user-images.githubusercontent.com/141232/41759457-dbfb230a-7618-11e8-9159-03aa62e57152.png" height="100" alt="fx">

# [![Build Status](https://travis-ci.org/antonmedv/xx.svg?branch=master)](https://travis-ci.org/antonmedv/xx)
 
[fx](https://github.com/antonmedv/fx)-like command-line JSON processing tool 

## Features

* Don't need to learn new syntax
* Written in Go
* Formatting and highlighting

## Differences

* Only ES5 (no arrow functions, no spread)
* Small binary size (xx ~5mb vs fx ~30mb)
* Can't use npm packages

## Install

```
$ go get github.com/antonmedv/xx
```

Or download precompiled binary from [releases](https://github.com/antonmedv/xx/releases) page.

## Usage

Pipe into `xx` any JSON and JS code for reducing it.

```
$ xx [code ...]
```

Pretty print JSON without passing any arguments:
```
$ echo '{"key":"value"}' | xx
{
    "key": "value"
}
```

### This Binding

You can get access to JSON by `this` keyword:
```
$ echo '{"foo": [{"bar": "value"}]}' | xx 'this.foo[0].bar'
"value"
```

### Chain

You can pass any number of code blocks for reducing JSON:
```
$ echo '{"foo": [{"bar": "value"}]}' | xx 'this.foo' 'this[0]' 'this.bar'
"value"
```

### Formatting

If you need something different then JSON (for example arguments for xargs) do not return anything from reducer.
`undefined` value printed into stderr by default.
```
$ echo '[]' | xx 'void 0'
undefined
```

```
$ echo '[1,2,3]' | xx 'this.forEach(function (x) { console.log(x) })' 2>/dev/null | xargs echo
1 2 3
```

### Modifying

To modify object use command separated by comma `,` and return `this` as the end.

```
$ echo '{"a": 2}' | xx 'this["b"] = Math.pow(this.a, 10), this'
{
  "a": 2,
  "b": 1024
}
``` 

### Object keys

Get all object keys:
```
$ echo '{"foo": 1, "bar": 2}' | xx 'Object.keys(this)'
[
  "foo",
  "bar"
]
```

By the way, xx has shortcut for `Object.keys(this)`. Previous example can be rewritten as:

```
$ echo '{"foo": 1, "bar": 2}' | xx ?
``` 

## Related

* [fx](https://github.com/antonmedv/fx) – original `fx` package
* [ymlx](https://github.com/matthewadams/ymlx) - `fx`-like YAML cli processor

## License

MIT
