# flagenv

[![godoc](https://godoc.org/github.com/kahlys/flagenv?status.svg)](https://pkg.go.dev/github.com/kahlys/flagenv)
[![go report](https://goreportcard.com/badge/github.com/kahlys/flagenv)](https://goreportcard.com/report/github.com/kahlys/flagenv)

Flagenv is a golang command-line and/or or environment variable flags parser. It is an extension of the [flag package](https://golang.org/pkg/flag/), and both works fine together.

## Usage

Define some flags, and parse them.

```go
var (
	name  = flagenv.String("name", "NAME", "coulson", "target name")
	level = flagenv.Int("level", "LEVEL", 7, "level value")
)

func main() {
	flagenv.Parse()
	fmt.Println(*name, *level)
}
```

Run it using command-line flags or environment variables. If a command-line flag and an environment variable are used for the same flag, the value of the command-line flag will be used.

```none
$ go run main.go -name "coulson" -level 7
$ NAME=coulson LEVEL=7 go run main.go
```