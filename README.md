prompt
======

For use in PS1. Returns a concatenated version of $(pwd) based on a max length parameter. 

features
========

* elides intermediate directories in favor of longer ones.
* `/Users/cvandyck/code/go/src/github.com/collinvandyck/prompt` -> `~/.../github.com/collinvandyck/prompt`

usage
=====

    go get github.com/collinvandyck/prompt

In your .bash_profile:

    export PS1='`prompt -length 30` % '
