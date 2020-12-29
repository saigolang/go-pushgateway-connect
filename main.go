package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {

	var labelNames []string
	metriclabels := map[string]string{
		"Name": "testing",
	}

	for label, _ := range metriclabels {
		labelNames = append(labelNames, label)
	}

	record := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "testing",
	}, labelNames)

	record.With(metriclabels).Add(1)

	registry := prometheus.NewRegistry()
	pusher := push.New("http://localhost:9091/", "testJob").Gatherer(registry)
	err := pusher.Push()

	if err != nil {
		fmt.Println("We got an error and the error is ", err.Error())
	}

}
