Running the mesos-influxdb-collector on marathon
====

The easiest way to deploy the mesos-influxdb-collector in your cluster is by setting the `instance` value to `2` and forcing the scheduller to deploy them in different slaves.

## Create a marathon app definition

Fix the [mesos-influxdb-collector.json](https://github.com/kpacha/mesos-influxdb-collector/blob/master/marathon/mesos-influxdb-collector.json) file. The critical params to customize are `MESOS_HOST` and `INFLUXDB_HOST`.

## Deploy it!

After updating the app definition, send it to your marathon instance (once again, replace the `$MARATHON_HOST` & `$MARATHON_PORT` with the right value for your environment)

```
$ curl -iH'Content-Type: application/json' -XPUT \
    -d@mesos-influxdb-collector.json \
    http://$MARATHON_HOST:$MARATHON_PORT/v2/apps/mesos-influxdb-collector
```
