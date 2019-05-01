workflow "Release" {
  on = "push"
  resolves = ["goreleaser"]
}

workflow "PR" {
  on = "pull_request"
  resolves = ["build", "test", "verify"]
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
  args = "go test ./..."
}

action "verify" {
  uses = "docker://golang:1.11"
  args = "go mod verify"
}
