package config

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNSResolverMesosMasters(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/v1/hosts/leader.mesos." {
			fmt.Fprintln(w, `[{"ip":"123.123.123.123"}]`)
		}
		if r.RequestURI == "/v1/hosts/master.mesos." {
			fmt.Fprintln(w, `[{"ip":"123.123.123.123"},{"ip":"123.123.123.124"}]`)
		}
	}))
	defer ts.Close()

	host, port := splitHostPort(ts.URL)

	c := Config{MesosDNS: &MesosDNS{Domain: "mesos", Marathon: false, Host: host, Port: port}}
	dnsResolver := NewDNSResolver(&c)
	if err := dnsResolver.resolveMesosMasters(); err != nil {
		fmt.Println("Error processing the mocked response:", err.Error())
		t.Fail()
	}

	assert.Equal(t, 0, len(c.Slave), "Unexpected number of slaves")
	assert.Nil(t, c.Marathon)
	assert.Equal(t, 2, len(c.Master), "Unexpected number of masters")
	assert.Equal(t, "123.123.123.123", c.Master[0].Host, "Wrong master host")
	assert.Equal(t, 5050, c.Master[0].Port, "Wrong master port")
	assert.Equal(t, true, c.Master[0].Leader, "The first master must be also the leader")
	assert.Equal(t, "123.123.123.124", c.Master[1].Host, "Wrong master host")
	assert.Equal(t, 5050, c.Master[1].Port, "Wrong master port")
	assert.Equal(t, false, c.Master[1].Leader, "The second master shouldn't be the leader")
}

func TestDNSResolverMesosSlaves(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()

	host, port := splitHostPort(ts.URL)

	c := Config{MesosDNS: &MesosDNS{Domain: "mesos", Marathon: false, Host: host, Port: port}}
	dnsResolver := NewDNSResolver(&c)
	if err := dnsResolver.resolveMesosSlaves(); err != nil {
		fmt.Println("Error processing the mocked response:", err.Error())
		t.Fail()
	}

	assert.Equal(t, 0, len(c.Master), "Unexpected number of master")
	assert.Nil(t, c.Marathon)
	assert.Equal(t, 2, len(c.Slave), "Unexpected number of slaves")
	assert.Equal(t, "123.123.123.123", c.Slave[0].Host, "Wrong slave host")
	assert.Equal(t, 5051, c.Slave[0].Port, "Wrong slave port")
	assert.Equal(t, "123.123.123.124", c.Slave[1].Host, "Wrong slave host")
	assert.Equal(t, 5051, c.Slave[1].Port, "Wrong slave port")
}

func TestDNSResolverMarathonInstances(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()

	host, port := splitHostPort(ts.URL)

	c := Config{MesosDNS: &MesosDNS{Domain: "mesos", Marathon: true, Host: host, Port: port}}
	fmt.Println(c.MesosDNS)
	dnsResolver := NewDNSResolver(&c)
	if err := dnsResolver.resolveMarathon(); err != nil {
		fmt.Println("Error processing the mocked response:", err.Error())
		t.Fail()
	}

	assert.Equal(t, 0, len(c.Master), "Unexpected number of master")
	assert.Equal(t, 0, len(c.Slave), "Unexpected number of slaves")
	assert.NotNil(t, c.Marathon)
	assert.Equal(t, "123.123.123.123", c.Marathon.Server[0].Host, "Wrong marathon instance host")
	assert.Equal(t, 8080, c.Marathon.Server[0].Port, "Wrong marathon instance port")
	assert.Equal(t, "123.123.123.124", c.Marathon.Server[1].Host, "Wrong marathon instance host")
	assert.Equal(t, 8080, c.Marathon.Server[1].Port, "Wrong marathon instance port")
}

func TestDNSResolver(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/v1/hosts/leader.mesos." {
			fmt.Fprintln(w, `[{"ip":"123.123.123.123"}]`)
		} else {
			fmt.Fprintln(w, `[{"ip":"123.123.123.123"},{"ip":"123.123.123.124"}]`)
		}
	}))
	defer ts.Close()

	host, port := splitHostPort(ts.URL)

	c := Config{MesosDNS: &MesosDNS{Domain: "mesos", Marathon: true, Host: host, Port: port}}
	dnsResolver := NewDNSResolver(&c)
	if err := dnsResolver.resolve(); err != nil {
		fmt.Println("Error processing the mocked response:", err.Error())
		t.Fail()
	}

	assert.Equal(t, 2, len(c.Master), "Unexpected number of masters")
	assert.Equal(t, "123.123.123.123", c.Master[0].Host, "Wrong master host")
	assert.Equal(t, 5050, c.Master[0].Port, "Wrong master port")
	assert.Equal(t, true, c.Master[0].Leader, "The first master must be also the leader")
	assert.Equal(t, "123.123.123.124", c.Master[1].Host, "Wrong master host")
	assert.Equal(t, 5050, c.Master[1].Port, "Wrong master port")
	assert.Equal(t, false, c.Master[1].Leader, "The second master shouldn't be the leader")
	assert.Equal(t, 2, len(c.Slave), "Unexpected number of slaves")
	assert.Equal(t, "123.123.123.123", c.Slave[0].Host, "Wrong slave host")
	assert.Equal(t, 5051, c.Slave[0].Port, "Wrong slave port")
	assert.Equal(t, "123.123.123.124", c.Slave[1].Host, "Wrong slave host")
	assert.Equal(t, 5051, c.Slave[1].Port, "Wrong slave port")
	assert.NotNil(t, c.Marathon)
	assert.Equal(t, "123.123.123.123", c.Marathon.Server[0].Host, "Wrong marathon instance host")
	assert.Equal(t, 8080, c.Marathon.Server[0].Port, "Wrong marathon instance port")
	assert.Equal(t, "123.123.123.124", c.Marathon.Server[1].Host, "Wrong marathon instance host")
	assert.Equal(t, 8080, c.Marathon.Server[1].Port, "Wrong marathon instance port")
}

func buildTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[{"ip":"123.123.123.123"},{"ip":"123.123.123.124"}]`)
	}))
}

func splitHostPort(original string) (string, int) {
	u, _ := url.Parse(original)
	host, stringPort, _ := net.SplitHostPort(u.Host)
	port, _ := strconv.Atoi(stringPort)
	return host, port
}
