package mdate

import "time"

type Option func(t *config)

type config struct {
	layout  string
	loc     *time.Location
	defTime time.Time
	hasDef  bool
}

func WithLoc(loc *time.Location) Option {
	return func(c *config) {
		c.loc = loc
	}
}

func WithLayout(layout string) Option {
	return func(c *config) {
		c.layout = layout
	}
}

func WithDefault(dt time.Time) Option {
	return func(c *config) {
		c.defTime = dt
		c.hasDef = true
	}
}
