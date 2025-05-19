package time

import (
	"fmt"
	"testing"
	"time"
)

func TestUnixTime(t *testing.T) {

	timeString := time.Unix(int64(1729493640), 0).Format("2006-01-02 15:04:05")
	timeStr := time.Unix(int64(1729493640), 0).Format("2006-01-02 15:04:05 (UTC-07:00)")

	fmt.Printf("time = %s\n", timeString)
	fmt.Printf("time = %s\n", timeStr)
}

func TestTimeLo(tt *testing.T) {
	// 时间戳
	timestamp := int64(1729493640)
	// 将时间戳转换为时间
	t := time.Unix(timestamp, 0)
	// 获取时区信息
	_, offset := t.Zone()
	// 格式化输出
	formattedTime := t.Format("2006-01-02 15:04:05")
	timeZone := fmt.Sprintf("(UTC%+d:00)", offset/3600)
	// 输出
	fmt.Println(formattedTime, timeZone)

}

func TestTimeLocation(t *testing.T) {
	occurTime := time.Unix(int64(1729493640), 0)
	_, offset := occurTime.Zone()
	sprintf := fmt.Sprintf("%s (UTC%+d:00)", occurTime.Format("2006-01-02 15:04:05"), offset/3600)

	fmt.Println(sprintf)

}
// 解决不了非整点的时区
func TestTimeZone(tt *testing.T) {
	// 时间戳
	timestamp := int64(1729493640)
	t := time.Unix(timestamp, 0)
	// 定义不同的时区
	timezones := []string{
		"Asia/Kolkata",    // EST/EDT
		"Europe/London",       // GMT/BST
		"Asia/Tokyo",          // JST
		"Africa/Johannesburg", // SAST
	}
	for _, tz := range timezones {
		location, err := time.LoadLocation(tz)
		if err != nil {
			fmt.Println("Error loading location:", err)
			continue
		}
		// 转换为指定时区
		localTime := t.In(location)
		// 格式化输出
		output := localTime.Format("2006-01-02 15:04:05 (UTC-07:00)")
		fmt.Println(output)

	}

}

func TestTimeForRFC(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now()
	format := now.In(loc).Format("2006-01-02 15:04:05 (UTC-07:00)")
	fmt.Println(format)
	format = time.Unix(now.Unix(), 0).In(loc).Format(time.RFC3339)
	fmt.Println(format)
}

func TestTimeZoneUTC(tt *testing.T) {
	// 时间戳
	timestamp := int64(1729493640)
	// 转换为UTC时间
	utcTime := time.Unix(timestamp, 0).UTC()
	// 加载指定时区
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	// 转换为指定时区的时间
	localTime := utcTime.In(location)
	// 格式化输出
	output := localTime.Format("2006-01-02 15:04:05 (UTC-07:00)")
	fmt.Println(output)

}
