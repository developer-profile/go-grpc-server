package birthday

import (
	context "context"
	// "strconv"
	// "time"
)

type Server struct{}

func (s *Server) CheckBirthday(ctx context.Context, in *Date) (*BirthdayStatus, error) {

	return &BirthdayStatus{Status: true, Age: 2}, nil

	// t := time.Now()
	// year := in.Year
	// month := in.Month
	// day := in.Month

	// userBirthday := strconv.FormatInt(year, 10) + "-" + strconv.FormatInt(month, 10) + "-" + strconv.FormatInt(day, 10)

	// today := getCurrentDate(t)

	// Check if user has birthday
	// if userBirthday == today {
	// return &BirthdayStatus{Status: true, Age: getAge(year, month, t)}, nil
	// }
	// return &BirthdayStatus{Status: false, Age: getAge(year, month, t)}, nil
	// }

	// func getCurrentDate(Time t) string {
	// Get current date
	// today := getYear(t) + "-" + getMonth(t) + "-" + getDay(t)
	// return today
	// }

	// func getAge(int64 year, int64 month, Time t) int64 {
	// Calculate birthday
	// age := getYear(t) - year
	// if age == 0 {
	// age := getMonth(t) - month
	// }

	// return age
	// }

	// func getYear(Time t) string { return strconv.Itoa(int(t.Month())) }

	// func getMonth(Time t) string { return strconv.Itoa(int(t.Month())) }

	// func getDay(Time t) string { return strconv.Itoa(t.Day())

}
