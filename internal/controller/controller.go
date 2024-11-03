package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/ribeirosaimon/shortify-read/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Start() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/health", NewHealth)
	http.HandleFunc("/", NewUrlRecord().NewUrlRecord)

	http.Handle("/metrics", promhttp.Handler())

}

func getIDFromURL(path string) (string, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid URL")
	}
	return parts[1], nil
}
