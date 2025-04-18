package user_update

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ctfloyd/hazelmere-api/src/pkg/api"
	"github.com/ctfloyd/hazelmere-bot/src/internal/constant"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command"
	"github.com/ctfloyd/hazelmere-bot/src/internal/gain"
)

func CreateMessage(username string, time int, unit gain.TimeUnit, gains gain.UserGains) command.DiscordMessage {
	return command.DiscordMessage{
		Embeds: []*discordgo.MessageEmbed{
			{
				URL:         "https://api.hazelmere.xyz",
				Type:        discordgo.EmbedTypeRich,
				Title:       fmt.Sprintf("%s gains (%d %s)", username, time, unit),
				Description: buildDescription(username, gains),
				Color:       2719929,
			},
		},
	}
}

func buildDescription(username string, gains gain.UserGains) string {
	if len(gains.Skills) == 0 {
		return fmt.Sprintf("%s is stinky and has not made any gains in the time range!", username)
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
