package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

const (
	namespace = "grpc"
	subsystem = "client"
)

func summary(name, help string, labels []string) *prometheus.SummaryVec {
	return promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  namespace,
		Subsystem:  subsystem,
		Name:       name,
		Help:       help,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		MaxAge:     time.Minute,
	}, labels)
}

func counter(name, help string, labels []string) *prometheus.CounterVec {
	return promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	}, labels)
}

func histogram(name, help string, labels []string, buckets []float64) *prometheus.HistogramVec {
	return promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labels)
}

func gauge(name, help string, labels []string) *prometheus.GaugeVec {
	return promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	}, labels)
}

var (
	http2RespDataFrameLen = summary(
		"http2_resp_data_frame_len",
		"http2 response data frame len",
		[]string{"conn_id"})

	http2ReqDataFrameLen = summary(
		"http2_req_data_frame_len",
		"http2 req data frame len",
		[]string{"conn_id"})

	waitingStreamsNum = gauge(
		"waiting_streams_num",
		"waiting streams number",
		[]string{"conn_id"})

	streamsQuota = gauge(
		"streams_quota",
		"streams quota",
		[]string{"conn_id"})

	activeStreamsNum = gauge(
		"active_streams_num",
		"active streams number",
		[]string{"conn_id"})

	activeStreamItemsNum = gauge(
		"active_stream_items_num",
		"active stream items num",
		[]string{"conn_id"})

	oriReqItemDataLen = gauge(
		"ori_req_item_data_len",
		"original req item data len",
		[]string{"conn_id"})

	streamBytesOutStanding = gauge(
		"stream_bytes_out_standing",
		"stream bytes out standing",
		[]string{"conn_id", "stream_id"})

	sendQuota = gauge(
		"send_quota",
		"send quota",
		[]string{"conn_id"})

	sendBufSize = gauge(
		"send_buf_size",
		"send buffer size",
		[]string{"conn_id"})

	recvBufSize = gauge(
		"recv_buf_size",
		"recv buf size",
		[]string{"conn_id"})
)

func RecordSendBufSize(connId string, n uint) {
	if n > 0 {
		sendBufSize.WithLabelValues(connId).Set(float64(n))
	}
}

func RecordRecvBufSize(connId string, n uint) {
	if n > 0 {
		recvBufSize.WithLabelValues(connId).Set(float64(n))
	}
}

func RecordStreamBytesOutStanding(connId string, streamId string, n uint) {
	if n > 0 {
		streamBytesOutStanding.WithLabelValues(connId, streamId).Set(float64(n))
	}
}

func RecordSendQuota(connId string, n uint) {
	if n > 0 {
		sendQuota.WithLabelValues(connId).Set(float64(n))
	}
}

func RecordOriReqItemDataLen(connId string, n uint) {
	if n > 0 {
		oriReqItemDataLen.WithLabelValues(connId).Set(float64(n))
	}
}

func RecordHttp2ReqDataFrameLen(connId string, n uint) {
	if n > 0 {
		http2ReqDataFrameLen.WithLabelValues(connId).Observe(float64(n))
	}
}

func RecordHttp2RespDataFrameLen(connId string, n uint) {
	if n > 0 {
		http2RespDataFrameLen.WithLabelValues(connId).Observe(float64(n))
	}
}

func RecordWaitingStreamsNum(connId string, n uint) {
	if n > 0 {
		waitingStreamsNum.WithLabelValues(connId).Set(float64(n))
	}
}

func RecordStreamsQuota(connId string, n uint) {
	if n > 0 {
		streamsQuota.WithLabelValues(connId).Set(float64(n))
	}
}

func RecordActiveStreamsNum(connId string, n uint) {
	if n > 0 {
		activeStreamsNum.WithLabelValues(connId).Set(float64(n))
	}
}

func RecordActiveItemsNum(connId string, n uint) {
	if n > 0 {
		activeStreamItemsNum.WithLabelValues(connId).Set(float64(n))
	}
}
