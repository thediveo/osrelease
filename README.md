<!-- markdownlint-disable-next-line MD022 -->
# A Pedantic `os-release` Parser

[![GitHub](https://img.shields.io/github/license/thediveo/osrelease)](https://img.shields.io/github/license/thediveo/osrelease)
![build and test](https://github.com/thediveo/osrelease/workflows/build%20and%20test/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/osrelease)](https://goreportcard.com/report/github.com/thediveo/osrelease)

`osrelease` is a Go module implementing a slightly pedantic parser for the
os-release file format, as specified in
[os-release(5)](https://www.freedesktop.org/software/systemd/man/os-release.html).
This implementation actually _parses_ quoted assignment values for escape
sequences in _correct sequence_, instead of simply doing blind
search-and-replaces (which can result in unexpected results).

Oh, and this module has _tests_.

There's a reason why we call it the "_pedantic_" parser after all.

## ⚖️ Copyright and License

`osrelease` is Copyright 2021 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
