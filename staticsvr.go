// author: wongoo
// since: 2019-10-08
package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/vogo/logger"
	"os"
	"strconv"
	"time"
)

func main() {
	port := 9999
	staticDir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}

	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			logger.Fatal(err)
		}
	}
	if len(os.Args) > 2 {
		staticDir = os.Args[2]
	}

	logger.Infof("serve static on port %d at directory %s", port, staticDir)

	fs := &fasthttp.FS{
		Root:               staticDir,
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: false,
		Compress:           false,
		AcceptByteRange:    false,
		CacheDuration:      time.Minute,
	}

	server := &fasthttp.Server{
		Handler:            fs.NewRequestHandler(),
		ReadTimeout:        3600 * time.Second,
		WriteTimeout:       3600 * time.Second,
		MaxRequestBodySize: 1 << 20,
	}

	logger.Info(server.ListenAndServe(fmt.Sprintf(":%d", port)))
}
