workflow "Release" {
  on = "push"
  resolves = ["goreleaser"]
}

workflow "PR" {
  on = "pull_request"
  resolves = ["build", "test", "coverage", "verify"]
}

action "is-tag" {
  uses = "actions/bin/filter@master"
  args = "tag"
}

action "goreleaser" {
  uses = "docker://goreleaser/goreleaser"
  secrets = [
    "GITHUB_TOKEN",
    "DOCKER_USERNAME",
    "DOCKER_PASSWORD",
  ]
  args = "release"
  needs = ["is-tag"]
}

action "build" {
  uses = "docker://golang:1.11"
  args = "go build"
}

action "test" {
  uses = "docker://golang:1.11"
  args = "go test -race -coverprofile=coverage.txt -covermode=atomic ./..."
}

action "coverage" {
  uses = "docker://debian:9.5-slim"
  # needs = ["test"]
  args = "./.github/upload-coverage.sh"
  secrets = ["CODECOV_TOKEN"]
}

action "verify" {
  uses = "docker://golang:1.11"
  args = "go mod verify"
}
