package utils

import (
	"fmt"
	"strings"
)

// 将秒转换成年天时分秒
func ConvertSecondsToHumanReadable(seconds int64) string {
	const dayInSeconds = 24 * 60 * 60
	const hourInSeconds = 60 * 60
	const minuteInSeconds = 60

	days := int(seconds / dayInSeconds)
	seconds %= dayInSeconds

	hours := int(seconds / hourInSeconds)
	seconds %= hourInSeconds

	minutes := int(seconds / minuteInSeconds)
	seconds %= minuteInSeconds

	var components []string

	if days > 0 {
		components = append(components, fmt.Sprintf("%d 天", days))
	}
	if hours > 0 {
		components = append(components, fmt.Sprintf("%d 时", hours))
	}
	if minutes > 0 {
		components = append(components, fmt.Sprintf("%d 分", minutes))
	}
	if seconds > 0 {
		components = append(components, fmt.Sprintf("%d 秒", seconds))
	}

	return strings.Join(components, " ")
}
