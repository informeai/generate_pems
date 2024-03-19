# Generate pems

Create keys the cryptograph ecdsa with sha256 utility

### Required

go version >= 1.18

## Build

### Windows

```
GOOS=windows GOARCH=amd64 go build -o generate_pems main.go
```

### Linux

```
GOOS=linux GOARCH=amd64 go build -o generate_pems main.go
```

### Macos

```
GOOS=darwin GOARCH=amd64 go build -o generate_pems main.go
```

## Run

without build

```
go run main.go
```

with build

```
./generate_pems
```
