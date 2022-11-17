package convert

import "strconv"

const base = 10

func ParseUint(s string) (uint64, error) {
	return strconv.ParseUint(s, base, 64)
}

func ParseUint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, base, 32)
	if err != nil {
		return 0, err
	}
	return uint32(i), nil
}

func ParseInt(s string) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

func ParseInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, base, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func FormatInt(i int64) string {
	return strconv.FormatInt(i, base)
}

func FormatInt32(i int32) string {
	return FormatInt(int64(i))
}

func FormatUint(i uint64) string {
	return strconv.FormatUint(i, base)
}

func FormatUint32(i uint32) string {
	return FormatUint(uint64(i))
}
