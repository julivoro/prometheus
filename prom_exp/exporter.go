package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	fileCountExporter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_folder_files_count",
			Help: "My folder file count exporter",
		},
		[]string{"directory"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(fileCountExporter)
}

func recursiveDirFileCounter(directory string) {

	readDir, _ := os.Open(directory)
	readDirFiles, _ := readDir.Readdir(0)

	for idx := range readDirFiles {
		file := readDirFiles[idx]
		if file.IsDir() {
			recursiveDirFileCounter(directory + string(os.PathSeparator) + file.Name())
		} else {
			fileCountExporter.With(prometheus.Labels{"directory": directory}).Inc()
		}
	}
}

func main() {

	myFolder := flag.String("dir", "E:"+string(os.PathSeparator), "Directory where to count files")
	httpHostPort := flag.String("listen", ":8080", "Host and port of prom http server, default :8080")
	flag.Parse()

	go recursiveDirFileCounter(*myFolder)

	//run http server
	//handle metrics, including newly created my_folder_files_count
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*httpHostPort, nil))
}
