package command

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

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

type DiscordMessage struct {
	Embeds []*discordgo.MessageEmbed
}

func (dm *DiscordMessage) ToWebhookEdit() discordgo.WebhookEdit {
	return discordgo.WebhookEdit{
		Embeds: &dm.Embeds,
	}
}

func (dm *DiscordMessage) ToEmbed() (*discordgo.MessageEmbed, error) {
	if len(dm.Embeds) == 0 {
		return nil, errors.New("no embeds found")
	}

	return dm.Embeds[0], nil
}
