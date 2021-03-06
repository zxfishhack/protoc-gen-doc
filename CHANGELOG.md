# Changelog

All noteworthy changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) (as of Feb 2018)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/pseudomuto/protoc-gen-doc/compare/v1.1.0...master)

### Added

* Added [Sprig](https://github.com/Masterminds/sprig) functions to renderer contexts
* Added `IsMap` to `MessageField` objects indicating whether or not the field is a map field

### Fixed

* CI issue related to Regexp comparison on Golang master

## [v1.1.0](https://github.com/pseudomuto/protoc-gen-doc/compare/v1.0.0...v1.1.0) - March 13, 2018

### Added

* Switch to using [protokit] to handle parsing and running the plugin
* Default values are not included in the built-in templates
* Added the ability to ignore certain files by adding `:pattern[,pattern]` to `--doc_opt`
* Added `-help` and `-version` flags to the binary (e.g. `./protoc-gen-doc -help`)

### Changed

* Dev tooling, now using dep and retool
* CI setup to use go 1.10 (and master) and protoc 3.5.1
* Docker image updated to use protoc 3.5.1 as well

## Fixed

* Several issues with the documentation
* `NrBr` filter was a little too aggressive

### Removed

* The entire `parser` package (in favor of [protokit])

[protokit]: https://github.com/pseudomuto/protokit

# v1.0.0 - September 26, 2017

This is the tenth official release. (However, it's the first 1.x release!!)

**New things**

* Support for both proto2 and proto3!
* Application is now dockerized, no need to install it (unless you want to of course)
* Ported to Go and precompiled for Linux, OSX, and Windows
* Added test coverage to ensure functionality and prevent regressions
* Ignore comments starting with `@exclude`
* Added (and backfilled) CHANGELOG.md
* Added CONTRIBUTING.md

**Bug fixes**

* Message enums no longer included in file enums (#288)
* Added top-level `files` key for JSON output and _camelCased_ all fields (#289)
* Fixed issue with TOC in Markdown not nesting correctly (#293)
* No need to worry about expiring apt keys (#295)
* Extra slashes in comment prefix no longer show up in output (#298)
* Added field details for templates (#300)
* Markdown headers include new line so they render correctly (#303)

**Breaking changes**

* File-level comments are now attached to the syntax directive
* JSON fields are now _camelCased_
* Mustache templates replaced with golang templates
* Dropped direct support for PDF generation (still possible with FOP)
* `doc_out` flag removed in favor of `doc_opt`


# v0.9 - February 26, 2017

This is the ninth official release.

**changes**

* Improve installation instructions for macOS (thanks @guozheng)
* Improve installation instructions for Debian/Ubuntu (thanks @mhaberler)
* Add asciidoc.mustache example template (thanks @ArcEye)
* Don't do HTML escaping in Markdown template (thanks @sunfmin)
* Add support for JSON output

# v0.8 - February 26, 2016

This is the eight official release.

**changes**

* Add support for documenting files (#9)
* Add support for default values (#11)
* Add no-exclude flag to ignore @exclude directives (#13)
* Add support for RPC services (#14) (thanks to @murph0 !)

# v0.7 - January 7, 2016

This is the seventh official release.

**changes**

* Added support for extensions (thanks @masterzen !)
* Added Custom Templates wiki page
* Added additional distro packages for Debian 8, Ubuntu 15.04 + 15.10 and Fedora 22 + 23

# v0.6 - April 8, 2015

This is the sixth official release.

No functional changes were made, but Linux distribution package repositories are now provided for Ubuntu, Arch, Fedora
and openSUSE through the Open Build Service, and an RPM for CentOS 7 here below.

# v0.5 - December 19, 2014

This is the fifth official release.

**changes**

* Support exclusion also of enum values (accidental omission in 0.4).

# v0.4 - December 19, 2014

This is the fourth official release.

**changes**

* Updated to a newer version of qt-mustache.
* Updated Windows zip to libprotobuf/libprotoc 2.6.1.
* Added support for excluding messages/enums/fields.

# v0.3 - August 19, 2014

This is the third official release.

**changes**

* Updated to a newer version of qt-mustache which is more spec compliant.
* Added missing documentation for enums to Markdown output.

# v0.2 - August 14, 2014

This is the second official release.

* No functional changes were made, but the build system was improved.

# v0.1 - August 6, 2014

* Initial release
