# ghtasklist

**ghtasklist** is a small utility to automatically prefix lines in order to
format them as [Github task lists][1].

[1]: https://github.com/blog/1375%0A-task-lists-in-gfm-issues-pulls-comments

## Install

    go get github.com/bfontaine/ghtasklist

## Usage

    $ cat my-input
    this
    is
    the
    result
    $ ghtasklist < my-input
    - [ ] this
    - [ ] is
    - [ ] the
    - [ ] result
