package datetime

import "time"

type Option func(t *Options)

type Options struct {
	layout string
	loc    *time.Location
}

func WithLoc(loc *time.Location) Option {
	return func(t *Options) {
		t.loc = loc
	}
}

type DateTime struct {
	t      time.Time
}

func New() DateTime {
	return DateTime{t: time.Now()}
}

func NewWithOption(value string, opts ...Option) DateTime {
	options := Options{layout: DefaultLayout,loc: time.Local}
	for _, op := range opts {
		op(&options)
	}
	t, err := time.ParseInLocation(options.layout, value, options.loc)
	if err != nil {
		return New()
	}
	return DateTime{t: t}
}

func (d *DateTime) Unix() int64 {
	return d.t.Unix()
}

func (d *DateTime) Start() int64 {
	return StartOfDay(d.t)
}

func (d *DateTime) End() int64 {
	return EndOfDay(d.t)
}

//StartOfDay 获取当天开始时间戳
func StartOfDay(dt time.Time) int64 {
	year, month, day := dt.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, dt.Location()).Unix()
}

//EndOfDay 获取当天结束时间戳
func EndOfDay(dt time.Time) int64 {
	year, month, day := dt.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, dt.Location()).Unix() - 1
}
