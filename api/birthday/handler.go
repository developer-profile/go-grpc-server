package api

import (
	context "context"
	"strconv"
	"time"
)

type server struct {
}

func (s *server) CheckBirthday(ctx context.Context, in *Date) (*BirthdayStatus, error) {
	// Create date 
	userBirthday := strconv.FormatInt(in.Day, 10) + "-" + istrconv.FormatInt(in.Month, 10) + strconv.FormatInt(in.Year, 10)
	currentDate := time.Now()

	if userBirthday == currentDate.Format("01-02-2006") {
		return &BirthdayStatus(Status: true, Age: 29)
	}
}
