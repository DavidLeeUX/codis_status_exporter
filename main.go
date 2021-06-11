/*
 * @Descripttion:
 * @version: xv1.0
 * @Author: changwei5
 * @Date: 2019-04-18 18:17:02
 * @LastEditors: changwei5
 * @LastEditTime: 2021-06-11 14:01:11
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"all_kinds_monitoring/codis_status_exporter/exporter"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	codisAddr := flag.String("codis.addr", "http://127.0.0.1:28001", "admin codis dashboard ")
	namespace := flag.String("namespace", "codis", "Namespace for metrics")
	productName := flag.String("productname", "", "product name")
	listenAddress := flag.String("web.listen-address", ":8082", "Address to listen on for web interface and telemetry.")
	metricPath := flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.Parse()

	exporter, err := exporter.NewExporter(*namespace, *codisAddr, *productName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(`
	Codis prometheus exporter
	Access http => `, *listenAddress)

	// Define parameters

	prometheus.MustRegister(exporter)

	// Launch http service

	http.Handle(*metricPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Codis Exporter</title></head>
			<body>
			<h1>Codis Exporter</h1>
			<p><a href='` + *metricPath + `'>Metrics</a></p>
			</body>
			</html>`))
	})
	fmt.Println(http.ListenAndServe(*listenAddress, nil))
}
