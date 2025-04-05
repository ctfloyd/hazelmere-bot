package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ctfloyd/hazelmere-api/src/pkg/api"
	"github.com/ctfloyd/hazelmere-bot/src/internal"
	"github.com/ctfloyd/hazelmere-bot/src/internal/constant"
	"github.com/ctfloyd/hazelmere-bot/src/internal/initialize"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_logger"
	"log"
	"os"
	"os/signal"
)

var OrderedSkills = []api.ActivityType{
	api.ActivityTypeAttack,
	api.ActivityTypeHitpoints,
	api.ActivityTypeMining,
	api.ActivityTypeStrength,
	api.ActivityTypeAgility,
	api.ActivityTypeSmithing,
	api.ActivityTypeDefence,
	api.ActivityTypeHerblore,
	api.ActivityTypeFishing,
	api.ActivityTypeRanged,
	api.ActivityTypeThieving,
	api.ActivityTypeCooking,
	api.ActivityTypePrayer,
	api.ActivityTypeCrafting,
	api.ActivityTypeFiremaking,
	api.ActivityTypeMagic,
	api.ActivityTypeFletching,
	api.ActivityTypeWoodcutting,
	api.ActivityTypeRunecraft,
	api.ActivityTypeSlayer,
	api.ActivityTypeFarming,
}

func main() {
	cfg := hz_config.NewConfigWithAutomaticDetection()
	err := cfg.Read()
	if err != nil {
		panic(err)
	}
	discord, err := discordgo.New(cfg.ValueOrPanic("token"))

	if err != nil {
		panic(err)
	}

	one := 1.0
	userUpdate := discordgo.ApplicationCommand{
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
	}

	logger := hz_logger.NewZeroLogAdapater(hz_logger.LogLevelDebug)
	hz := initialize.InitializeHazelmereClient(cfg, logger)

	gainedService := internal.NewGainedService(logger, hz)

	err = discord.Open()
	if err != nil {
		panic(err)
	}
	defer discord.Close()

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is now running.  Press CTRL-C to exit.")
	})

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "", &userUpdate)
	if err != nil {
		panic(err)
	}

	discord.AddHandler(func(discord *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "gained" {
			handleUserUpdate(discord, i, logger, gainedService)
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}

func handleUserUpdate(discord *discordgo.Session, i *discordgo.InteractionCreate, logger hz_logger.Logger, gs *internal.GainedService) {
	options := i.ApplicationCommandData().Options
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, option := range options {
		optionsMap[option.Name] = option
	}

	username := optionsMap["username"].StringValue()
	time := 1
	unit := "days"

	if tVal, ok := optionsMap["time"]; ok {
		time = int(tVal.IntValue())
	}

	if uval, ok := optionsMap["unit"]; ok {
		unit = uval.StringValue()
	}

	_ = discord.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	tu := internal.TimeUnitFromString(unit)
	gains, err := gs.CalculateUserGains(username, time, tu)
	if err != nil {
		if errors.Is(err, internal.ErrHiscoreTimeout) {
			message := "Could not get a response from Runescape hiscores in time. Blame Jagex."
			_, _ = discord.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &message,
			})
			return
		}
		logger.ErrorArgs(context.Background(), "An error occurred while calculating user gains: %v", err)
		message := "Something went really wrong."
		_, _ = discord.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &message,
		})
		return
	}

	description := fmt.Sprintf("%s **%s**\n",
		constant.Emojis[api.ActivityTypeOverall],
		gains.Skills[api.ActivityTypeOverall].Amount,
	)

	for _, skill := range OrderedSkills {
		emoji, ok := constant.Emojis[skill]
		if !ok {
			emoji = constant.Emojis[api.ActivityTypeUnknown]
		}

		gain, ok := gains.Skills[skill]
		if ok {
			description += fmt.Sprintf("%s %s\n", emoji, gain.Amount)
		}
	}

	_, _ = discord.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			{
				URL:         "https://api.hazelmere.xyz",
				Type:        discordgo.EmbedTypeRich,
				Title:       fmt.Sprintf("%s gains (%d %s)", username, time, unit),
				Description: description,
				Color:       2719929,
			},
		},
	})
}
