package parser

import (
	"github.com/kpacha/mesos-influxdb-collector/store"
	"io"
)

type Parser interface {
	Parse(r io.ReadCloser) ([]store.Point, error)
}
