package example

//go:generate gen
//go:generate jsonenums -type=Bits

// +gen stringer flags
type Bits uint32

const (
	Bit01 Bits = 1 << iota
	Bit02
	Bit03
	Bit04
	Bit05
	Bit06
	Bit07
	Bit08
	Bit09
	Bit10
	Bit11
	Bit12
	Bit13
	Bit14
	Bit15
	Bit16
	Bit17
	Bit18
	Bit19
	Bit20
	Bit21
	Bit22
	Bit23
	Bit24
	Bit25
	Bit26
	Bit27
	Bit28
	Bit29
	Bit30
	Bit31
	Bit32
)

const (
	AnyEven Bits = Bit02 | Bit04 | Bit06 | Bit08 | Bit10 | Bit12 | Bit14 | Bit16 | Bit18 | Bit20 | Bit22 | Bit24 | Bit26 | Bit28 | Bit30 | Bit32
	AnyOdd  Bits = Bit01 | Bit03 | Bit05 | Bit07 | Bit09 | Bit11 | Bit13 | Bit15 | Bit17 | Bit19 | Bit21 | Bit23 | Bit25 | Bit27 | Bit29 | Bit31
)

const MutuallyExclusive Bits = Bit01 | Bit02 | Bit03
