package initialize

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ctfloyd/hazelmere-bot/src/internal/discord"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
)

func InitializeDiscord(cfg *hz_config.Config) *discord.Bot {
	session, err := discordgo.New(cfg.ValueOrPanic("token"))
	if err != nil {
		panic(err)
	}
	return discord.NewBot(session)
}
