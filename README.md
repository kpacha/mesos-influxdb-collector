mesos-influxdb-collector
=====

Lightweight mesos stats collector for influxdb

# Installing

Docker images are available at [docker hub](https://hub.docker.com/r/kpacha/mesos-influxdb-collector). Just pull the latest image available with:

```
$ docker pull kpacha/mesos-influxdb-collector:latest
```

Alternatively, if you have Go installed:

```
$ go get github.com/kpacha/mesos-influxdb-collector
```

# Running

The collector use these environmental vars:

+ `INFLUXDB_DB`
+ `INFLUXDB_HOST`
+ `INFLUXDB_PORT`
+ `INFLUXDB_USER`
+ `INFLUXDB_PWD`
+ `MESOS_HOST`
+ `MESOS_PORT`
+ `COLLECTOR_LAPSE`
+ `COLLECTOR_LIFETIME`

## Dockerized version

Run the container with the default params (check the Dockerfile and overwrite whatever you need):

```
$ docker pull --name mesos-influxdb-collector \
    -e INFLUX_USER=admin \
    -e INFLUX_PWD=secret \
    -it --rm kpacha/mesos-influxdb-collector
```

Since the default value for `INFLUXDB_HOST` is `ìnfluxb`, you can link the collector to the influxdb container, dependeing on your environment.

```
$ docker pull -it --rm --name mesos-influxdb-collector --link influxdb kpacha/mesos-influxdb-collector
```

## Binary version

```
$ ./mesos-influxdb-collector -h
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

This is the relation between those params and the environmnetal variables listed above.

Flag | EnvVar
---- | ------
`Id` | `INFLUXDB_DB`
`Ih` | `INFLUXDB_HOST`
`Ip` | `INFLUXDB_PORT`
`Mh` | `MESOS_HOST`
`Mp` | `MESOS_PORT`
`d`  | `COLLECTOR_LAPSE`
`l`  | `COLLECTOR_LIFETIME`

The credentials for the influxdb database are accepted just as env_var (`INFLUXDB_USER` & `INFLUXDB_PWD`)
