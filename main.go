// Package storagesize package includes contents for represent
// storage sizes.
//
// Usage:
//   var storageSize = 10 * MB
package storagesize

// Size is type for storage size
type Size int64

const (
	// B means Byte
	B Size = 1
	// KB means kilobyte
	KB = 1024 * B
	// MB means megabyte
	MB = 1024 * KB
	// GB means gigabyte
	GB = 1024 * MB
	// TB means terabyte
	TB = 1024 * GB
	// PB means petabyte
	PB = 1024 * TB
	// EB means exabyte
	EB = 1024 * PB
)
