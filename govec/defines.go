package govec

type UniformInt		int
type UniformInt32	int32
type UniformInt64	int64
type UniformFloat32	float32
type UniformFloat64	float64
var ProgramIndex	UniformInt
var ProgramCount	UniformInt

func Range(a ...interface{}) []struct{} {
	panic("Range: Dummy function. Can't run this")
}

func DoubleRange(a ...interface{}) bool {
	panic("DoubleRange: Dummy function. Can't run this")
}

func ReduceAddInt(a int) int {
	panic("ReduceAdd: Dummy function. Can't run this")
}

func ReduceAddFloat32(a float32) float32 {
	panic("ReduceAdd: Dummy function. Can't run this")
}

func ReduceAddFloat64(a float64) float64 {
	panic("ReduceAdd: Dummy function. Can't run this")
}
