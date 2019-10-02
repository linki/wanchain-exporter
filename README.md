# Wanchain Prometheus Exporter

This service exports various metrics from Wanchain clients for consumption by [Prometheus](https://prometheus.io). It uses [JSON-RPC](https://iwan.wanchain.org/static/apidoc/docs.html) interface to collect the metrics. Any JSON-RPC 2.0 enabled client should be supported.

## Usage

You can deploy this exporter using the [linki/wanchain-exporter](https://hub.docker.com/r/linki/wanchain-exporter/) Docker image.

    docker run -d -p 9368:9368 --name wanchain-exporter linki/wanchain-exporter -url http://wanchain:8545

Keep in mind that your container needs to be able to communicate with the Wanchain client using the specified `url` (default is `http://localhost:8545`).

By default the exporter serves on `:9368` at `/metrics`. The listen address can be changed by specifying the `-addr` flag.

Here is an example [`scrape_config`](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config) for Prometheus.

```yaml
- job_name: wanchain
  static_configs:
  - targets:
    - wanchain-exporter:9368
```

## Exported Metrics

| Name | Description |
| ---- | ----------- |
| net_peers | The number of peers currently connected to the client. |
| eth_block_number | The number of most recent block. |
| eth_gas_price | The current price per gas in wei. *Might be inaccurate*. |
| eth_earliest_block_transactions | The number of transactions in an earliest block. |
| eth_latest_block_transactions | The number of transactions in a latest block. |
| eth_pending_block_transactions | The number of transactions in a pending block. |
| eth_hashrate | The number of hashes per second that the node is mining with. |
| eth_sync_starting | The block at which the import started. |
| eth_sync_current | The number of most recent block. |
| eth_sync_highest | The estimated highest block. |
| parity_net_active_peers | The number of active peers. *Available only for Parity*. |
| parity_net_connected_peers | The number of peers currently connected to the client. *Available only for Parity*. |

## Development

[Go modules](https://golang.org/doc/go1.11#modules) is used for dependency management. Hence Go 1.11 is a minimum required version.

[CircleCI Local CLI](https://circleci.com/docs/2.0/local-cli/) can be used to ensure that everything builds locally.

    circleci build --job lint
    circleci build --job test
    circleci build --job build

## Contributing

Contributions are greatly appreciated. The project follows the typical GitHub pull request model. Before starting any work, please either comment on an existing issue or file a new one.

## Support and Donate

You can show your appreciation for the project and support future development by donating.

[![Donate with Ethereum](https://en.cryptobadges.io/badge/big/0xcaDe516c2c2d916eDf44b958ED5B52C01039fad6)](https://en.cryptobadges.io/donate/0xcaDe516c2c2d916eDf44b958ED5B52C01039fad6)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
