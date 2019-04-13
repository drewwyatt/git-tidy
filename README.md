# 🧽 Gitclean

[![release](https://img.shields.io/github/release-pre/drewwyatt/gitclean.svg?style=for-the-badge)](https://github.com/drewwyatt/gitclean/releases)

Delete local git branches (in the current directory) that have a remote tracking branch that is `: gone`.

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

**With force delete:**

```bash
$ gitclean -f # same as above, but with "-D" instead of "-d"
```
