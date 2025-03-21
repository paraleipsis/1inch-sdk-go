# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html) starting with the *v1.0.0-beta.1* release.

## [v1.0.0-beta.3] - 2025-1-22
[v1.0.0-beta.3 release page](https://github.com/paraleipsis/1inch-sdk-go/releases/tag/v1.0.0-beta.3)
### Changed
- Fusion Plus support added
- Pending Fusion orders can now be tracked using the SDK
- Orderbook client updated to support new API schema

## [v1.0.0-beta.2] - 2024-10-23
[v1.0.0-beta.2 release page](https://github.com/paraleipsis/1inch-sdk-go/releases/tag/v1.0.0-beta.2)
### Changed
- Classic Swap updated to use V6 API
- Added examples for all Classic Swap endpoints
- When using TransactionBuilder, if no `gas` value is specified in the transaction config, `eth_estimateGas` will be used by default

## [v1.0.0-beta.1] - 2024-8-22
[v1.0.0-beta.1 release page](https://github.com/paraleipsis/1inch-sdk-go/releases/tag/v1.0.0-beta.1)

Note: This changelog summarizes all changes since the last *changelog* version of v0.0.3-developer-preview

### Added
- Web3 API added
- Fusion SDK added
- Portfolio API added
- Permit1 support added for Orderbook orders and Aggregator Swaps

### Changed
- Readme updated to link to all API docs and examples
- Updating Geth version
- Types generation script updated to handle Web3 API spec design
- Normalized and improved SDK examples
- Improved code generation to make optional parameters pointers

# [v0.0.3-developer-preview] - 2024-3-9
[v0.0.3-developer-preview](https://github.com/paraleipsis/1inch-sdk/releases/tag/v0.0.3-developer-preview)

### New Features and Enhancements:

- All non-global query configurations have been moved to the request-level
  params ([PR](https://github.com/paraleipsis/1inch-sdk/pull/6))
    - RPC providers for all chains will now be defined/set at SDK startup
- Query parameters now use concrete types instead of pointers ([PR](https://github.com/paraleipsis/1inch-sdk/pull/16))
- Limit orders created within the SDK now support auto-expiration ([PR](https://github.com/paraleipsis/1inch-sdk/pull/23))
- Permit1 properly supported for limit orders when possible (fallback to Approval if Permit1 does not
  work) ([commit](https://github.com/paraleipsis/1inch-sdk/commit/f2e79e5f0e81503bfeeff076e41455e86e5a5120))
- When creating a limit order, integrators can error out when an approval is needed. This is useful for integrators who
  want all onchain actions to be performed manually by their users ([PR](https://github.com/paraleipsis/1inch-sdk/pull/26))

### Optimizations and Bug Fixes:

- Tenderly forks are cleaned up automatically at the beginning of each test
  run ([PR](https://github.com/paraleipsis/1inch-sdk/pull/6))
- Validation pattern for swagger-generated input params is now fully handled on all
  endpoints ([PR](https://github.com/paraleipsis/1inch-sdk/pull/8))
- Project-wide validation scripts added to verify validation logic
  standards ([PR](https://github.com/paraleipsis/1inch-sdk/pull/11))

# [v0.0.2-developer-preview] 2024-1-23
Tag: [v0.0.2-developer-preview](https://github.com/paraleipsis/1inch-sdk/releases/tag/v0.0.2-developer-preview)

### New Features and Enhancements:

- **Added Tenderly support for e2e swap tests**
    - e2e tests will now create forks, apply state overrides, and run simulations when a Tenderly API key is provided.
- **Added approval type selection**
    - Users can choose between `Approve` and `Permit1` (`Permit2` currently unsupported)
- **Implemented nonce cache to address RPC lag**
    - Once a wallet has posted a transaction, the nonce of that transaction is tracked and incremented internally by the
      SDK.

### Optimizations and Bug Fixes:

- Updated orderbook to use string inputs instead of integers to support all of uint256.
- Increased gas limit and reduced permit duration to improve transactions success and debugging.
- Moved Actions into a service within the main client to consolidate SDK structure.
- Simplified tests and refactored onchain actions to have more uniformity across the library.

# Release (January 15, 2024)

Tag: [v0.0.1-developer-preview](https://github.com/paraleipsis/1inch-sdk/releases/tag/v0.0.1-developer-preview)

### New Features and Enhancements:

### Limit Order support

- Enables posting orders to the 1inch Limit Order Protocol
- Enables reading orders from the 1inch Limit Order Protocol
- Most endpoints from the Limit Order API supported
    - `has-active-orders-with-permit` REST endpoint still untested

### Aggregator Protocol support

- All REST endpoints supported
- Get quotes and swap data from the Aggregator Protocol

### Onchain execution support

- Execute swaps onchain from within the SDK


