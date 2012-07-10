package s3dm

import "strconv"

// Fixed-point number in Q32.32 format.
type Fixed64 int64

func NewFixed64(f float64) Fixed64 {
	return Fixed64(f*(1<<32))
}

func (f Fixed64) Modf() (int Fixed64, frac Fixed64) {
	int = f &^ 0xffffffff
	frac = f & 0xffffffff
	return
}

func (f Fixed64) Int64() int64 {
	return int64(f >> 32)
}

func (f Fixed64) Float64() float64 {
	return float64(f) / (1 << 32)
}

func (f Fixed64) String() string {
	_, frac := f.Modf()
	return strconv.FormatInt(f.Int64(), 10) + strconv.FormatFloat(frac.Float64(), 'f', -1, 64)[1:]
}
