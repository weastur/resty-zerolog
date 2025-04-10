# resty-zerolog

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/weastur/resty-zerolog)](https://goreportcard.com/report/github.com/weastur/resty-zerolog)
[![codecov](https://codecov.io/gh/weastur/resty-zerolog/graph/badge.svg?token=CYFYMFHNON)](https://codecov.io/gh/weastur/resty-zerolog)
[![test](https://github.com/weastur/resty-zerolog/actions/workflows/test.yaml/badge.svg)](https://github.com/weastur/resty-zerolog/actions/workflows/test.yaml)
[![lint](https://github.com/weastur/resty-zerolog/actions/workflows/lint.yaml/badge.svg)](https://github.com/weastur/resty-zerolog/actions/workflows/lint.yaml)</br>
![GitHub Release](https://img.shields.io/github/v/release/weastur/resty-zerolog)
![GitHub commits since latest release](https://img.shields.io/github/commits-since/weastur/resty-zerolog/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/weastur/resty-zerolog)
![GitHub License](https://img.shields.io/github/license/weastur/resty-zerolog)

</div>

**resty-zerolog** is a simple wrapper for zerolog to use it as an implementation of
[Logger](https://pkg.go.dev/github.com/go-resty/resty/v3#Logger) interface from [Resty](https://resty.dev/)

## Why?

I wanted to use [zerolog](https://github.com/rs/zerolog) as a logger for [Resty](https://resty.dev/),
but Resty does not support zerolog out of the box. So I created this simple wrapper.

Yes, [Logger](https://pkg.go.dev/github.com/go-resty/resty/v3#Logger) interface from Resty is very simple,
and you can implement it in a few lines of code,
but I wanted to have a ready-to-use solution to **not repeat myself** in every project I use resty and zerolog.

## Installation

```bash
go get -u github.com/weastur/resty-zerolog
```

## Usage

```go
import (
	"github.com/rs/zerolog/log"
	restyzerolog "github.com/weastur/resty-zerolog"
	"resty.dev/v3"
)

client := resty.New()
client.SetLogger(restyzerolog.New(log.Logger))
```

Despite it's extremely simple, you can refer to the [example](./_example/) and
[godoc](https://pkg.go.dev/github.com/weastur/resty-zerolog) to see a bit more.

## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on how to contribute to this project.

## Security

Refer to the [SECURITY.md](SECURITY.md) file for more information.

## License

Mozilla Public License 2.0

Refer to the [LICENSE](LICENSE) file for more information.
