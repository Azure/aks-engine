package metrics

import (
	"fmt"

	"github.com/Azure/aks-engine/pkg/helpers"
	"github.com/alexcesaro/statsd"
)

type mdmBucket struct {
	Namespace string            `json:"Namespace"`
	Metric    string            `json:"Metric"`
	Dims      map[string]string `json:"Dims"`
}

// AddMetric adds the defined metric to a list of metrics to send to MDM
func AddMetric(endpoint, namespace, metric string, count int64, dims map[string]string) error {
	bucket := mdmBucket{
		Namespace: namespace,
		Metric:    metric,
		Dims:      dims}
	data, err := helpers.JSONMarshal(bucket, false)
	if err != nil {
		return err
	}
	client, err := statsd.New(
		statsd.Address(endpoint),
		statsd.Network("udp"),
		statsd.ErrorHandler(
			func(err error) {
				fmt.Println(err.Error())
			}))
	if err != nil {
		return err
	}
	defer client.Close()
	client.Count(string(data), count)
	return nil
}
