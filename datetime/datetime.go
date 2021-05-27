package datetime

import (
	"strconv"
	"time"
)

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
	opt := option{layout: LayoutDefault, loc: time.Local}
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

//Unix 返回Unix时间戳
func (d DateTime) Unix() int64 {
	return d.t.Unix()
}

//Millis 返回Unix毫秒时间戳
func (d DateTime) Millis() int64 {
	return d.t.UnixNano() / 1e6
}

//Start 返回当天开始时间戳
func (d DateTime) Start() int64 {
	return StartOfDay(d.t)
}

//End 返回当天结束时间戳
func (d DateTime) End() int64 {
	return EndOfDay(d.t)
}

//Time 返回time.Time
func (d DateTime) Time() time.Time {
	return d.t
}

//Digit 将时间转换int类型 到秒结束
func (d DateTime) Digit() int64 {
	return d.toInt(d.t.Format(LayoutDigit))
}

//DayDigit 将时间转换为int，到天结束
func (d DateTime) DayDigit() int64 {
	return d.toInt(d.t.Format(LayoutDayDigit))
}

func (d DateTime) toInt(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

//Day 返回到天的格式化时间
func (d DateTime) FormatDay() string {
	return d.t.Format(LayoutDay)
}

//Month 返回到月的格式化时间
func (d DateTime) FormatMonth() string {
	return d.t.Format(LayoutMonth)
}

func (d DateTime) ToString() string {
	return d.t.Format(LayoutDefault)
}

func (d DateTime) Format(layout string) string {
	return d.t.Format(layout)
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
