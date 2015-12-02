package parser

import (
	"io"

	"github.com/kpacha/mesos-influxdb-collector/store"
)

type Parser interface {
	Parse(r io.ReadCloser) ([]store.Point, error)
}

type ParserFrom interface {
	Parse(r io.ReadCloser, from string) ([]store.Point, error)
}
