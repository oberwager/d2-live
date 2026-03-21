package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/husobee/vestigo"
	"github.com/oberwager/d2-live/internal/handlers"
	ctxlog "oss.terrastruct.com/d2/lib/log"
)

var Version string

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctxlog.Init()

	c := handlers.Controller{
		Logger:  logger,
		Version: Version,
	}

	router := vestigo.NewRouter()

	router.Get("/", c.GetD2SVGHandler, c.LoggingMiddleware)
	router.Get("/info", c.GetInfoHandler, c.LoggingMiddleware)
	router.Get("/svg/:encodedD2", c.GetD2SVGHandler, c.LoggingMiddleware)

	log.Fatal(http.ListenAndServe(":8090", router))
}
