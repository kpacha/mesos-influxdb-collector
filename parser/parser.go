package parser

import (
	"github.com/kpacha/mesos-influxdb-collector/store"
	"io"
)

type Parser interface {
	Parse(r io.Reader) ([]store.Point, error)
}
