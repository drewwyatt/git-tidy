workflow "Release" {
  on = "push"
  resolves = ["goreleaser"]
}

workflow "Build" {
  on = "pull_request"
  resolves = ["build"]
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
