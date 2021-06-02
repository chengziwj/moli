package mdate

import "time"

type Option func(t *config)

type config struct {
	layout string
	loc    *time.Location
}

func WithLoc(loc *time.Location) Option {
	return func(t *config) {
		t.loc = loc
	}
}

func WithLayout(layout string) Option {
	return func(t *config) {
		t.layout = layout
	}
}