package datetime

import "time"

type Option func(t *option)

type option struct {
	layout string
	loc    *time.Location
}

func WithLoc(loc *time.Location) Option {
	return func(t *option) {
		t.loc = loc
	}
}

func WithLayout(layout string) Option {
	return func(t *option) {
		t.layout = layout
	}
}

type DateTime struct {
	t time.Time
}

func New() DateTime {
	return DateTime{t: time.Now()}
}

func NewMust(value string, opts ...Option) (DateTime, error) {
	opt := option{layout: DefaultLayout, loc: time.Local}
	for _, op := range opts {
		op(&opt)
	}
	t, err := time.ParseInLocation(opt.layout, value, opt.loc)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime{t: t}, nil
}

func NewWithOption(value string, opts ...Option) DateTime {
	dt, err := NewMust(value, opts...)
	if err != nil {
		return New()
	}
	return dt
}

func (d DateTime) Unix() int64 {
	return d.t.Unix()
}

func (d DateTime) Start() int64 {
	return StartOfDay(d.t)
}

func (d DateTime) End() int64 {
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
