package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
)

type Bot struct {
	logger  hz_logger.Logger
	session *discordgo.Session
}

func NewBot(logger hz_logger.Logger, session *discordgo.Session) *Bot {
	return &Bot{logger: logger, session: session}
}

func (bot *Bot) Run(commands []command.DiscordCommand) {
	err := bot.session.Open()
	if err != nil {
		panic(err)
	}

	for _, discordCommand := range commands {
		_, err = bot.session.ApplicationCommandCreate(bot.session.State.User.ID, discordCommand.GuildId, &discordCommand.Command)
		if err != nil {
			panic(err)
		}
	}

	bot.session.AddHandler(func(discord *discordgo.Session, i *discordgo.InteractionCreate) {
		for _, discordCommand := range commands {
			if i.ApplicationCommandData().Name == discordCommand.Name {
				discordCommand.Handler.HandleCommand(discord, i)
			}
		}
	})
}

func (bot *Bot) SendMessageEmbed(channelId string, message command.DiscordMessage) {
	if bot.session == nil {
		bot.logger.WarnArgs(context.Background(), "Tried to send a message but session is nil")
		return
	}

	embed, err := message.ToEmbed()
	if err != nil {
		bot.logger.ErrorArgs(context.Background(), "Tried to send a message but could not convert to embed: %v", err)
		return
	}

	_, err = bot.session.ChannelMessageSendEmbed(channelId, embed)
	if err != nil {
		bot.logger.ErrorArgs(context.Background(), "Tried to send a message but could not send embed: %v", err)
	}
}

func (bot *Bot) Cleanup() {
	if bot.session != nil {
		_ = bot.session.Close()
	}
}
