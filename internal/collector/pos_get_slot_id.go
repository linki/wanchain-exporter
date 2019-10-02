package collector

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

type PosGetSlotID struct {
	rpc  *rpc.Client
	desc *prometheus.Desc
}

func NewPosGetSlotID(rpc *rpc.Client) *PosGetSlotID {
	return &PosGetSlotID{
		rpc: rpc,
		desc: prometheus.NewDesc(
			"pos_get_slot_id",
			"current slot in epoch",
			nil,
			nil,
		),
	}
}

func (collector *PosGetSlotID) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desc
}

func (collector *PosGetSlotID) Collect(ch chan<- prometheus.Metric) {
	var result int
	if err := collector.rpc.Call(&result, "pos_getSlotID"); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}

	value := float64(result)
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, value)
}
