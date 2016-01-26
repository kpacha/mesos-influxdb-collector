package haproxy

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/kpacha/mesos-influxdb-collector/store"
)

type HAProxy struct {
	Node string
	Ts   time.Time
}

func NewHAProxy(node string) HAProxy {
	return HAProxy{
		Node: node,
		Ts:   time.Now(),
	}
}

func (hap HAProxy) Parse(r io.ReadCloser) ([]store.Point, error) {
	defer r.Close()
	stats := []store.Point{}
	csvReader := csv.NewReader(r)
	csvReader.Comment = '#'
	hap.Ts = time.Now()

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error parsing HAProxy stats endpoint response:", err.Error())
			return stats, err
		}
		point, err := hap.getHAProxyPoint(record)
		if err != nil {
			log.Println("Error parsing to HAProxyStats:", err.Error())
			return stats, err
		}
		stats = append(stats, point)
	}

	return stats, nil
}

func (hap HAProxy) getHAProxyPoint(record []string) (store.Point, error) {
	stats := make([]float64, len(record))
	for k, v := range record {
		switch k {
		case 0, 1, 17: // do not touch those because they're strings!
		default:
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				f = 0
			}
			stats[k] = f
		}
	}
	return store.Point{
		Measurement: "haproxy",
		Tags: map[string]string{
			"node": hap.Node,
			"app":  hap.sanitizeName(record[0]),
		},
		Fields: map[string]interface{}{
			"task":           hap.sanitizeName(record[1]),
			"qcur":           stats[2],
			"qmax":           stats[3],
			"scur":           stats[4],
			"smax":           stats[5],
			"slim":           stats[6],
			"stot":           stats[7],
			"bin":            stats[8],
			"bout":           stats[9],
			"dreq":           stats[10],
			"dresp":          stats[11],
			"ereq":           stats[12],
			"econ":           stats[13],
			"eresp":          stats[14],
			"wretr":          stats[15],
			"wredis":         stats[16],
			"status":         record[17],
			"weight":         stats[18],
			"act":            stats[19],
			"bck":            stats[20],
			"chkfail":        stats[21],
			"chkdown":        stats[22],
			"lastchg":        stats[23],
			"downtime":       stats[24],
			"qlimit":         stats[25],
			"pid":            stats[26],
			"iid":            stats[27],
			"sid":            stats[22],
			"throttle":       stats[24],
			"lbtot":          stats[25],
			"tracked":        stats[26],
			"type":           stats[27],
			"rate":           stats[28],
			"rate_lim":       stats[29],
			"rate_max":       stats[30],
			"check_status":   stats[31],
			"check_code":     stats[32],
			"check_duration": stats[33],
			"hrsp_1xx":       stats[34],
			"hrsp_2xx":       stats[35],
			"hrsp_3xx":       stats[36],
			"hrsp_4xx":       stats[37],
			"hrsp_5xx":       stats[38],
			"hrsp_other":     stats[39],
			"hanafail":       stats[40],
			"req_rate":       stats[41],
			"req_rate_max":   stats[42],
			"req_tot":        stats[43],
			"cli_abrt":       stats[44],
			"srv_abrt":       stats[45],
			"comp_in":        stats[46],
			"comp_out":       stats[47],
			"comp_byp":       stats[48],
			"comp_rsp":       stats[49],
			"lastsess":       stats[50],
			"last_chk":       stats[51],
			"last_agt":       stats[52],
			"qtime":          stats[53],
			"ctime":          stats[54],
			"rtime":          stats[55],
			"ttime":          stats[56],
		},
		Time: hap.Ts,
	}, nil
}

func (hap HAProxy) sanitizeName(name string) string {
	return strings.Replace(strings.TrimPrefix(name, "::"), "::", "_", -1)
}
