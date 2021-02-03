# prometheus
Learning Prometheus

prom_exp
Usage of Prometheus go client for counting files in the directory, and exposing them to prometheus metrics by format:
my_folder_files_count{directory="path"} count

0) install prometheus go client 
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp

Start client:
1) go run exporter.go
Fetch metrics from server:
2) curl http://localhost:8080/metrics
