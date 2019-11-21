package collector

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

type EthGetBalance struct {
	rpc  *rpc.Client
	desc *prometheus.Desc
}

func NewEthGetBalance(rpc *rpc.Client) *EthGetBalance {
	return &EthGetBalance{
		rpc: rpc,
		desc: prometheus.NewDesc(
			"eth_balance",
			"the current balance of the coinbase address",
			nil,
			nil,
		),
	}
}

func (collector *EthGetBalance) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desc
}

func (collector *EthGetBalance) Collect(ch chan<- prometheus.Metric) {
	var address common.Address
	if err := collector.rpc.Call(&address, "eth_coinbase"); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	var balance string
	if err := collector.rpc.Call(&balance, "eth_getBalance", address, "latest"); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	balanceBig, err := hexutil.DecodeBig(balance)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	balanceBigFloat := big.NewFloat(0).SetInt(balanceBig)

	normalBalanace := balanceBigFloat.Quo(balanceBigFloat, big.NewFloat(1000000000000000000))

	value, _ := normalBalanace.Float64()
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, value)
}
