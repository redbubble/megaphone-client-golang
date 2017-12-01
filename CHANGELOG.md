# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and
this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## 0.2.1 - 2017-12-01

### Fixed

- missing _tag prefix_ in the underlying Fluentd logger configuration (prevented messages from being published in their destination stream)

## 0.2.0 - 2017-11-29

### Changed

- Make `megaphone.Client` an interface so it could be mocked (e.g. using [pegomocks](https://github.com/petergtz/pegomock))
- Change the `megaphone.Client` constructor to require the `origin`, `host` and `port` directly instead of requiring a `megaphone.Config` structure

## 0.1.0 - 2017-11-28

### Added

- Initial implementation of the `megaphone.Client`
