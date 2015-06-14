# submit - to submit a code to AOJ

[![Build Status](https://travis-ci.org/upamune/submit.svg?branch=master)](https://travis-ci.org/upamune/submit)

## Description
ターミナルからAOJに提出できるコマンドです。提出時の言語は拡張子から自動的に決められます。
AOJのIDとパスワードを環境変数 ```AOJID``` と ```AOJPASS``` にあらかじめ登録しておく必要があります。

## Usage

```submit [PROBLEM NUMBER] [FILE]```


### Same

- ```submit 1 hello.cpp```
- ```submit 01 hello.cpp```
- ```submit 001 hello.cpp```
- ```submit 0001 hello.cpp```
- ```submit 0001 hello.cpp```

## Install

To install, use `go get`:

```bash
$ go get github.com/upamune/submit
```

## Contribution

1. Fork ([https://github.com/upamune/submit/fork](https://github.com/upamune/submit/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[upamune](https://github.com/upamune)
