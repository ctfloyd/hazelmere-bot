package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command"
)

type Bot struct {
	session *discordgo.Session
}

func NewBot(session *discordgo.Session) *Bot {
	return &Bot{session: session}
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

func (bot *Bot) Cleanup() {
	if bot.session != nil {
		_ = bot.session.Close()
	}
}
