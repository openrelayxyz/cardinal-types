package ship

import (
	"net"
	_ "net/http/pprof"
	"fmt"
	"github.com/openrelayxyz/cardinal-types/metrics"
		log "github.com/inconshreveable/log15"
		"github.com/pubnub/go-metrics-statsd"
	"strconv"
	"time"
	"github.com/savaki/cloudmetrics"
)

func StatsD(port string, address string, interval time.Duration, prefix string, minor bool) {
		addr := "127.0.0.1:" + port
		if address != "" {
			addr = fmt.Sprintf("%v:%v", address, port)
		}
		udpAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			log.Error("Invalid Address. Statsd will not be configured.", "error", err.Error())
		} else {
			if interval == 0 {
				interval = time.Second
			}
			go statsd.StatsD(
				metrics.MajorRegistry,
				interval,
				prefix,
				udpAddr,
			)
			if minor {
				go statsd.StatsD(
					metrics.MinorRegistry,
					interval,
					prefix,
					udpAddr,
				)
			}
		}
}

func CloudWatch(namespace string, cwDimensions map[string]string, chainid int64, interval time.Duration, percentiles []float64, minor bool) {
		if namespace == "" {
			namespace = "Cardinal"
		}
		dimensions := []string{}
		for k, v := range cwDimensions {
			dimensions = append(dimensions, k, v)
		}
		if len(dimensions) == 0 {
			dimensions = append(dimensions, "chainid", strconv.Itoa(int(chainid)))
		}
		cwcfg := []func(*cloudmetrics.Publisher){
			cloudmetrics.Dimensions(dimensions...),
		}
		if interval > 0 {
			cwcfg = append(cwcfg, cloudmetrics.Interval(time.Duration(interval) * time.Second))
		}
		if len(percentiles) > 0 {
			cwcfg = append(cwcfg, cloudmetrics.Percentiles(percentiles))
		}
		go cloudmetrics.Publish(metrics.MajorRegistry,
			namespace,
			cwcfg...
		)
		if minor {
			go cloudmetrics.Publish(metrics.MinorRegistry,
				namespace,
				cwcfg...
			)
	}
}
