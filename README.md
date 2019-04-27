# 🧽 Gitclean

[![release](https://img.shields.io/github/release-pre/drewwyatt/gitclean.svg?style=for-the-badge)](https://github.com/drewwyatt/gitclean/releases)

Delete local git branches that have a remote tracking branch that is `: gone`.

## Installation

### Homebrew

```bash
$ brew tap drewwyatt/tap
$ brew install gitclean
```

## Usage

```bash
$ gitclean # executes "git branch -d" on ": gone" branches
```

### With force delete

```bash
$ gitclean -f # same as above, but with "-D" instead of "-d"
# or
$ gitclean --force
```

### Interactive

Present all ": gone" branches in a checkbox list, allowing user to opt-in to deletions.

```bash
$ gitclean -i
# or
$ gitclean --interactive
# with force
$ gitclean -if
# or
$ gitclean --interactive --force
```

### Directory

By default, `gitclean` will execute all commands in the **current directory** (`.`), however, you can pass a path to another git repository after any/all other flags.

```bash
$ gitclean ../some/other/repo
# with flags
$ gitclean -if ../some/other/repo
```
