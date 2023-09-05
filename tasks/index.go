package tasks

import (
	"github.com/go-co-op/gocron"
	"time"
)

func Init() {
	s := gocron.NewScheduler(time.UTC)

	_, err := s.Every(1).Day().Do(happyBirthday)

	if err != nil {
		panic("Job happy birthday not run correctly")
	}

	_, err = s.Every(10).Minutes().Do(luckyUser)

	if err != nil {
		panic("Job lucky user not run correctly")
	}

	_, err = s.Every(1).Minute().Do(sayHello)

	if err != nil {
		panic("Job say hello not run correctly")
	}

	s.StartAsync()
}
