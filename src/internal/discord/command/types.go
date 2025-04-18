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

func (dm *DiscordMessage) ToWebhookEdit() (discordgo.WebhookEdit, error) {
	if len(dm.Embeds) == 0 {
		return discordgo.WebhookEdit{}, errors.New("no embeds found")
	}

	return discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{dm.Embeds[0]},
	}, nil
}

func (dm *DiscordMessage) ToFollowupMessageCreates() []discordgo.WebhookParams {
	if len(dm.Embeds) <= 1 {
		return []discordgo.WebhookParams{}
	}

	var webhookParams []discordgo.WebhookParams
	for i := 1; i < len(dm.Embeds); i++ {
		webhookParams = append(webhookParams, discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{dm.Embeds[i]},
		})
	}
	return webhookParams
}

func (dm *DiscordMessage) ToEmbed() (*discordgo.MessageEmbed, error) {
	if len(dm.Embeds) == 0 {
		return nil, errors.New("no embeds found")
	}

	return dm.Embeds[0], nil
}
