# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project do not adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [21.3.1] - 2021-08-30
### Added
- Support to role `tenkai-manager` added in keycloak
- Filter by email on endpoint /users
- Audit on endpoints:
  - /createOrUpdateUserEnvironmentRole
  - /createOrUpdateUser

### Changed
- The behavior of endpoint /users/{id} was changed according to permission level of requester
  - to `tenkai-admin` requester was keeped the same behavior
  - to `tenkai-manager` requester the user's environments returned now are restricted to those who requester have access

### Fixed
- Delete role associated with user and environment when user loses his access to this environment

