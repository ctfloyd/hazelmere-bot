package user_update

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ctfloyd/hazelmere-api/src/pkg/api"
	"github.com/ctfloyd/hazelmere-bot/src/internal/constant"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command"
	"github.com/ctfloyd/hazelmere-bot/src/internal/gain"
)

const BossesPerMessage = 25

func CreateMessage(username string, time int, unit gain.TimeUnit, gains gain.UserGains) command.DiscordMessage {
	descriptions := buildDescriptions(username, gains)
	message := command.DiscordMessage{
		Embeds: []*discordgo.MessageEmbed{},
	}
	for _, description := range descriptions {
		message.Embeds = append(message.Embeds, &discordgo.MessageEmbed{
			URL:         "https://api.hazelmere.xyz",
			Type:        discordgo.EmbedTypeRich,
			Title:       fmt.Sprintf("%s gains (%d %s)", username, time, unit),
			Description: description,
			Color:       2719929,
		})
	}
	return message
}

func buildDescriptions(username string, gains gain.UserGains) []string {
	descriptions := []string{
		buildSkillDescription(username, gains),
	}

	bossDescriptions := buildBossDescriptions(gains)
	for _, description := range bossDescriptions {
		descriptions = append(descriptions, description)
	}

	return descriptions
}

func buildBossDescriptions(gains gain.UserGains) []string {
	descriptions := []string{}

	description := ""

	count := 0
	for _, boss := range api.AllBossActivityTypes {
		if count%BossesPerMessage == 0 && count != 0 {
			descriptions = append(descriptions, description)
			description = ""
		}

		emoji, ok := constant.Emojis[boss]
		if !ok {
			continue
		}

		gain, ok := gains.Bosses[boss]
		if ok {
			description += fmt.Sprintf("%s %s\n", emoji, gain.Amount)
			count += 1
		}
	}

	if len(description) > 0 {
		descriptions = append(descriptions, description)
	}

	return descriptions
}

func buildSkillDescription(username string, gains gain.UserGains) string {
	if len(gains.Skills) == 0 {
		return fmt.Sprintf("%s is stinky and has not made any skill gains in the time range!", username)
	}

	description := fmt.Sprintf("%s **%s**\n",
		constant.Emojis[api.ActivityTypeOverall],
		gains.Skills[api.ActivityTypeOverall].Amount,
	)

	for _, skill := range api.AllSkillActivityTypes {
		if skill == api.ActivityTypeOverall {
			continue
		}

		emoji, ok := constant.Emojis[skill]
		if !ok {
			emoji = constant.Emojis[api.ActivityTypeUnknown]
		}

		gain, ok := gains.Skills[skill]
		if ok {
			description += fmt.Sprintf("%s %s\n", emoji, gain.Amount)
		}
	}
	return description
}
