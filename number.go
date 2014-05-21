package drawille

// Can represent an int or a float64, default is a float64
type Number struct {
	floatType bool
	fval      float64
	ival      int
}

func NewFloat(val float64) *Number {
	return &Number{true, val, 0}
}

func NewInt(val int) *Number {
	return &Number{false, 0.0, val}
}

func (fi *Number) New() *Number {
	return &Number{fi.floatType, fi.fval, fi.ival}
}

// Switch a coordinate type from float64 to int, or keep it as an int
func (fi *Number) Normalize() {
	if fi.floatType {
		fi.floatType = false
		// Round float64 to int
		fi.ival = int(fi.fval + 0.5)
	}
}

func (fi *Number) Normalized() *Number {
	n := fi.New()
	n.Normalize()
	return n
}

func (fi *Number) Int() int {
	if fi.floatType {
		// Round float64 to int
		return int(fi.fval + 0.5)
	}
	return fi.ival
}

func (fi *Number) Float() float64 {
	if fi.floatType {
		return fi.fval
	}
	return float64(fi.ival)
}

func (fi *Number) Bool() bool {
	return fi.Int() != 0
}
