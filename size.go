package format

import "fmt"

func Size(size int64) string {
	x := float64(size)
	prefixes := []string{"B", "kB", "MB", "GB", "TB", "PB"}
	for i, prefix := range prefixes {
		if x < 1000 || i == len(prefixes)-1 {
			return fmt.Sprintf("%.3g %s", x, prefix)
		}
		x /= 1024
	}
	return ""
}
