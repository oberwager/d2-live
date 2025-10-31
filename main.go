package main

import (
	"log"
	"net/http"

	statsd "github.com/DataDog/datadog-go/v5/statsd"
	"github.com/husobee/vestigo"
	"github.com/oberwager/d2-live/internal/handlers"
	ctxlog "oss.terrastruct.com/d2/lib/log"
)

var Version string

func main() {
	metricsClient, err := statsd.New("statsd-exporter-service.hyperion.svc.cluster.local:8125",
		statsd.WithTags([]string{"env:prod"}),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctxlog.Init()

	c := handlers.Controller{
		Metrics: metricsClient,
		Version: Version,
	}

	router := vestigo.NewRouter()

	router.Get("/", c.GetD2SVGHandler, c.StatsdMiddleware)

	router.Get("/info", c.GetInfoHandler, c.StatsdMiddleware)

	router.Get("/svg/:encodedD2", c.GetD2SVGHandler, c.StatsdMiddleware)

	log.Fatal(http.ListenAndServe(":8090", router))
}
