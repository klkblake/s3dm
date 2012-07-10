package s3dm

import "strconv"

const (
	fracBits = 32
	unit = 1 << fracBits
	lowMask = unit - 1
)

// Fixed-point number in Q32.32 format.
type Fixed64 int64

func NewFixed64(f float64) Fixed64 {
	return Fixed64(f * unit)
}

func (f Fixed64) Abs() Fixed64 {
	if f < 0 {
		return -f
	}
	return f
}

func (f Fixed64) Ceil() Fixed64 {
	if f <= 0 || f & lowMask == 0 {
		return f &^ lowMask
	}
	return f &^ lowMask + 1
}

func (f Fixed64) Floor() Fixed64 {
	if f >= 0 || f & lowMask == 0 {
		return f &^ lowMask
	}
	return f &^ lowMask - 1
}

func (f Fixed64) Modf() (int Fixed64, frac Fixed64) {
	int = f &^ lowMask
	frac = f & lowMask
	return
}

func (f Fixed64) Int64() int64 {
	return int64(f >> fracBits)
}

func (f Fixed64) Float64() float64 {
	return float64(f) / unit
}

func (f Fixed64) String() string {
	_, frac := f.Modf()
	return strconv.FormatInt(f.Int64(), 10) + strconv.FormatFloat(frac.Float64(), 'f', -1, 64)[1:]
}
