# go-laboratory

This repository serves as a collection of code studies and experiments using the Go programming language. The contents
of this repository are not intended for production use, but rather for personal learning and exploration.

## Contents

The repository is organized into folders for each project or experiment, with a brief description provided below.

- **[bbolt](bbolt)**: a simple key-value store implementation in Go, where [bbolt](https://github.com/etcd-io/bbolt) is
  used as the embedded database engine.
- **[diskv](diskv)**: an implementation of a key-value store in Go that is based on file storage, utilizing
  the [diskv](https://github.com/peterbourgon/diskv) library.

## Getting Started

To run any of the experiments or projects in this repository, you will need to have Go installed on your system. Please
refer to the [official Go documentation](https://golang.org/doc/install) for installation instructions.

To run a specific project or experiment, navigate to its directory and run the `go run` command. For example, to run
the "bbolt" program, navigate to the [bbolt/](bbolt) directory and run:

```
$ go run ./...
```

## Contributing

Contributions to this repository are welcome! If you have an experiment or project you would like to share, please feel
free to submit a pull request.

## License

This repository is licensed under the [MIT License](LICENSE).
