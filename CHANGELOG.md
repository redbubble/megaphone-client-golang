# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and
this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## 1.1.0 - 2018-11-15

### Added

- Add `PublishRawMessage` method to all the client to publish raw messages to a Kinesis stream. 


## 1.0.0 - 2018-03-09

### Changed

- Rename `Client` to `FluentdClient`
- Rename `Config` to `FluentdConfig`

### Added

- Implement `megaphone.Publisher` using a AWS Kinesis Client, which writes to Kinesis synchronously. 

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
