mesos-influxdb-collector
=====

Lightweight mesos stats collector for influxdb

Since this collector is intended to be deployed as a [marathon](https://mesosphere.github.io/marathon) app, it comes with a *lifetime* param. This defines how long the collector will run until it dies, so marathon will re-launch it, allowing easy allocation optimizations. Check the [fixtures/marathon/](https://github.com/kpacha/mesos-influxdb-collector/tree/master/fixtures/marathon) folder for more details on how to launch it.

# Goals

+ Discover the mesos cluster through `mesos-dns`
+ Collect the mesos leader stats
+ Collect the mesos master stats
+ Collect the mesos slave stats
+ Collect the marathon master stats
+ Collect the mesos executors stats
+ Collect the marathon events (experimental)
+ Collect the haproxy stats (experimental)
+ Collect the chronos task stats (TODO)

# Installing

Docker images are available at [docker hub](https://hub.docker.com/r/kpacha/mesos-influxdb-collector). Just pull one of the available images with:

```
$ docker pull kpacha/mesos-influxdb-collector:latest-min
```

The `-min` versions are images with just the binary and a config file.

**Be careful with the `latest` and the `latest-min` versions. They could have experimental features, not battle-proven yet**

Alternatively, get the last released binary from the [releases](https://github.com/kpacha/mesos-influxdb-collector/releases) section.


Finally, if you love the hard way and have Go installed:

```
$ go get github.com/kpacha/mesos-influxdb-collector
$ cd $GOPATH/src/github.com/kpacha/mesos-influxdb-collector
$ make
```

# Integration with `mesos-dns`

The `mesos-influxdb-collector` is able to discover all your mesos nodes (masters and slaves) and the marathon master using the REST API exposed by the [mesos-dns](http://mesosphere.github.io/mesos-dns/) service. Check the next section for details.

# Configuration

The collector implements the 12 factor app methodology, so it has several ways to be configured: environmental vars, flags and a config file:

+ The config file is where all the defaults should be placed
+ Flags could be used in order to customize the format and the path of the configuration file
+ Environmental variables should be use to keep the secrets away from the repository and/or the container image

Flags also could be passed as environmental vars:

+ `FORMAT`
+ `CONFIG_PATH`
+ `CONFIG_FILE`

The config file should contain the list of nodes to monitor or the details about the `mesos-dns` service among these other params:

+ *Lapse*: time between consecutive collections. Default: 30 seconds
+ *DieAfter*: duration of the running instance. Default: 1 hour

### MesosDNS

Optional. Add it if you have a `mesos-dns` service running in your mesos cluster.

```
   "mesosDNS":{  
      "domain":"mesos", // the domain used by the mesos-dns service
      "marathon":true, // resolve marathon master
      "host":"slave.mesos", // host of the mesos-dns service
      "port":8123 // port of the REST API
   }
```

### InfluxDB

Required.

```
   "influxdb":{  
      "host":"influxdb.marathon.mesos", // host of the influxdb instance
      "port":8086, // port of the REST API
      "db":"mesos", // name of the database to use
      "checkLapse":30 // ping frequency
   }
```

### Mesos masters

Optional. For manual definition of some (or all) mesos masters, use the `Master` struct:

```
    "master": [
        {
            "host": "localhost",
            "port": 5050,
            "leader": true
        },
        {
            "host": "localhost",
            "port": 5051
        }
    ]
```

### Mesos slaves

Optional. For manual definition of some (or all) mesos slave, use the `Slave` struct:

```
    "slave": [
        {
            "host": "slave0.example.com",
            "port": 5051
        },
        {
            "host": "slave1.example.com",
            "port": 5051
        },
        {
            "host": "slave2.example.com",
            "port": 5051
        }
    ]
```

### Marathon instances

Optional. For manual definition of some (or all) marathon instances, use the `Marathon` struct:

```
    "marathon": {
        "server": [
            {
                "host": "marathon1",
                "port": 8080
            },
            {
                "host": "marathon2",
                "port": 8080
            }
        ],
        "events": true,
        "host": "$HOST",
        "port": 8088,
        "bufferSize": 10000
    }
```

### HAProxy

Optionsl. If you want to also collect the HAProxy stats, add a `haproxy` section to your config file

```
   "haproxy":{  
      "user":"admin",
      "password":"admin",
      "port":9090,
      "endPoint":"haproxy?stats"
   }
```

Check [`config/configuration_test.go`](https://github.com/kpacha/mesos-influxdb-collector/blob/master/config/configuration_test.go), [`fixtures/`](https://github.com/kpacha/mesos-influxdb-collector/tree/master/fixtures) and [`conf.json`](https://github.com/kpacha/mesos-influxdb-collector/blob/master/conf.json) for examples.

# Running

## Dockerized version

Run the container with the default params:

```
$ docker pull --name mesos-influxdb-collector \
    -it --rm kpacha/mesos-influxdb-collector
```

If you need to customize something, there are some alternatives:

### 1: Config file

Just copy the `conf.json`, make your changes and link it as a volume:

```
$ docker pull --name mesos-influxdb-collector \
    -v /path/to/my/custom/conf.json:/tmp/conf.json \
    -it --rm kpacha/mesos-influxdb-collector -d /tmp/
```

Tip: if you link your config file to `/go/src/github.com/kpacha/mesos-influxdb-collector/conf.json` you don't need to worry about that flag!

### 2: ENV VARS

Override the credentials so you secrets won't be stored neither the repo nor the container image. This just works for root config values and nested fields from the HAProxy and the Influxdb sections. These envvars use the `MIC_` namespace.

```
$ docker pull --name mesos-influxdb-collector \
    -e MIC_INFLUXDB_HOST=influxdb.example.com \
    -v /path/to/my/custom/conf.json:/tmp/conf.json \
    -it --rm kpacha/mesos-influxdb-collector -d /tmp/
```

### 3: Flags

Use the flags accepted by the binary and defined below. Remeber to set them as commands to the defined entrypoint for the docker container.

```
$ docker pull --name mesos-influxdb-collector \
    -v /path/to/my/custom/conf.hcl:/tmp/conf.hcl \
    -it --rm kpacha/mesos-influxdb-collector -d /tmp/ -f hcl
```

## Binary version

```
$ ./mesos-influxdb-collector -h
Usage of ./mesos-influxdb-collector:
  -c string
      name of the config file (default "conf")
  -d string
      path to the config folder (default ".")
  -dns
      enable mesos-dns (default true)
  -f string
      config format (default "json")
```

This is the relation between those params and the environmnetal variables listed above.

Flag | EnvVar
---- | ------
`c`  | `CONFIG_FILE`
`d`  | `CONFIG_PATH`
`f`  | `FORMAT`

# Grafana dashboards

The [fixtures/grafana](https://github.com/kpacha/mesos-influxdb-collector/tree/master/fixtures/grafana) folder contains several grafana dashboard definitions. Go to the grafana website and, after configuring the influxdb datasource, import them and start monitoring your mesos cluster.
