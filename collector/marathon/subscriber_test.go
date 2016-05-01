package marathon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/kpacha/mesos-influxdb-collector/config"
	parser "github.com/kpacha/mesos-influxdb-collector/parser/marathon"
	"github.com/stretchr/testify/assert"
)

func TestNewMarathonEventsSubscriber(t *testing.T) {
	callbacks := []string{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.RequestURI, "/v2/eventSubscriptions") {
			if strings.Contains(r.RequestURI, "callbackUrl") {
				u, err := url.Parse(r.RequestURI)
				if err != nil {
					t.Fatalf("unable to parse the requested uri", err.Error())
				}
				m, _ := url.ParseQuery(u.RawQuery)
				callbacks = append(callbacks, m["callbackUrl"][0])
			}
			b, _ := json.Marshal(callbacks)
			fmt.Fprintf(w, `{"callbackUrls":%s}`, b)
		}
	}))
	defer ts.Close()

	host, port := splitHostPort(ts.URL)
	c := &config.Marathon{
		Server:     []config.Server{config.Server{Host: host, Port: port}},
		Host:       "127.0.0.1",
		Port:       9876,
		BufferSize: 1,
		Events:     true,
	}
	p := parser.MarathonEventsParser{}
	jobs := NewMarathonEventsSubscriber(c, p)

	<-time.After(1 * time.Second)

	var jsonStr = []byte(`{"eventType":"typeA", "taskStatus": "up", "appId": "supu"}`)
	req, _ := http.NewRequest("POST", "http://127.0.0.1:9876/marathon", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal("Error sending the test event:", err.Error())
	}
	defer resp.Body.Close()

	<-time.After(100 * time.Millisecond)
	j := <-jobs

	assert.Equal(t, 1, len(j), "Unexpected number of points")
	if 0 == len(j) {
		return
	}
	assert.Equal(t, "marathon-event", j[0].Measurement, "Unexpected measurement name")
	assert.Equal(t, "127.0.0.1", j[0].Tags["node"], "Unexpected node name")
	assert.Equal(t, "typeA", j[0].Tags["type"], "Unexpected type name")
	assert.Equal(t, "up", j[0].Tags["status"], "Unexpected status name")
	assert.Equal(t, "supu", j[0].Tags["appId"], "Unexpected appId name")
	assert.Equal(t, 1, j[0].Fields["value"], "Unexpected field value")
}

func splitHostPort(original string) (string, int) {
	u, _ := url.Parse(original)
	host, stringPort, _ := net.SplitHostPort(u.Host)
	port, _ := strconv.Atoi(stringPort)
	return host, port
}
