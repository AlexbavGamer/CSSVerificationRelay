package bot

import (
	"github.com/Necroforger/dgrouter/exrouter"
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
	session.AddHandler(ready)

	session.State.TrackEmojis = false
	session.State.TrackPresences = false
	session.State.TrackVoice = false

	err = session.Open()

	if err != nil {
		logrus.WithField("Error", err).Fatal("Unable to open bot session")
	}

	router := exrouter.New()

	session.AddHandler(func(session *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.Bot && !config.Config.Bot.ListenToBots {
			return
		}

	})

}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	go Listen()

	logrus.WithFields(logrus.Fields{
		"Username":    event.User.Username,
		"Guild Count": len(event.Guilds),
	}).Info("Bot is now running")
}
