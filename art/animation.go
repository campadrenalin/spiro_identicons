package art

type Animation struct {
	Value int // 3 bits (0-7 looping)
	Delta int // 3 bits (0-3 with sign bit, saturating)
	Accel int // 2 bits (0-1 with sign bit, does not change)

	Min, Max float64 // Range of output values
}

func NewAnimation(source int, min, max float64) (anim Animation) {
	// Value is simple, because no sign bit
	anim.Value = int(source & 0x7)
	source >>= 3

	// Delta is 2 bits plus sign
	anim.Delta = int(source & 0x3)
	if (source>>2)&1 == 1 {
		anim.Delta = -anim.Delta
	}
	source >>= 3

	// Accel is 1 bit plus sign bit
	anim.Accel = int(source & 0x1)
	if (source>>1)&1 == 1 {
		anim.Accel = -anim.Accel
	} // No shift after, we're done

	anim.Min = min
	anim.Max = max

	return anim
}
func (anim Animation) Float64() float64 {
	return (anim.Max-anim.Min)*float64(anim.Value)/8 + anim.Min
}
func (anim *Animation) Step() {
	// Add delta to value, looping
	anim.Value = (anim.Value + anim.Delta + 32) % 8 // The +32 is an easy way to loop when result < 0

	// Add accel to delta, and clamp delta to range
	anim.Delta += anim.Accel
	if anim.Delta < -3 {
		anim.Delta = -3
	} else if anim.Delta > 3 {
		anim.Delta = 3
	}
}
