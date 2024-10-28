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

	// 获取当前时区
	localTimeZone, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println("Error loading local timezone:", err)
		return
	}
	// 时间戳
	timestamp := int64(1729501805)
	// 将时间戳转换为时间
	ti := time.Unix(timestamp, 0).In(localTimeZone).Format("2006-01-02 15:04:05")
	// 输出带有时区的时间信息
	fmt.Println("当前时区:", localTimeZone)
	fmt.Println("转换后的时间:", ti) // 转换后的时间: 2024-10-21 17:10:05 +0800 CST

}

func TestTimeZone(tt *testing.T) {
	// 时间戳
	timestamp := int64(1729493640)
	t := time.Unix(timestamp, 0)
	// 定义不同的时区
	timezones := []string{
		"America/New_York",    // EST/EDT
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
		// 获取时区信息
		_, offset := localTime.Zone()
		// 格式化输出
		formattedTime := localTime.Format("2006-01-02 15:04:05")
		timeZone := fmt.Sprintf("(UTC%+d:00)", offset/3600)
		// 输出
		fmt.Println(formattedTime, timeZone)
	}

}
