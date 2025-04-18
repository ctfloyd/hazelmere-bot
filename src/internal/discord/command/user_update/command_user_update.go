package user_update

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ctfloyd/hazelmere-api/src/pkg/api"
	"github.com/ctfloyd/hazelmere-bot/src/internal/constant"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord/command"
	"github.com/ctfloyd/hazelmere-bot/src/internal/gain"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
)

type UserUpdateHandler struct {
	logger        hz_logger.Logger
	gainedService gain.GainedService
}

func NewUserUpdateHandler(logger hz_logger.Logger, service gain.GainedService) *UserUpdateHandler {
	return &UserUpdateHandler{
		logger:        logger,
		gainedService: service,
	}
}

func NewUserUpdateCommand(logger hz_logger.Logger, gainedService gain.GainedService) command.DiscordCommand {
	one := 1.0
	name := "gained"

	return command.DiscordCommand{
		Name:    name,
		GuildId: command.GlobalCommandGuildId,
		Handler: NewUserUpdateHandler(logger, gainedService),
		Command: discordgo.ApplicationCommand{
			Name:        "gained",
			Description: "View a player's gains.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "The username of the user to check.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "time",
					Description: "The quantity of how far back to evaluate",
					Required:    false,
					MinValue:    &one,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "unit",
					Description: "The time unit to use for evaluation",
					Required:    false,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "hours",
							Value: "hours",
						},
						{
							Name:  "days",
							Value: "days",
						},
					},
				},
			},
		},
	}
}

func (uh *UserUpdateHandler) HandleCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	_ = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	options := parseOptions(interaction)

	gains, err := uh.gainedService.CalculateUserGains(options.Username, options.Time, options.Unit)
	if err != nil {
		message := "Something went really wrong."
		if errors.Is(err, gain.ErrHiscoreTimeout) {
			message = "Could not get a response from Runescape hiscores in time. Blame Jagex."
		} else {
			uh.logger.ErrorArgs(context.Background(), "An error occurred while calculating user gains: %v", err)
		}

		_, _ = session.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
			Content: &message,
		})

		return
	}

	description := buildDescription(options.Username, gains)
	_, _ = session.InteractionResponseEdit(interaction.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			{
				URL:         "https://api.hazelmere.xyz",
				Type:        discordgo.EmbedTypeRich,
				Title:       fmt.Sprintf("%s gains (%d %s)", options.Username, options.Time, options.Unit),
				Description: description,
				Color:       2719929,
			},
		},
	})
}

func parseOptions(interaction *discordgo.InteractionCreate) UserUpdateOptions {
	options := interaction.ApplicationCommandData().Options
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, option := range options {
		optionsMap[option.Name] = option
	}

	username := optionsMap["username"].StringValue()
	time := 1
	unit := "days"

	if timeValue, ok := optionsMap["time"]; ok {
		time = int(timeValue.IntValue())
	}

	if unitValue, ok := optionsMap["unit"]; ok {
		unit = unitValue.StringValue()
	}

	return UserUpdateOptions{
		Username: username,
		Time:     time,
		Unit:     gain.TimeUnitFromString(unit),
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
