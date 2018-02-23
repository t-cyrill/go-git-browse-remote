# go-git-browse-remote

git-browser-remote open github on repository in golang

It clone ruby's git-browse-remote. Original [motemen/git-browse-remote](https://github.com/motemen/git-browse-remote)

## Installation

```
go get -u github.com/t-cyrill/go-git-browser-remote/...
```

## Build

Just use go build.

```
go build -o git-browser-remote
```

## Usage

```
NAME:
   git-browse-remote - open git remote repository on your web browser

USAGE:
   git-browse-remote [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --stdout             prints URL instead of opening browser
   --pullrequest, --pr  open pull request URL instead of top
   --directory value    change working directory
   --help, -h           show help
   --version, -v        print the version
```

