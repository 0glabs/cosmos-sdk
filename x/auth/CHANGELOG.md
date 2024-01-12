<!--
Guiding Principles:
Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.
Usage:
Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:
* (<tag>) [#<issue-number>] Changelog message.
Types of changes (Stanzas):
"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"API Breaking" for breaking exported APIs used by developers building on SDK.
Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [Unreleased]

### Features

* [#18641](https://github.com/cosmos/cosmos-sdk/pull/18641) Support the ability to broadcast unordered transactions per ADR-070. See UPGRADING.md for more details on integration.
* [#18281](https://github.com/cosmos/cosmos-sdk/pull/18281) Support broadcasting multiple transactions.
* (vesting) [#17810](https://github.com/cosmos/cosmos-sdk/pull/17810) Add the ability to specify a start time for continuous vesting accounts.
* (tx) [#18772](https://github.com/cosmos/cosmos-sdk/pull/18772) Remove misleading gas wanted from tx simulation failure log.

### Improvements

* [#18780](https://github.com/cosmos/cosmos-sdk/pull/18780) Move sig verification out of the for loop, into the authenticate method.

### CLI Breaking Changes

* (vesting) [#18100](https://github.com/cosmos/cosmos-sdk/pull/18100) `appd tx vesting create-vesting-account` takes an amount of coin as last argument instead of second. Coins are space separated.

### API Breaking Changes

* [#17985](https://github.com/cosmos/cosmos-sdk/pull/17985) Remove `StdTxConfig`

### Consensus Breaking Changes

* [#18817](https://github.com/cosmos/cosmos-sdk/pull/18817) SigVerification, GasConsumption, IncreaseSequence ante decorators have all been joined into one SigVerification decorator. Gas consumption during TX validation flow has reduced.

### Bug Fixes