package nyaa

import (
	"strconv"
	"strings"
)

func convertSize(size string) int64 {
	if strings.HasSuffix(size, "MiB") {
		mib := strings.TrimSuffix(size, " MiB")
		size, err := strconv.ParseFloat(mib, 64)
		if err != nil {
			return 0
		}

		return int64(size*10) * 10_000
	}

	if strings.HasSuffix(size, "GiB") {
		mib := strings.TrimSuffix(size, " GiB")
		size, err := strconv.ParseFloat(mib, 64)
		if err != nil {
			return 0
		}

		return int64(size*10) * 10_000_000
	}

	return 0
}

func isAnime(item *Item) bool {
	return strings.HasPrefix(item.CategoryId, "1_")
}
