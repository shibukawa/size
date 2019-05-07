// Package size includes contents for representing
// sizes.
//
// Usage:
//   var storageSize = 10 * MB
package size

// Size is type for storage size
type Size int64

func fmtInt(v int64, o *[6]byte, offset int) int {
	for v > 0 {
		digit := v % 10
		o[offset] = byte(digit) + '0'
		offset--
		v /= 10
	}
	return offset
}

func (s Size) String() string {
	var buf [6]byte
	buf[5] = 'B'
	offset := 4
	u := int64(s)
	neg := u < 0
	if neg {
		u = -u
	}
	switch {
	case u == 0:
		buf[offset] = '0'
		offset--
	case u < int64(KB):
		offset = fmtInt(u, &buf, offset)
	case u < int64(MB):
		buf[offset] = 'K'
		offset = fmtInt(u/int64(KB), &buf, offset-1)
	case u < int64(GB):
		buf[offset] = 'M'
		offset = fmtInt(u/int64(MB), &buf, offset-1)
	case u < int64(TB):
		buf[offset] = 'G'
		offset = fmtInt(u/int64(GB), &buf, offset-1)
	case u < int64(PB):
		buf[offset] = 'T'
		offset = fmtInt(u/int64(TB), &buf, offset-1)
	case u < int64(EB):
		buf[offset] = 'P'
		offset = fmtInt(u/int64(PB), &buf, offset-1)
	default:
		buf[offset] = 'E'
		offset = fmtInt(u/int64(EB), &buf, offset-1)
	}
	if neg {
		buf[offset] = '-'
		offset--
	}
	return string(buf[offset+1:])
}

const (
	// B means Byte
	B Size = 1
	// KB means kilobyte
	KB = 1000 * B
	// MB means megabyte
	MB = 1000 * KB
	// GB means gigabyte
	GB = 1000 * MB
	// TB means terabyte
	TB = 1000 * GB
	// PB means petabyte
	PB = 1000 * TB
	// EB means exabyte
	EB = 1000 * PB
)
