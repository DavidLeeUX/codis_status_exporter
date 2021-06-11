/*
 * @Descripttion:
 * @version: xv1.0
 * @Author: changwei5
 * @Date: 2019-04-18 18:51:48
 * @LastEditors: changwei5
 * @LastEditTime: 2021-06-11 13:59:30
 */
package exporter

import (
	"fmt"
	"log"
	"sync"

	"all_kinds_monitoring/codis_status_exporter/codisstat"

	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	MetricsPrefix   string
	Addr            string
	ProductName     string
	CodisData       *codisstat.CodisData
	UptimeInSeconds *prometheus.GaugeVec
	TotalQps        *prometheus.GaugeVec
	GroupMetric     map[string]*prometheus.GaugeVec
	ProxyMetric     map[string]*prometheus.GaugeVec
	GroupMetricsMtx sync.RWMutex
	ProxyMetricsMtx sync.RWMutex
}

func NewExporter(metricsPrefix string, addr string, productName string) (*Exporter, error) {

	UptimeInSeconds := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "uptime_in_seconds_Total",
		Help:      "uptime_in_secondsmetric_Total",
	}, []string{"addr", "product"})

	TotalQps := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "total_qps",
		Help:      "tital QPS",
	}, []string{"addr", "product"})
	e := &Exporter{
		Addr:            addr,
		ProductName:     productName,
		MetricsPrefix:   metricsPrefix,
		UptimeInSeconds: UptimeInSeconds,
		TotalQps:        TotalQps,
	}
	e.InitGroupMetric()
	e.InitProxyMetric()

	return e, nil
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.RunCollect(ch)
}

func (e *Exporter) RunCollect(ch chan<- prometheus.Metric) {

	var url string
	var err error
	if e.ProductName == "" {
		url = fmt.Sprintf("%s/topom", e.Addr)

	} else {
		url = fmt.Sprintf("%s/topom?forward=%s", e.Addr, e.ProductName)
	}
	e.CodisData, err = codisstat.NewCodisData(url)

	if err != nil {
		log.Printf("Codis connect Error : %v", err)
		return
	}
	e.UptimeInSeconds.WithLabelValues(e.Addr, e.ProductName).Set(float64(e.CodisData.GetStartTimeSec()))
	e.TotalQps.WithLabelValues(e.Addr, e.ProductName).Set(float64(e.CodisData.GetTotalQps()))
	e.GetGroupMetric(ch)
	e.GetProxyMetric(ch)
	e.UptimeInSeconds.Collect(ch)
	e.TotalQps.Collect(ch)

}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.TotalQps.Describe(ch)
	e.UptimeInSeconds.Describe(ch)
	for _, vm := range e.GroupMetric {
		vm.Describe(ch)
	}
	for _, vm := range e.ProxyMetric {
		vm.Describe(ch)
	}
}
