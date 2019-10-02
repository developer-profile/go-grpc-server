package birthday

import (
	context "context"
	"strconv"
	"time"
)

type Server struct {
}

func (s *Server) CheckBirthday(ctx context.Context, in *Date) (*BirthdayStatus, error) {
	// Create date 
	userBirthday := strconv.FormatInt(in.Day, 10) + "-" + strconv.FormatInt(in.Month, 10) + strconv.FormatInt(in.Year, 10)
	currentDate := time.Now()

	if userBirthday == currentDate.Format("01-02-2006") {
		// TODO change age status
		return &BirthdayStatus{Status: true, Age: 29}, nil
	} else {
		// TODO change age status
		return &BirthdayStatus{Status: false, Age: 29}, nil
	}
}
