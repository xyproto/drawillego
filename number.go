package drawille

// Can represent an int or a float64, default is a float64
type FloatOrInt struct {
	floatType bool
	fval      float64
	ival      int
}

func NewFloat(val float64) *FloatOrInt {
	return &FloatOrInt{true, val, 0}
}

func NewInt(val int) *FloatOrInt {
	return &FloatOrInt{false, 0.0, val}
}

func (fi *FloatOrInt) New() *FloatOrInt {
	return &FloatOrInt{fi.floatType, fi.fval, fi.ival}
}

// Switch a coordinate type from float64 to int, or keep it as an int
func (fi *FloatOrInt) Normalize() {
	if fi.floatType {
		fi.floatType = false
		// Round float64 to int
		fi.ival = Round(fi.fval)
	}
}

func (fi *FloatOrInt) Normalized() *FloatOrInt {
	n := fi.New()
	n.Normalize()
	return n
}

func (fi *FloatOrInt) Int() int {
	if fi.floatType {
		// Round float64 to int
		return Round(fi.fval)
	}
	return fi.ival
}

func (fi *FloatOrInt) Float() float64 {
	if fi.floatType {
		return fi.fval
	}
	return float64(fi.ival)
}

func (fi *FloatOrInt) Bool() bool {
	return fi.Int() != 0
}
