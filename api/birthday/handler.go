package birthday

import (
	"context"
	"fmt"
	"time"
)

type Server struct{}

var t = time.Now()

func (s *Server) CheckBirthday(ctx context.Context, in *Date) (*BirthdayStatus, error) {

	// return &BirthdayStatus{Status: true, Age: 2}, nil

	var year = in.Year
	var month = in.Month
	var day = in.Month

	userBirthday := fmt.Sprint(year) + "-" + fmt.Sprint(month) + "-" + fmt.Sprint(day)
	today := getCurrentDate()

	//Check if user has birthday
	if userBirthday == today {
		return &BirthdayStatus{Status: true, Age: getAge(year, month)}, nil
	}
	return &BirthdayStatus{Status: false, Age: getAge(year, month)}, nil
}

func getCurrentDate() string {
	//Get current date
	today := fmt.Sprint(getYear()) + "-" + fmt.Sprint(getMonth()) + "-" + fmt.Sprint(getDay())
	return today
}

func getAge(year int32, month int32) int32 {
	var age int32

	//Calculate birthday
	age = getYear() - year
	if age == 0 {
		age = getMonth() - month
	}

	return age
}

func getYear() int32 { return int32(t.Year()) }

func getMonth() int32 { return int32(t.Month()) }

func getDay() int32 {
	return int32(t.Day())

}
