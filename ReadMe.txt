//Run node-exporter
docker run -itd --name node-exporter --net="host" --pid="host" -v "/:/host:ro,rslave" quay.io/prometheus/node-exporter:latest --path.rootfs=/host

//Run cadvisor
docker run --volume=/:/rootfs:ro --volume=/var/run:/var/run:rw --volume=/sys:/sys:ro --volume=/var/lib/docker/:/var/lib/docker:ro --publish=8080:8080 --detach=true --name=cadvisor google/cadvisor:latest

//Run prometheus
docker run -itd -p 9090:9090 --name prometheus -v /root/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml -v /root/prometheus/rules/hoststats-alert.rules:/etc/prometheus/rules/hoststats-alert.rules prom/prometheus

//Run alertmanager
docker run -itd -p 9093:9093 --name alertmanager -v /root/prometheus/alertmanager/alertmanager.yml:/etc/alertmanager/alertmanager.yml prom/alertmanager

//Run webhook
./webhook/webhook

//Run grafana
docker run -d -p 3000:3000 grafana/grafana
