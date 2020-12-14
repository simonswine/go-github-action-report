# go-github-action-report

Golang test formatter for grouping multi-line output.

It parses the verbose `go test` output (code base on [go-junit-report]) and
groups it according to [Github commands]

## Usage

```shell
go test -v ./... | go-github-action-report
```

[go-junit-report]: https://github.com/jstemmer/go-junit-report
[GitHub commands]: https://github.com/actions/toolkit/blob/1cc56db0ff126f4d65aeb83798852e02a2c180c3/docs/commands.md#group-and-ungroup-log-lines
