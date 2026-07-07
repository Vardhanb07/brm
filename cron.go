package brm

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrParseDuration = errors.New("unable to parse duration, use formats: `x`min, `x`h, `x`d, `x`m and `x` should be a integer")
)

// Example of job definition:
// .---------------- minute (0 - 59)
// |  .------------- hour (0 - 23)
// |  |  .---------- day of month (1 - 31)
// |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
//
//	|  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
//	|  |  |  |  |
//
// *  *  *  *  * user-name  command to be executedd
// duration should be `x`min, `x`h, `x`d, `x`m
// `x`min - current time + `x` minutes after file will be deleted
// `x`h - current time + `x` hours after file will be deleted
// `x`d - current time + `x` days after file will be deleted
// `x`m - current time + `x` months after file will be deleted
// composition is allowed with no space i.e, `x`h`y`min. Order is irrelevant
func Cron(duration string) error {
	return nil
}

func convertToIntAndReset(t *string) (int, error) {
	n, err := strconv.ParseInt(*t, 10, 64)
	if err != nil {
		return 0, ErrParseDuration
	}
	*t = ""
	return int(n), nil
}

func parseDuration(duration string) (mins, hours, days, months int, err error) {
	t := ""
	i := 0
	for i < len(duration) {
		if i+3 < len(duration) && strings.Compare(duration[i:i+3], "min") == 0 {
			n, _err := convertToIntAndReset(&t)
			mins += n
			err = _err
			i += 2
		} else if string(duration[i]) == "h" {
			n, _err := convertToIntAndReset(&t)
			hours += n
			err = _err
		} else if string(duration[i]) == "d" {
			n, _err := convertToIntAndReset(&t)
			days += n
			err = _err
		} else if string(duration[i]) == "m" {
			n, _err := convertToIntAndReset(&t)
			months += n
			err = _err
		} else {
			t += string(duration[i])
			_, _err := strconv.ParseInt(string(duration[i]), 10, 64)
			if _err != nil {
				err = ErrParseDuration
			}
			if err != nil {
				return
			}
		}
		i += 1
	}
	return
}
