# Gist Backup

Gist Backup is a command-line tool for managing backups of your Gists.

Gist Backupe is published under the
[MIT license](http://opensource.org/licenses/MIT). It requires Go version 1.8
or later.


## Installation

    go get -u github.com/stefanbirkner/gist-backup
    go install github.com/stefanbirkner/gist-backup


## Usage

Provide your GitHub API token

    export GITHUB_TOKEN=<token>

You can get your gist `token` here [https://github.com/settings/tokens/new](https://github.com/settings/tokens/new), remember to select `gist` scope.

Backup all your Gists including the Gists you are starring

    $GOPATH/bin/gist-backup

The backup of each Gist is stored in a folder with the Gist's name.


## Contributing

You have three options if you have a feature request, found a bug or
simply have a question about Gist Backup.

* [Write an issue.](https://github.com/stefanbirkner/gist-backup/issues/new)
* Create a pull request. (See [Understanding the GitHub Flow](https://guides.github.com/introduction/flow/index.html))
* [Write a mail to mail@stefan-birkner.de](mailto:mail@stefan-birkner.de)


## Development Guide

If you want to contribute code then Fork the repo and create a pull request.
(See
[Understanding the GitHub Flow](https://guides.github.com/introduction/flow/index.html))
