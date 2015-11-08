Running the mesos-influxdb-collector on marathon
====

The easiest way to deploy the mesos-influxdb-collector in your cluster is by setting the `instance` value to `2` and forcing the scheduller to deploy them in different slaves.

## Create a marathon app definition

Fix the [mesos-influxdb-collector.json](https://github.com/kpacha/mesos-influxdb-collector/blob/master/marathon/mesos-influxdb-collector.json) file. The critical params to customize are `MESOS_HOST` and `INFLUXDB_HOST.

## Deploy it!

After updating the app definition, send it to your marathon instance (once again, replace the `$MARATHON_HOST` & `$MARATHON_PORT` with the right value for your environment)

```
$ curl -iH'Content-Type: application/json' -XPUT \
    -d@mesos-influxdb-collector.json \
    http://$MARATHON_HOST:$MARATHON_PORT/v2/apps/mesos-influxdb-collector
```

curl -iXPUT -H"Content-Type: application/json" 172.28.128.3:8080/v2/apps/mesos-influxdb-collector -d'{
  "id": "mesos-influxdb-collector",
  "cpus": 0.1,
  "mem": 64.0,
  "instances": 2,
  "container": {
    "type": "DOCKER",
    "docker": {
      "image": "kpacha/mesos-influxdb-collector"
    }
  },
  "env": {
    "MESOS_HOST": "172.17.42.1",
    "INFLUXDB_HOST": "$HOST",
    "COLLECTOR_LIFETIME": "300"
  },
  "constraints": [
    ["hostname", "UNIQUE"]
  ],
  "backoffSeconds": 1,
  "backoffFactor": 1.15,
  "maxLaunchDelaySeconds": 300,
  "upgradeStrategy": {
    "minimumHealthCapacity": 1,
    "maximumOverCapacity": 1
  }
}'

curl -iXPUT -H"Content-Type: application/json" 172.28.128.3:8080/v2/apps/influxdb -d'{
  "id": "influxdb",
  "cpus": 0.4,
  "mem": 400.0,
  "instances": 1,
  "container": {
    "type": "DOCKER",
    "docker": {
      "image": "tutum/influxdb",
      "network": "BRIDGE",
      "portMappings": [
        { "containerPort": 8083, "hostPort": 0, "servicePort": 8083, "protocol": "tcp" },
        { "containerPort": 8086, "hostPort": 0, "servicePort": 8086, "protocol": "tcp" }
      ]
    }
  },
  "env": {
    "MESOS_HOST": "10.0.2.15",
    "INFLUXDB_HOST": "10.0.2.15",
    "COLLECTOR_LIFETIME": "300"
  },
  "constraints": [
    ["hostname", "UNIQUE"]
  ],
  "backoffSeconds": 1,
  "backoffFactor": 1.15,
  "maxLaunchDelaySeconds": 300,
  "upgradeStrategy": {
    "minimumHealthCapacity": 1,
    "maximumOverCapacity": 1
  }
}'

curl -iXPUT -H"Content-Type: application/json" 172.28.128.3:8080/v2/apps/grafana -d'{
  "id": "grafana",
  "cpus": 0.2,
  "mem": 400.0,
  "instances": 1,
  "container": {
    "type": "DOCKER",
    "docker": {
      "image": "grafana/grafana",
      "network": "BRIDGE",
      "portMappings": [
        { "containerPort": 3000, "hostPort": 0, "servicePort": 3000, "protocol": "tcp" }
      ]
    }
  },
  "env": {
    "INFLUXDB_HOST": "10.0.2.15",
    "COLLECTOR_LIFETIME": "300"
  },
  "constraints": [
    ["hostname", "UNIQUE"]
  ],
  "backoffSeconds": 1,
  "backoffFactor": 1.15,
  "maxLaunchDelaySeconds": 300,
  "upgradeStrategy": {
    "minimumHealthCapacity": 1,
    "maximumOverCapacity": 1
  }
}'