package format

import "fmt"

func Size[T int | uint | int32 | uint32 | int64 | uint64](size T) string {
	s := float64(size)
	prefixes := []string{"B", "kB", "MB", "GB", "TB", "PB"}
	for i, prefix := range prefixes {
		if s < 1000 || i == len(prefixes)-1 {
			return fmt.Sprintf("%.3g %s", s, prefix)
		}
		s /= 1024
	}
	return ""
}
