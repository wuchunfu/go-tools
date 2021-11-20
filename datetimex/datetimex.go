package datetimex

import "time"

// GetCurrentFormatTime 获取当前格式化时间
func GetCurrentFormatTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurrentTime 获取当前时间
func GetCurrentTime() time.Time {
	return time.Now()
}

// FormatTimeStr 格式化时间字符串
func FormatTimeStr(timeStr string) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02T15:04:05.000Z", timeStr, loc)
	return theTime.Format("2006-01-02 15:04:05"), err
}

// StringToTime 时间字符转为时间
func StringToTime(date interface{}) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	ret, _ := time.ParseInLocation(timeLayout, date.(string), loc)
	return ret
}

// StringToTimestamp 时间字符转为时间戳
func StringToTimestamp(date interface{}) int64 {
	return StringToTime(date).Unix()
}

// TimeStampToTime 时间戳转为 time.Time
func TimeStampToTime(timeStamp int32) time.Time {
	return time.Unix(int64(timeStamp), 0)
}

// NowTime 当前时间，单位：秒
func NowTime() int64 {
	return time.Now().Unix()
}

// NowTimeToInt 当前时间，单位：秒
func NowTimeToInt() int {
	nowTime := NowTime()
	return int(nowTime)
}

// NowNanoTime 当前时间，单位：纳秒
func NowNanoTime() int64 {
	return time.Now().UnixNano()
}

// NowNanoTimeToInt 当前时间，单位：纳秒
func NowNanoTimeToInt() int {
	return int(NowNanoTime())
}

// BeforeTime 获取几天前时间，单位：秒
func BeforeTime(day int) int64 {
	return time.Now().AddDate(0, 0, day).Unix()
}

// BeforeTimeToInt 获取几天前时间，单位：秒
func BeforeTimeToInt(day int) int {
	beforeTime := BeforeTime(day)
	return int(beforeTime)
}
