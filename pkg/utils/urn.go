package utils

import "strings"

func InfoHashFromMagnet(magnet string) string {
	urn := strings.Split(magnet, "&")[0]
	return urn[20 : len(urn)-1]
}
