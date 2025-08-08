<img width="1920" height="1080" alt="din x402" src="https://github.com/user-attachments/assets/c922e0ab-3fc2-4a1c-af64-1b23e072a9d4" />

# DIN+x402 examples

[DIN](https://din.build) is making several of our supported networks available to users without a subscription.  Instead, the networks below can be accessed by making an x402 micro-payment for each call.  This gateway is provided as a beta release and feedback from BUIDLRS is very appreciated.  Submit feedback, feature requests and bug reports by filing [issues in this repository](https://github.com/DIN-center/din-x402-examples/issues).

## Supported networks

Crypto-currency networks are specified as specific sub-domains on the DIN gateway and match the `*.gateway.din.dev/rpc` hostname pattern.  The following networks are currently supported:

* Ethereum Mainnet: https://ethereum-mainnet.gateway.din.dev/rpc (`beta`/near-head only)
* Ethereum Hoodi: https://ethereum-hoodi.gateway.din.dev/rpc (`stable`)
* Base Mainnet: https://base-mainnet.gateway.din.dev/rpc (`beta`)
* Base Sepolia: https://base-sepolia.gateway.din.dev/rpc (`beta`)
* Blast Mainnet: https://blast-mainnet.gateway.din.dev/rpc (`stable`)
* Blast Sepolia: https://blast-sepolia.gateway.din.dev/rpc (`stable`)


Also note that the words `stable` and `beta` shown after each network above refer to the quality of the network connection available through DIN.

## Examples

Examples are available for the following languages:

* [Go](/main.go)
* [Python](/main.py)

### Go

To run the Go example, first ensure the dependencies are up to date:

```bash
go mod tidy
```

Then, set the `X402_BUYER_PRIVATE_KEY` environment variable and run the script:

```bash
export X402_BUYER_PRIVATE_KEY=<your_private_key>
go run main.go
```

### Python

To run the Python example, first create a virtual environment and install the dependencies:

```bash
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

Then, set the `X402_BUYER_PRIVATE_KEY` environment variable and run the script:

```bash
export X402_BUYER_PRIVATE_KEY=<your_private_key>
python main.py
```

## Clients

Clients are provided by external libraries and are available for the following languages:

* [Go](https://github.com/selesy/x402-buyer)
* [Java](https://github.com/coinbase/x402/tree/main/java/src/main/java/com/coinbase/x402/client)
* [Python](https://github.com/coinbase/x402/tree/main/python/x402/src/x402/clients)
* [Rust](https://github.com/x402-rs/x402-rs/tree/b13a4714432ae3421b0421c61103e5a57f8b5d38/crates/x402-reqwest)
* [TypeScript](https://github.com/coinbase/x402/tree/main/typescript/packages)

## Roadmap

The following feature list is a very rough roadmap for the DIN gateway.  We'll be building more features and expanding this product after the code nyc hackathon. If you'd like to provide feedback or work with us, send us an email to din@consensys.net or post your thoughts to the [discussions in this repository](https://github.com/DIN-center/din-x402-examples/discussions).

* [x] Ethereum-compatible JSON-RPC service
* [ ] Ethereum-compatible MCP service
* [ ] API to list available networks and methods
* [ ] The rest of the DIN networks
* [x] Micro-payments (one per request)
* [ ] Daily access tokens for different tiers

## References

* [DIN](https://din.build/)
* [x402](https://www.x402.org/)
* [CoinBase CodeNYC](https://www.coinbase.com/developer-platform/codenyc)
