// author: wongoo
// since: 2019-10-08
package main

import (
	"fmt"
	"github.com/vogo/logger"
	"net/http"
	"os"
	"strconv"
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

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)

	logger.Info(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
