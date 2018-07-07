# Test goreleaser

Simple recipe for a GitHub project using [goreleaser](https://github.com/goreleaser/goreleaser), and [dep](https://github.com/golang/dep) with the [no vendor directory commit](https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory) strategy.

## CI

This uses CircleCI to update each release tag with defined binaries. I also added a simple `go test` workflow to branches (including PR integration) to show that other CircleCI jobs/workflows can be combined with this recipe (note if you want to enable Pull Requests from forks see [this doc page](https://circleci.com/docs/2.0/oss/#build-pull-requests-from-forked-repositories)).

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
