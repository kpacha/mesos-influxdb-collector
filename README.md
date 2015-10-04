mesos-influxdb-collector
=====

Lightweight mesos stats collector for influxdb

# Installing

Docker images are available at [docker hub](https://hub.docker.com/r/kpacha/mesos-influxdb-collector).

Alternatively, if you have Go installed:

```
$ go get github.com/kpacha/mesos-influxdb-collector
```

# Running

```
./mesos-influxdb-collector -h
Usage of ./mesos-influxdb-collector:
  -Id string
      influxdb database (default "mesos")
  -Ih string
      influxdb host (default "localhost")
  -Ip int
      influxdb port (default 8086)
  -Mh string
      mesos host (default "localhost")
  -Mp int
      mesos port (default 5050)
  -d int
      die after N seconds (default 300)
  -l int
      sleep time between collections in seconds (default 1)
```

The binary also accepts those params as environmnetal variables.

Flag | EnvVar
---- | ------
`Id` | `INFLUXDB_DB`
`Ih` | `INFLUXDB_HOST`
`Ip` | `INFLUXDB_PORT`
`Mh` | `MESOS_HOST`
`Mp` | `MESOS_PORT`
`d`  | `COLLECTOR_LAPSE`
`l`  | `COLLECTOR_LIFETIME`

The credentials for the influxdb database are accepted just as env_var (`INFLUX_USER` & `INFLUX_PWD`)
