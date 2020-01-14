package vegeta

import (
	"io"
	"log"
	"strconv"

	"github.com/DataDog/datadog-go/statsd"
)

type DatadogReporter struct {
	statsd *statsd.Client
}

func (dr *DatadogReporter) Add(m *Result) {
	tags := []string{"attack:" + m.Attack, "error:" + m.Error, "url:" + m.URL, "code:" + strconv.Itoa(int(m.Code))}

	dr.statsd.TimeInMilliseconds("datadog.vegeta.latency", float64(m.Latency.Seconds()*1000), tags, 1.0)
	dr.statsd.Histogram("datadog.vegeta.bytes.out", float64(m.BytesOut), tags, 1.0)
	dr.statsd.Histogram("datadog.vegeta.bytes.in", float64(m.BytesIn), tags, 1.0)
}

func NewDatadogReporter() *DatadogReporter {
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	return &DatadogReporter{
		statsd: statsd,
	}
}

func (dr *DatadogReporter) FlushMetrics() Reporter {
	return func(w io.Writer) error {
		dr.statsd.Flush()
		return nil
	}
}
