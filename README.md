# Distribute Go with dep + no vendor dir

This repo is a simple example of how to easily distribute a Go project using [dep](https://github.com/golang/dep) without [committing the vendor directory](https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory).

Tl;dr, use [goreleaser](https://github.com/goreleaser/goreleaser) with your favorite CI.

Details below…

## Why not commit the vendor dir?

I prefer not to commit the vendor directory becuase of the [cons listed in the dep FAQ](https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory) (PR diffs are confusing when code changes are mixed in with potentially thousands of dependency file changes, and repo size can get HUGE). Also, in principle I prefer not to duplicate vendor dependencies in source code (IMO the packaging step is a much better place to do that).

## Awesome! But the distribution problem…

Simple distribution of the project isn't obvious anymore. For example, `go get` requires additional manual steps:
```console
$ go get -d X/Y/Z
$ cd $GOPATH/src/X/Y/Z
$ dep ensure -vendor-only
$ go install
```

This is fine for developers, but it should be simpler for end users.

## Simplify distribution with binaries

A great tool for this is [goreleaser](https://github.com/goreleaser/goreleaser):
> GoReleaser builds Go binaries for several platforms, creates a GitHub release and then pushes a Homebrew formula to a tap repository. All that wrapped in your favorite CI.

I made a `.goreleaser.yml` file with `goreleaser init`, and just modified the file to include Windows and ARM.

Simple usage: running `goreleaser` in the project root will update the lastest git release tag with binary distribution artefacts. Pretty neat.

## Automate it

You can automate goreleaser with your favorite CI. This example repo uses [CircleCI](https://circleci.com) to automatically update each release tag with defined binaries. goreleaser has an example [circle 1.0 config file](https://goreleaser.com/#circle), but I wanted to make a more updated circle 2.0 config file (check the `.circleci/config.yml` file git commit history for what and why).

The circle config file also contains a simple `go test` workflow to branches (including PR integration) to show that other jobs/workflows can be easily combined (note if you want to enable Pull Requests from forks like I did for this example repo, see [this doc page](https://circleci.com/docs/2.0/oss/#build-pull-requests-from-forked-repositories)).

## Example installation instructions

### Binaries

Download your preferred asset from the [releases page](https://github.com/scottrigby/test-goreleaser/releases) and install manually.

### Go get

```console
$ go get -d github.com/scottrigby/test-goreleaser
$ cd $GOPATH/src/github.com/scottrigby/test-goreleaser
$ dep ensure -vendor-only
$ go install
```

### Homebrew

To-do: set up a tap for this example, and define in `.goreleaser.yml`.
