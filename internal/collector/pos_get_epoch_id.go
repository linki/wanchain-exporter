package collector

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

type PosGetEpochID struct {
	rpc  *rpc.Client
	desc *prometheus.Desc
}

func NewPosGetEpochID(rpc *rpc.Client) *PosGetEpochID {
	return &PosGetEpochID{
		rpc: rpc,
		desc: prometheus.NewDesc(
			"pos_get_epoch_id",
			"current epoch",
			nil,
			nil,
		),
	}
}

func (collector *PosGetEpochID) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desc
}

func (collector *PosGetEpochID) Collect(ch chan<- prometheus.Metric) {
	var result int
	if err := collector.rpc.Call(&result, "pos_getEpochID"); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	value := float64(result)
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, value)
}
