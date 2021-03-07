package bot

import (
	"github.com/sirupsen/logrus"

	"github.com/AlexbavGamer/CSSVerificationRelay/config"
	"github.com/bwmarrin/discordgo"
)

var RelayBot *discordgo.Session

func Initialize() {
	session, err := discordgo.New("Bot " + config.Config.Bot.Token)

	if err != nil {
		logrus.WithField("error", err).Fatal("Unable to initiate bot session")
	}

	RelayBot = session
}
