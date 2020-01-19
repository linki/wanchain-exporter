package collector

import (
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

type EthLastBlock struct {
	rpc  *rpc.Client
	desc *prometheus.Desc
}

func NewEthLastBlock(rpc *rpc.Client) *EthLastBlock {
	return &EthLastBlock{
		rpc: rpc,
		desc: prometheus.NewDesc(
			"eth_last_block",
			"todo",
			nil,
			nil,
		),
	}
}

func (collector *EthLastBlock) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desc
}

type getBlockResult struct {
	Timestamp hexutil.Uint64
}

func (collector *EthLastBlock) Collect(ch chan<- prometheus.Metric) {
	var result hexutil.Uint64
	if err := collector.rpc.Call(&result, "eth_blockNumber"); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	var raw json.RawMessage
	if err := collector.rpc.Call(&raw, "eth_getBlockByNumber", "latest", true); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	var latestBlock *getBlockResult
	if err := json.Unmarshal(raw, &latestBlock); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	ts, _ := hexutil.DecodeUint64(latestBlock.Timestamp.String())

	blockTime := time.Unix(int64(ts), 0)
	currentTime := time.Now()

	diff := currentTime.Sub(blockTime)

	value := float64(diff)
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, value)
}
