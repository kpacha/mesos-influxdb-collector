mesosDNS {
	domain = "mesos"
	marathon = true
	host = "master.mesos"
	port = 8123
}
influxdb {
	host = "influxdb.marathon.mesos"
	port = 8086
	db = "mesos"
	checkLapse = 30
}
lapse=5
dieAfter = 300
