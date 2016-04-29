mesosDNS {
	domain = "ws.ideaportriga.lv"
	marathon = true
	host = "mesos1.ideaportriga.lv"
	port = 8123
}
influxdb {
	host = "influx.ws.ideaportriga.lv"
	port = 8086
	db = "metrics"
	checkLapse = 30
}
marathon {
	port = 8080
}
lapse=5
dieAfter = 300
haproxy {
	endPoint = "haproxy?stats"
	port = 9090
}
