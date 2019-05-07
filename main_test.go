package size

import (
	"strings"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestSize_String(t *testing.T) {
	testcases := []struct {
		value  Size
		result string
	}{
		{
			0 * B,
			"0B",
		},
		{
			10 * B,
			"10B",
		},
		{
			10000 * B,
			"10KB",
		},
		{
			10000000 * B,
			"10MB",
		},
		{
			10000000000 * B,
			"10GB",
		},
		{
			10000000000000 * B,
			"10TB",
		},
		{
			10000000000000000 * B,
			"10PB",
		},
		{
			9000000000000000000 * B,
			"9EB",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.result, func(t *testing.T) {
			if tt.value.String() != tt.result {
				t.Errorf("String() of %v should be %#v, but %#v", int64(tt.value), tt.result, tt.value.String())
			}
		})
	}
}

var suffix = map[Size]string{
	B:  "B",
	KB: "KB",
	MB: "MB",
	GB: "GB",
	TB: "TB",
	PB: "PB",
	EB: "EB",
}

func TestSize_String_PBT(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("positive", prop.ForAll(
		func(number int64, order Size) bool {
			size := Size(number) * order
			str := size.String()
			return strings.HasSuffix(str, suffix[order])
		},
		gen.Int64Range(1, 999),
		gen.OneConstOf(B, KB, MB, GB, TB, PB),
	))

	properties.Property("negative", prop.ForAll(
		func(number int64, order Size) bool {
			size := Size(number) * order
			str := size.String()
			return strings.HasSuffix(str, suffix[order]) && strings.HasPrefix(str, "-")
		},
		gen.Int64Range(-999, -1),
		gen.OneConstOf(B, KB, MB, GB, TB, PB),
	))

	properties.TestingRun(t)
}
