package mdate

import (
	"strconv"
	"time"
)

type DateTime struct {
	time.Time
}

func Now() DateTime {
	return DateTime{Time: time.Now()}
}

//NewFromUnix 根据Unix时间戳创建
func NewFromUnix(sec int64, nsec int64) DateTime {
	return DateTime{Time: time.Unix(sec, nsec)}
}

//New 通过字符串创建时间，失败返回error
func New(value string, opts ...Option) (DateTime, error) {
	cfg := config{layout: LayoutDefault, loc: time.Local}
	for _, op := range opts {
		op(&cfg)
	}
	t, err := time.ParseInLocation(cfg.layout, value, cfg.loc)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime{Time: t}, nil
}

func NewDefault(value string, defVal time.Time, opts ...Option) DateTime {
	dt, err := New(value, opts...)
	if err != nil {
		return DateTime{Time: defVal}
	}
	return dt
}

//Millis 返回Unix毫秒时间戳
func (dt DateTime) Millis() int64 {
	return dt.UnixNano() / 1e6
}

//Start 返回当天开始时间戳
func (dt DateTime) Start() int64 {
	return StartOfDay(dt.Time)
}

//End 返回当天结束时间戳
func (dt DateTime) End() int64 {
	return EndOfDay(dt.Time)
}

//StartOfMonth 返回月份第一秒时间戳
func (dt DateTime) StartOfMonth() int64 {
	return StartOfMonth(dt.Time)
}

//EndOfMonth 返回月份最后一秒时间戳
func (dt DateTime) EndOfMonth() int64 {
	return EndOfMonth(dt.Time)
}

//Diff 返回日期相差天数
func (dt DateTime) Diff(d1 DateTime) int64 {
	return (dt.Start() - d1.Start()) / Day
}

//Digit 将时间转换int类型 到秒结束
func (dt DateTime) Digit() int {
	return dt.toInt(dt.Format(LayoutDigit))
}

//DayDigit 将时间转换为int，到天结束
func (dt DateTime) DayDigit() int {
	return dt.toInt(dt.Format(LayoutDayDigit))
}

func (dt DateTime) toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

//FormatDay 返回到天的格式化时间，格式：yyyy-MM-dd
func (dt DateTime) FormatDay() string {
	return dt.Format(LayoutDay)
}

//FormatMonth 返回到月的格式化时间，格式：yyyy-MM-dd
func (dt DateTime) FormatMonth() string {
	return dt.Format(LayoutMonth)
}

//ToString 返回日期字符串，格式：yyyy-MM-dd HH:mm:ss
func (dt DateTime) ToString() string {
	return dt.Format(LayoutDefault)
}

