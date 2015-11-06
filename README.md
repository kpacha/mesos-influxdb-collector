mesos-influxdb-collector
=====

Lightweight mesos stats collector for influxdb

Since this collector is intended to be deployed as a [marathon](https://mesosphere.github.io/marathon) app, it comes with a *lifetime* param. This defines how long the collector will run until it dies, so marathon will re-launch it, allowing easy allocation optimizations. Check the [marathon/](https://github.com/kpacha/mesos-influxdb-collector/tree/master/marathon) folder for more details on how to launch it.

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
+ `MESOS_MASTER_HOST`
+ `MESOS_MASTER_PORT`
+ `MESOS_SLAVE_HOST`
+ `MESOS_SLAVE_PORT`
+ `MARATHON_HOST`
+ `MARATHON_PORT`
+ `COLLECTOR_LAPSE`
+ `COLLECTOR_LIFETIME`

## Dockerized version

Run the container with the default params (check the Dockerfile and overwrite whatever you need):

```
$ docker pull --name mesos-influxdb-collector \
    -e INFLUXDB_USER=admin \
    -e INFLUXDB_PWD=secret \
    -it --rm kpacha/mesos-influxdb-collector
```

Since the default value for `INFLUXDB_HOST` is `ìnfluxb`, you can link the collector to the influxdb container, dependeing on your environment.

```
$ docker run --name mesos-influxdb-collector \
    --link influxdb \
    -it --rm kpacha/mesos-influxdb-collector
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
  -Mmh string
      mesos master host (default "localhost")
  -Mmp int
      mesos master port (default 5050)
  -Msh string
      mesos slave host (default "localhost")
  -Msp int
      mesos slave port (default 5051)
  -d int
      die after N seconds (default 300)
  -l int
      sleep time between collections in seconds (default 1)
  -mh string
      marathon host (default "localhost")
  -mp int
      marathon port (default 8080)
```

This is the relation between those params and the environmnetal variables listed above.

Flag  | EnvVar
----  | ------
`Id`  | `INFLUXDB_DB`
`Ih`  | `INFLUXDB_HOST`
`Ip`  | `INFLUXDB_PORT`
`Mmh` | `MESOS_MASTER_HOST`
`Mmp` | `MESOS_MASTER_PORT`
`Msh` | `MESOS_SLAVE_HOST`
`Msp` | `MESOS_SLAVE_PORT`
`mh`  | `MARATHON_HOST`
`mp`  | `MARATHON_PORT`
`d`   | `COLLECTOR_LAPSE`
`l`   | `COLLECTOR_LIFETIME`

The credentials for the influxdb database are accepted just as env_var (`INFLUXDB_USER` & `INFLUXDB_PWD`)

## Testing environment

In order to do a quick test of the collector, you can use one of the available mesos test environments: [playa-mesos](https://github.com/mesosphere/playa-mesos) & [mesoscope](https://github.com/schibsted/mesoscope). The other components can be deployed with public containers. Replace the `$DOCKER_IP` and `$MESOS_HOST` with the correct values. If you are running the mesoscope env, `MESOS_HOST=$DOCKER_IP`. For the playa-mesos option, `MESOS_HOST=10.141.141.10`.

```
$ docker run --name influxdb -p 8083:8083 -p 8086:8086 \
    --expose 8090 --expose 8099 \
    -d tutum/influxdb
$ docker run --name grafana -p 3000:3000 \
    -e GF_SERVER_ROOT_URL="http://$DOCKER_IP" \
    -e GF_SECURITY_ADMIN_PASSWORD=secret \
    -d grafana/grafana
$ docker run --name mesos-influxdb-collector \
    --link influxdb \
    -e MESOS_HOST=$MESOS_HOST \
    -it --rm kpacha/mesos-influxdb-collector
```

The `grafana` folder contains several grafana dashboard definitions. Go to the grafana website (`http://$DOCKER_IP:3000/) and, after configuring the influxdb datasource, import them and start monitoring your mesos cluster.
