mesosDNS {
	domain = "mesos"
	marathon = true
	host = "master.mesos"
	port = 8123
	checkLapse = 30
}
influxdb {
	host = "influxdb.marathon.mesos"
	port = 8086
	db = "mesos"
}
lapse=5
dieAfter = 300
