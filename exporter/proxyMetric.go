package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

// 初始化proxy监控信息
func (e *Exporter) InitProxyMetric() {
	var (
		ProxyMetricLabelNames = []string{
			"id",
			"addr",
			"product",
			"ProxyAddr",
			"ProtoType",
		}
	)

	e.ProxyMetric = map[string]*prometheus.GaugeVec{}
	e.ProxyMetric["proxy_start_time"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_start_time",
		Help:      "proxy_start_time",
	}, ProxyMetricLabelNames)
	e.ProxyMetric["proxy_online_status"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_online_status",
		Help:      "proxy_online_status  status :online or  offline",
	}, append(ProxyMetricLabelNames, "status"))

	e.ProxyMetric["proxy_closed_status"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_closed_status",
		Help:      "proxy_closed_status  status: closed or open",
	}, append(ProxyMetricLabelNames, "status"))

	e.ProxyMetric["proxy_total_ops"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_total_ops",
		Help:      "proxy_total_ops",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_fails_ops"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_fails_ops",
		Help:      "proxy_fails_ops",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_curr_ops"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_curr_ops",
		Help:      "proxy_curr_ops",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_redis_ops_error"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_redis_ops_error",
		Help:      "proxy_redis_ops_error",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_sessions_total"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_sessions_total",
		Help:      "proxy_sessions_total",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_sessions_alive"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_sessions_alive",
		Help:      "proxy_sessions_alive",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_rusage_cpu"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_rusage_cpu",
		Help:      "proxy_rusage_cpu",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_rusage_mem"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_rusage_mem",
		Help:      "proxy_rusage_mem",
	}, ProxyMetricLabelNames)
	e.ProxyMetric["proxy_utime"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_utime",
		Help:      "proxy_utime",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_stime"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_stime",
		Help:      "proxy_stime",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_cutime"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_cutime",
		Help:      "proxy_cutime",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_cstime"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_cstime",
		Help:      "proxy_cstime",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_num_threads"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_num_threads",
		Help:      "proxy_num_threads",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_vm_size"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_vm_size",
		Help:      "proxy_vm_size",
	}, ProxyMetricLabelNames)

	e.ProxyMetric["proxy_vm_rss"] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: e.MetricsPrefix,
		Name:      "proxy_vm_rss",
		Help:      "proxy_vm_rss",
	}, ProxyMetricLabelNames)
}

func (e *Exporter) GetProxyMetric(ch chan<- prometheus.Metric) {
	ProxyMetricLabel := make([]string, 5)

	proxyData := e.CodisData.ProcessProxyData()

	for _, v := range proxyData {
		ProxyMetricLabel[0] = string(v.Id)
		ProxyMetricLabel[1] = e.Addr
		ProxyMetricLabel[2] = v.Product_name
		ProxyMetricLabel[3] = v.Proxy_addr
		ProxyMetricLabel[4] = v.Proto_type
		e.ProxyProcessValues(ProxyMetricLabel, v.Start_time, "proxy_start_time")
		if v.Online {
			e.ProxyProcessValues(append(ProxyMetricLabel, "online"), 1, "proxy_online_status")
		} else {
			e.ProxyProcessValues(append(ProxyMetricLabel, "offline"), 1, "proxy_online_status")
		}
		if v.Closed {
			e.ProxyProcessValues(append(ProxyMetricLabel, "closed"), 1, "proxy_closed_status")
		} else {
			e.ProxyProcessValues(append(ProxyMetricLabel, "open"), 1, "proxy_closed_status")
		}
		e.ProxyProcessValues(ProxyMetricLabel, v.Ops_total, "proxy_total_ops")
		e.ProxyProcessValues(ProxyMetricLabel, v.Ops_fails, "proxy_fails_ops")

		e.ProxyProcessValues(ProxyMetricLabel, v.Curr_qps, "proxy_curr_ops")
		e.ProxyProcessValues(ProxyMetricLabel, v.Ops_redis_error, "proxy_redis_ops_error")
		e.ProxyProcessValues(ProxyMetricLabel, v.Sessions_alive, "proxy_sessions_alive")
		e.ProxyProcessValues(ProxyMetricLabel, v.Sessions_total, "proxy_sessions_total")
		e.ProxyProcessValues(ProxyMetricLabel, v.Rusage_mem, "proxy_rusage_mem")
		e.ProxyProcessValues(ProxyMetricLabel, v.Utime, "proxy_utime")
		e.ProxyProcessValues(ProxyMetricLabel, v.Stime, "proxy_stime")
		e.ProxyProcessValues(ProxyMetricLabel, v.Cutime, "proxy_cutime")
		e.ProxyProcessValues(ProxyMetricLabel, v.Cstime, "proxy_cstime")
		e.ProxyProcessValues(ProxyMetricLabel, v.Num_threads, "proxy_num_threads")
		e.ProxyProcessValues(ProxyMetricLabel, v.Vm_size, "proxy_vm_size")
		e.ProxyProcessValues(ProxyMetricLabel, v.Vm_rss, "proxy_vm_rss")
		e.ProxyMetricsMtx.Lock()
		e.ProxyMetric["proxy_rusage_cpu"].WithLabelValues(ProxyMetricLabel...).Set(v.Rusage_cpu)
		e.ProxyMetricsMtx.Unlock()
	}
	for _, vm := range e.ProxyMetric {
		vm.Collect(ch)
	}
}

func (e *Exporter) ProxyProcessValues(ProxyMetricLabel []string, value int64, mapKey string) {
	e.ProxyMetricsMtx.Lock()
	e.ProxyMetric[mapKey].WithLabelValues(ProxyMetricLabel...).Set(float64(value))
	e.ProxyMetricsMtx.Unlock()
}
