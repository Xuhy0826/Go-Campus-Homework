package timeutil

import "time"

const dateTimeLayout = "2006-01-02 15:04:05"
const dateLayout = "2006-01-02"

func ToDateTimeString(time time.Time) string {
	return time.Format(dateTimeLayout)
}

func ToDateString(time time.Time) string {
	return time.Format(dateLayout)
}

//ToDateTime time template is "yyyy-MM-dd HH:mm:ss"
func ToDateTime(timeStr string) (time.Time, error) {
	return time.ParseInLocation(dateTimeLayout, timeStr, time.Local)
}

func GetAge(birthDay time.Time) int32 {
	age := time.Now().Year() - birthDay.Year()

	if time.Now().Month() < birthDay.Month() {
		age--
	}
	return age
}
