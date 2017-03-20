# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]
### Added
* Basic implementation.

### Fixed
* Don't drop `setgroups` anymore, as it previously would allow for a user to
  drop their supplementary groups and subvert ACL modes like `0X0X`. The only
  really sane way of handling it is to not touch supplementary groups.
