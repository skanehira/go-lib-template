![GitHub Repo stars](https://img.shields.io/github/stars/skanehira/go-lib-template?style=social)
![GitHub license](https://img.shields.io/github/license/skanehira/go-lib-template)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/skanehira/go-lib-template)
![GitHub CI Status](https://img.shields.io/github/workflow/status/skanehira/go-lib-template/ci?label=CI)
![GitHub Release Status](https://img.shields.io/github/workflow/status/skanehira/go-lib-template/Release?label=release)

# go-lib-template
This is template that help you to quick implement some CLI using Go.

This repository is contains following.

- minimal CLI implementation using [spf13/cobra](https://github.com/spf13/cobra)
- CI/CD
  - golangci-lint
  - go test
  - goreleaser
  - dependabot for github-actions and Go
  - CodeQL Analysis (Go)

## How to use
1. fork this repository
2. replace `skanehira` to your user name using `sed`(or others)
3. run `make init`

## Author
skanehira
