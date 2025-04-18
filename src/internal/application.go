package internal

import (
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command/user_update"
	"github.com/ctfloyd/hazelmere-bot/src/internal/gain"
	"github.com/ctfloyd/hazelmere-bot/src/internal/initialize"
	"github.com/ctfloyd/hazelmere-bot/src/internal/job"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"github.com/go-co-op/gocron/v2"
)

type Application struct {
	discord   *discord.Bot
	commands  []command.DiscordCommand
	scheduler gocron.Scheduler
}

func (app *Application) Initialize(cfg *hz_config.Config) {
	logger := hz_logger.NewZeroLogAdapater(hz_logger.LogLevelFromString(cfg.ValueOrPanic("log.level")))

	hz := initialize.InitializeHazelmereClient(cfg, logger)
	hzResilient := initialize.InitializeHazelmereClientResilient(cfg, logger)
	gainedService := gain.NewGainedService(logger, hz)
	gainedServiceResilient := gain.NewGainedService(logger, hzResilient)
	app.discord = initialize.InitializeDiscord(cfg, logger)

	userUpdateJob := job.NewUserUpdateJob(logger, gainedServiceResilient, hzResilient, app.discord, cfg.ValueOrPanic("channelId"))
	app.scheduler = initialize.InitializeScheduler(userUpdateJob)

	userUpdateCommand := user_update.NewUserUpdateCommand(logger, gainedService)
	app.commands = []command.DiscordCommand{
		userUpdateCommand,
	}
}

func (app *Application) Run() {
	app.discord.Run(app.commands)
	app.scheduler.Start()
}

func (app *Application) Cleanup() {
	if app.discord != nil {
		app.discord.Cleanup()
	}
}
