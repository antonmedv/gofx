#!/usr/bin/env bash
set -euxo pipefail

echo '{"key":"value"}' | xx
echo '{"foo": [{"bar": "value"}]}' | xx 'this.foo[0].bar'
echo '{"foo": [{"bar": "value"}]}' | xx 'this.foo' 'this[0]' 'this.bar'
echo '[]' | xx 'void 0'
echo '[1,2,3]' | xx 'this.forEach(function (x) {console.log(x)})' 2>/dev/null | xargs echo
echo '{"foo": 1, "bar": 2}' | xx 'Object.keys(this)'
echo '{"foo": 1, "bar": 2}' | xx ?
