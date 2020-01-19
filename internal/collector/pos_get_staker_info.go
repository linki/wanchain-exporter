package collector

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/linki/wanchain-cli/client"
	"github.com/linki/wanchain-cli/types"
	"github.com/linki/wanchain-cli/util"
	"github.com/prometheus/client_golang/prometheus"
)

type PosGetStakerInfo struct {
	rpc  *rpc.Client
	url  string
	desc *prometheus.Desc
}

func NewPosGetStakerInfo(rpc *rpc.Client, url string) *PosGetStakerInfo {
	return &PosGetStakerInfo{
		rpc: rpc,
		url: url,
		desc: prometheus.NewDesc(
			"pos_get_staker_info",
			"amount staked",
			nil,
			nil,
		),
	}
}

func (collector *PosGetStakerInfo) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desc
}

func (collector *PosGetStakerInfo) Collect(ch chan<- prometheus.Metric) {
	client, err := client.NewClient(collector.url)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	defer client.Close()

	blockHeight, err := client.GetCurrentBlockHeight()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	validators, err := client.GetValidators(blockHeight)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	var address common.Address
	if err := collector.rpc.Call(&address, "eth_coinbase"); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	var validator types.Validator
	for i, v := range validators {
		if v.Address == address {
			validator = validators[i]
			break
		}
	}

	validatorAmount := util.TotalAmountValidators([]types.Validator{validator})
	result := util.WeiToEth(validatorAmount)

	value, _ := result.Float64()
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, value)
}
