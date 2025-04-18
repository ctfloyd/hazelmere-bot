package command

import "github.com/bwmarrin/discordgo"

const GlobalCommandGuildId = ""

type DiscordCommandHandler interface {
	HandleCommand(*discordgo.Session, *discordgo.InteractionCreate)
}

type DiscordCommand struct {
	Name    string
	GuildId string
	Command discordgo.ApplicationCommand
	Handler DiscordCommandHandler
}
