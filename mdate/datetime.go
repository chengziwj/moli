package mdate

import (
	"strconv"
	"time"
)

type Option func(t *config)

type config struct {
	layout  string
	defTime *time.Time
	loc     *time.Location
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

type DateTime struct {
	t time.Time
}

func Now() DateTime {
	return DateTime{t: time.Now()}
}

//NewFromUnix 根据Unix时间戳创建
func NewFromUnix(sec int64, nsec int64) DateTime {
	return DateTime{t: time.Unix(sec, nsec)}
}

//New 通过字符串创建时间，创建失败则防护当前时间
func New(value string, opts ...Option) DateTime {
	dt, err := NewMust(value, opts...)
	if err != nil {
		return Now()
	}
	return dt
}

func NewWithDefault(value string, defVal time.Time, opts ...Option) DateTime {
	dt, err := NewMust(value, opts...)
	if err != nil {
		return DateTime{t: defVal}
	}
	return dt
}

//NewMust 通过字符串创建时间
func NewMust(value string, opts ...Option) (DateTime, error) {
	cfg := config{layout: LayoutDefault, loc: time.Local}
	for _, op := range opts {
		op(&cfg)
	}
	t, err := time.ParseInLocation(cfg.layout, value, cfg.loc)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime{t: t}, nil
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

//FirstSecOfMonth 返回月份第一秒时间戳
func (d DateTime) FirstSecOfMonth() int64 {
	year, month, _ := d.t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, d.t.Location()).Unix()
}

//LastSecOfMonth 返回月份最后一秒时间戳
func (d DateTime) LastSecOfMonth() int64 {
	year, month, _ := d.t.Date()
	return time.Date(year, month+1, 1, 0, 0, 0, 0, d.t.Location()).Unix() - 1
}

//Time 返回time.Time
func (d DateTime) Time() time.Time {
	return d.t
}

//Digit 将时间转换int类型 到秒结束
func (d DateTime) Digit() int {
	return d.toInt(d.t.Format(LayoutDigit))
}

//DayDigit 将时间转换为int，到天结束
func (d DateTime) DayDigit() int {
	return d.toInt(d.t.Format(LayoutDayDigit))
}

func (d DateTime) toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

//FormatDay 返回到天的格式化时间，格式：yyyy-MM-dd
func (d DateTime) FormatDay() string {
	return d.t.Format(LayoutDay)
}

//FormatMonth 返回到月的格式化时间，格式：yyyy-MM-dd
func (d DateTime) FormatMonth() string {
	return d.t.Format(LayoutMonth)
}

//ToString 返回日期字符串，格式：yyyy-MM-dd HH:mm:ss
func (d DateTime) ToString() string {
	return d.t.Format(LayoutDefault)
}

//Format 格式化日期
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
