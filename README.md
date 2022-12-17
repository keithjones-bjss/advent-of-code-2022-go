# Advent of Code 2022 (Go)

An attempt at the Advent of Code 2022 challenge in Go. I have not written any Go before, so this probably won't be
pretty.

## Testing
```shell
go test ./...
```

## Running

To run all days, excluding the expensive parts:
```shell
go run main.go
```

To run with command line options:
```shell
go build main.go
./main --help
./main --days=1-10,17 --include-all
```
