package initialize

import (
	"github.com/ctfloyd/hazelmere-bot/src/internal/job"
	"github.com/go-co-op/gocron/v2"
	"time"
)

func InitializeScheduler(
	userUpdateJob *job.UserUpdateJob,
) gocron.Scheduler {
	location, err := time.LoadLocation("America/Chicago")
	if err != nil {
		panic(err)
	}

	scheduler, err := gocron.NewScheduler(gocron.WithLocation(location))
	if err != nil {
		panic(err)
	}

	everyDayAtSixAm := gocron.DailyJob(1,
		gocron.NewAtTimes(
			gocron.NewAtTime(6, 0, 0),
		),
	)
	_, err = scheduler.NewJob(
		everyDayAtSixAm,
		gocron.NewTask(userUpdateJob.Run),
	)

	if err != nil {
		panic(err)
	}

	return scheduler
}
