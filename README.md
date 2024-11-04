# Machina

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/machina.svg)](https://pkg.go.dev/github.com/aethiopicuschan/machina)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/machina)](https://goreportcard.com/report/github.com/aethiopicuschan/machina)
[![CI](https://github.com/aethiopicuschan/machina/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/machina/actions/workflows/ci.yaml)

Machina is a tool that automatically generates configuration files for [Kataribe](https://github.com/matsuu/kataribe).
This tool allows you to generate the necessary bundle based on the provided source code.

This tool was developed for [ISUCON](https://isucon.net/).

## Installation

```sh
go install github.com/aethiopicuschan/machina@latest
```

## Usage

```sh
machina app.py > kataribe.toml
```

## Supported languages

- [x] Python(Flask)
- [ ] Golang
- [ ] Node.js
- [ ] Perl
- [ ] PHP
- [ ] Ruby
- [ ] Rust
