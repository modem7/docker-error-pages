# Changelog

All notable changes to this package will be documented in this file.

The format is based on [Keep a Changelog][keepachangelog] and this project adheres to [Semantic Versioning][semver].

## V1.2.1

### Added

- Added parallax template from [Thom-X](https://github.com/Thom-x/docker-error-pages)

## V1.2.0

### Added

- Added dependabot.yml

### Fixed

- Switched from nginx:mainline-alpine to nginx:1-alpine to fix [libgcrypt vulnerability](https://snyk.io/vuln/SNYK-ALPINE314-LIBGCRYPT-1582772)

## V1.1.1

### Added

- Added dependabot.yml
- Added hexxone template from [hexxone](https://github.com/hexxone/error-pages)

## V1.0.0

### Added

- First project release. See upstream [CHANGELOG.MD](https://github.com/tarampampam/error-pages/blob/master/CHANGELOG.md)

### Changed

- Set X-Forwarded-For for logs and enable multiarch builds
- Added additional error statuses
- Added healthcheck endpoint

[keepachangelog]:https://keepachangelog.com/en/1.0.0/
[semver]:https://semver.org/spec/v2.0.0.html