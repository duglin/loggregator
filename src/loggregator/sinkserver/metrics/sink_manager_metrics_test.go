package metrics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"loggregator/sinks/dump"
	"loggregator/sinks/syslog"
	"loggregator/sinks/websocket"

	"loggregator/sinks"
	"loggregator/sinkserver/metrics"
)

var _ = Describe("SinkManagerMetrics", func() {

	var sinkManagerMetrics *metrics.SinkManagerMetrics
	var sink sinks.Sink

	BeforeEach(func() {
		sinkManagerMetrics = metrics.NewSinkManagerMetrics()
	})

	It("Should have metrics for dump sinks", func() {

		Expect(sinkManagerMetrics.Emit().Metrics[0].Name).To(Equal("numberOfDumpSinks"))
		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))

		sink = &dump.DumpSink{}
		sinkManagerMetrics.Inc(sink)

		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(1))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))

		sinkManagerMetrics.Dec(sink)

		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))
	})

	It("Should have metrics for syslog sinks", func() {

		Expect(sinkManagerMetrics.Emit().Metrics[1].Name).To(Equal("numberOfSyslogSinks"))
		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))

		sink := &syslog.SyslogSink{}
		sinkManagerMetrics.Inc(sink)

		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(1))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))

		sinkManagerMetrics.Dec(sink)

		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))
	})

	It("Should have metrics for websocket sinks", func() {

		Expect(sinkManagerMetrics.Emit().Metrics[2].Name).To(Equal("numberOfWebsocketSinks"))
		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))

		sink := &websocket.WebsocketSink{}
		sinkManagerMetrics.Inc(sink)

		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(1))

		sinkManagerMetrics.Dec(sink)

		Expect(sinkManagerMetrics.Emit().Metrics[0].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[1].Value).To(Equal(0))
		Expect(sinkManagerMetrics.Emit().Metrics[2].Value).To(Equal(0))
	})

})
