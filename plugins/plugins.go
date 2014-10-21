package plugins

import (
)

type MetricValue struct {
	Val string
}

type Collection struct {
	Items map[string]*MetricValue
}

type Plugin interface {
	OsType() string
	CollectData() (string, map[string]*MetricValue)
}
