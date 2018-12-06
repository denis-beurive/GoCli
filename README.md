# Description

This repository contains a CLI parser. This parser is inspired from the Perl's [Getopt::Long](http://perldoc.perl.org/Getopt/Long.html).

> Please note that, although the package is usable, it misses an example and a document.
> I will write documentation and examples when I have time for it.

# Prerequisites

First, make sure that the environment variable `GOPATH` is well configured. If not, then you can fix this quickly:

    export GOPATH=$(pwd):${GOPATH}
    echo ${GOPATH}

# Build the package

    cd src/beurive.com/cli
    go install

The package can be found under the directory `pkg`.

# Testing

    go test beurive.com/cli

Or:

    go test -test.v beurive.com/cli

# Compiling

    go build -v -x -work beurive.com/cli && echo $?

# Building the example program

    go install main

