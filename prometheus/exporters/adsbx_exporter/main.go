/*
Exports ADS-B Exchange data to Prometheus. It currently exports:
	- The number of MLAT feeders connected to ADS-B Exchange.

Based on the example at https://stackoverflow.com/a/73240605/8135687.
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const adsbx_url = "https://map.adsbexchange.com/sync/totalcount.json"

var httpClient = http.Client{
	Timeout: 5 * time.Second,
}

// The response from ADS-B Exchange.
type adsbxResponse struct {
	UpdateTime string        `json:"UPDATED"`
	Region1A   []interface{} `json:"1A"`
	Region1B   []interface{} `json:"1B"`
	Region1C   []interface{} `json:"1C"`
	Region2A   []interface{} `json:"2A"`
	Region2B   []interface{} `json:"2B"`
	Region2C   []interface{} `json:"2C"`
	Region3A   []interface{} `json:"3A"`
	Region3B   []interface{} `json:"3B"`
	Region3C   []interface{} `json:"3C"`
	Region4A   []interface{} `json:"4A"`
	Region4B   []interface{} `json:"4B"`
	Region4C   []interface{} `json:"4C"`
	Region5A   []interface{} `json:"5A"`
	Region5B   []interface{} `json:"5B"`
	Region5C   []interface{} `json:"5C"`
}

// AdsbxCollector metrics.
type adsbxCollector struct {
	adsbxMLATFeeders *prometheus.Desc
}

// Create a new AdsbxCollector.
func newAdsbxCollector() *adsbxCollector {
	return &adsbxCollector{
		adsbxMLATFeeders: prometheus.NewDesc("mlat_feeders",
			"The number of MLAT feeders connected to ADS-B Exchange.",
			nil, prometheus.Labels{"aggregator": "ADSBx"},
		),
	}
}

// Declare Describe method. Runs when Prometheus scrapes the exporter.
func (collector *adsbxCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.adsbxMLATFeeders
}

// Declare Collect method. Runs when Prometheus scrapes the exporter.
func (collector *adsbxCollector) Collect(ch chan<- prometheus.Metric) {
	// Get the ADS-B Exchange data.
	response, err := httpClient.Get(adsbx_url)
	if err != nil {
		log.Fatalf("Error getting ADS-B Exchange data: %v", err)
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading ADS-B Exchange data: %v", err)
		os.Exit(1)
	}

	// Parse the JSON response.
	var responseObject adsbxResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatalf("Error unmarshalling ADS-B Exchange data: %v", err)
		os.Exit(1)
	}

	// Sum the number of MLAT feeders in all regions.
	adsbxMLATFeeders := float64(responseObject.Region1A[0].(float64) + responseObject.Region1B[0].(float64) + responseObject.Region1C[0].(float64) +
		responseObject.Region2A[0].(float64) + responseObject.Region2B[0].(float64) + responseObject.Region2C[0].(float64) +
		responseObject.Region3A[0].(float64) + responseObject.Region3B[0].(float64) + responseObject.Region3C[0].(float64) +
		responseObject.Region4A[0].(float64) + responseObject.Region4B[0].(float64) + responseObject.Region4C[0].(float64) +
		responseObject.Region5A[0].(float64) + responseObject.Region5B[0].(float64) + responseObject.Region5C[0].(float64))

	// Set the exporter metrics.
	ch <- prometheus.MustNewConstMetric(collector.adsbxMLATFeeders, prometheus.GaugeValue, adsbxMLATFeeders)
}

// Create exporter and start the web server.
func main() {
	log.Print("Starting ADS-B Exchange exporter.")

	// Create new prometheus registry and add the adsbx collector.
	promReg := prometheus.NewRegistry()
	collector := newAdsbxCollector()
	promReg.MustRegister(collector)

	// Create the prometheus http handler and start the web server.
	handler := promhttp.HandlerFor(
		promReg,
		promhttp.HandlerOpts{
			EnableOpenMetrics: false, // Disable default golang metrics.
		})
	http.Handle("/metrics", handler)
	log.Fatal(http.ListenAndServe(":19100", nil))
}
