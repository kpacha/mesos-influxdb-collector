package haproxy

import (
	"fmt"
	"io"
	"strings"
)

func ExampleHAProxy() {
	parser := NewHAProxy("testing-node")
	content := newTestingReader(`# pxname,svname,qcur,qmax,scur,smax,slim,stot,bin,bout,dreq,dresp,ereq,econ,eresp,wretr,wredis,status,weight,act,bck,chkfail,chkdown,lastchg,downtime,qlimit,pid,iid,sid,throttle,lbtot,tracked,type,rate,rate_lim,rate_max,check_status,check_code,check_duration,hrsp_1xx,hrsp_2xx,hrsp_3xx,hrsp_4xx,hrsp_5xx,hrsp_other,hanafail,req_rate,req_rate_max,req_tot,cli_abrt,srv_abrt,comp_in,comp_out,comp_byp,comp_rsp,lastsess,last_chk,last_agt,qtime,ctime,rtime,ttime,
http-in,FRONTEND,,,1,2,2000,12,8393,141134,0,0,0,,,,,OPEN,,,,,,,,,1,2,0,,,,0,1,0,2,,,,0,9,0,2,0,0,,1,1,12,,,0,0,0,0,,,,,,,,
::groupA::appA-cluster,::groupA::appA-slave06.mesos-31373,0,0,0,0,,0,0,0,,0,,0,0,0,0,no check,1,1,0,,,,,,1,3,1,,0,,2,0,,0,,,,0,0,0,0,0,0,0,,,,0,0,,,,,-1,,,0,0,0,0,
::groupA::appA-cluster,BACKEND,0,0,0,0,200,0,0,0,0,0,,0,0,0,0,UP,1,1,0,,0,49,0,,1,3,0,,0,,1,0,,0,,,,0,0,0,0,0,0,,,,,0,0,0,0,0,0,-1,,,0,0,0,0,
::groupB::appB-cluster,::groupB::appB-slave05.mesos-31227,0,0,0,1,,2,1639,680,,0,,0,0,0,0,UP,1,1,0,0,0,49,0,,1,4,1,,2,,2,0,,1,L7OK,200,1,0,2,0,0,0,0,0,,,,0,0,,,,,12,OK,,0,1,1,1,
::groupB::appB-cluster,::groupB::appB-slave06.mesos-31103,0,0,0,1,,2,5122,3937,,0,,0,0,0,0,UP,1,1,0,0,0,49,0,,1,4,2,,2,,2,0,,1,L7OK,200,0,0,2,0,0,0,0,0,,,,0,0,,,,,10,OK,,0,0,1,1,
::groupB::appB-cluster,BACKEND,0,0,0,1,200,4,6761,4617,0,0,,0,0,0,0,UP,2,2,0,,0,49,0,,1,4,0,,4,,1,0,,1,,,,0,4,0,0,0,0,,,,,0,0,0,0,0,0,10,,,0,0,1,2,
::groupB::appC-cluster,::groupB::appC-slave07.mesos-31520,0,0,0,0,,0,0,0,,0,,0,0,0,0,no check,1,1,0,,,,,,1,5,1,,0,,2,0,,0,,,,0,0,0,0,0,0,0,,,,0,0,,,,,-1,,,0,0,0,0,
::groupB::appC-cluster,BACKEND,0,0,0,0,200,0,0,0,0,0,,0,0,0,0,UP,1,1,0,,0,49,0,,1,5,0,,0,,1,0,,0,,,,0,0,0,0,0,0,,,,,0,0,0,0,0,0,-1,,,0,0,0,0,
`)
	points, err := parser.Parse(content)
	fmt.Println(len(points))
	fmt.Println(err)
	// Output:
	// 8
	// <nil>
}

type TestingReader struct {
	in io.Reader
}

func newTestingReader(in string) io.ReadCloser {
	return TestingReader{strings.NewReader(in)}
}

func (t TestingReader) Read(p []byte) (n int, err error) {
	return t.in.Read(p)
}

func (t TestingReader) Close() error {
	return nil
}
