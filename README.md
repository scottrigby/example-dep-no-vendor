# Test goreleaser

Simple recipe for a GitHub project using [goreleaser](https://github.com/goreleaser/goreleaser), and [dep](https://github.com/golang/dep) with the [no vendor directory commit](https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory) strategy.

## Installation

Example installation steps for a project following this recipe:

### Binaries

Download your preferred asset from the [releases page](https://github.com/scottrigby/test-goreleaser/releases) and install manually.

### Go get

```sh
go get -d github.com/scottrigby/test-goreleaser
cd $GOPATH/src/github.com/scottrigby/test-goreleaser
dep ensure -vendor-only
make setup build
```

### Homebrew

To-do: set up a tap for this example, and define in `.goreleaser.yml`.
