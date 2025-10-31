package handlers

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/sloghuman"
	"github.com/husobee/vestigo"
	"github.com/oberwager/d2-live/internal/urlenc"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/log"
	"oss.terrastruct.com/d2/lib/textmeasure"
	"oss.terrastruct.com/util-go/go2"
)

func DiscardSlog(ctx context.Context) context.Context {
	return log.With(ctx, slog.Make(sloghuman.Sink(io.Discard)))
}

func (c *Controller) GetD2SVGHandler(rw http.ResponseWriter, req *http.Request) {
	ctx := DiscardSlog(req.Context())

	// First, try to get encodedD2 from the path.
	urlencoded := vestigo.Param(req, "encodedD2")

	// If encodedD2 is not found in the path, look for the ?script= variable.
	if urlencoded == "" {
		urlencoded = req.URL.Query().Get("script")
	}

	// If still not found, return an error.
	if urlencoded == "" {
		http.Error(rw, "encodedD2 or script parameter not provided", http.StatusBadRequest)
		return
	}

	// Get theme if provided
	themeStr := req.URL.Query().Get("theme")
	var theme int64
	var err error
	if themeStr != "" {
		theme, err = strconv.ParseInt(themeStr, 10, 64)
		if err != nil {
			http.Error(rw, "Invalid theme parameter", http.StatusBadRequest)
			return
		}
	} else {
		// Use a default theme if none is provided
		theme = d2themescatalog.NeutralDefault.ID
	}

	// Emit complexity metric
	c.Metrics.Histogram("d2-live.complexity", float64(len(urlencoded)), []string{}, 1)

	svg, err := c.handleGetD2SVG(ctx, urlencoded, theme)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "image/svg+xml")
	rw.Write(svg)
}

func (c *Controller) handleGetD2SVG(ctx context.Context, encoded string, theme int64) ([]byte, error) {
	decoded, err := urlenc.Decode(encoded)
	if err != nil {
		return nil, errors.New("Invalid Base64 data.")
	}

	ruler, _ := textmeasure.NewRuler()
	layoutResolver := func(engine string) (d2graph.LayoutGraph, error) {
		return d2dagrelayout.DefaultLayout, nil
	}
	renderOpts := &d2svg.RenderOpts{
		Pad:     go2.Pointer(int64(5)),
		ThemeID: &theme,
	}
	compileOpts := &d2lib.CompileOptions{
		LayoutResolver: layoutResolver,
		Ruler:          ruler,
	}

	diagram, _, _ := d2lib.Compile(ctx, decoded, compileOpts, renderOpts)

	// Render to SVG
	out, err := d2svg.Render(diagram, renderOpts)
	if err != nil {
		return nil, errors.New("Invalid D2 data.")
	}
	return out, nil
}
