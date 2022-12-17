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
go run aoc2022.go
```

To run with command line options:
```shell
go build aoc2022.go
./aoc2022 --help
./aoc2022 --days=1-10,17 --include-all
```
