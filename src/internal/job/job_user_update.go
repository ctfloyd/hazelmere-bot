package job

import (
	"context"
	"github.com/ctfloyd/hazelmere-api/src/pkg/api"
	"github.com/ctfloyd/hazelmere-api/src/pkg/client"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command/user_update"
	"github.com/ctfloyd/hazelmere-bot/src/internal/gain"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
)

type UserUpdateJob struct {
	logger        hz_logger.Logger
	gainedService gain.GainedService
	hazelmere     *client.Hazelmere
	bot           *discord.Bot
	channelId     string
}

func NewUserUpdateJob(
	logger hz_logger.Logger,
	gainedService gain.GainedService,
	hazelmere *client.Hazelmere,
	bot *discord.Bot,
	channelId string,
) *UserUpdateJob {
	return &UserUpdateJob{
		logger:        logger,
		gainedService: gainedService,
		hazelmere:     hazelmere,
		bot:           bot,
		channelId:     channelId,
	}
}

func (job *UserUpdateJob) Run() {
	ctx := context.Background()

	userResponse, err := job.hazelmere.User.GetAllUsers()
	if err != nil {
		job.logger.ErrorArgs(ctx, "Error getting all users: %v", err)
		return
	}

	for _, user := range userResponse.Users {
		if user.TrackingStatus == api.TrackingStatusEnabled {
			job.logger.InfoArgs(ctx, "Job is calculating gains and posting update for user (%s) on channel (%s).", user.RunescapeName, job.channelId)
			gains, err := job.gainedService.CalculateUserGains(user.RunescapeName, 1, gain.TimeUnitDays)
			if err != nil {
				job.logger.ErrorArgs(ctx, "Error calculating gains for user (%s): %v", user.RunescapeName, err)
				continue
			}

			message := user_update.CreateMessage(user.RunescapeName, 1, gain.TimeUnitDays, gains)
			job.bot.SendMessageEmbed(job.channelId, message)
		}
	}

}
