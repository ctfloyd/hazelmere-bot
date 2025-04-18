package internal

import (
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command/user_update"
	"github.com/ctfloyd/hazelmere-bot/src/internal/gain"
	"github.com/ctfloyd/hazelmere-bot/src/internal/initialize"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
)

type Application struct {
	discord  *discord.Bot
	commands []command.DiscordCommand
}

func (app *Application) Initialize(cfg *hz_config.Config) {
	logger := hz_logger.NewZeroLogAdapater(hz_logger.LogLevelFromString(cfg.ValueOrPanic("log.level")))

	app.discord = initialize.InitializeDiscord(cfg)

	hz := initialize.InitializeHazelmereClient(cfg, logger)
	gainedService := gain.NewGainedService(logger, hz)
	userUpdateCommand := user_update.NewUserUpdateCommand(logger, gainedService)
	app.commands = []command.DiscordCommand{
		userUpdateCommand,
	}
}

func (app *Application) Run() {
	app.discord.Run(app.commands)
}

func (app *Application) Cleanup() {
	if app.discord != nil {
		app.discord.Cleanup()
	}
}
