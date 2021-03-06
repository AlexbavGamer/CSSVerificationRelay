package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AlexbavGamer/CSSVerificationRelay/bot"
	"github.com/AlexbavGamer/CSSVerificationRelay/config"

	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
)

var (
	action string
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)

	flag.StringVar(&action, "service", "", "Install, uninstall, start, stop, restart")
	flag.StringVar(&config.Path, "config", "config.toml", "Path to the config file")

	flag.Parse()
}

type program struct{}

func (p *program) Start(s service.Service) error {
	logrus.Infof("Server is now running on version %s. Press CTRL-C to exit.", config.SCRVER)

	config.ParseString()

	bot.Initialize()

	return nil
}

func (p *program) Stop(s service.Service) error {
	logrus.Info("Received exit signal. Terminating.")

	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "source-chat-relay",
		DisplayName: "Source Chat Relay",
		Description: "Service for Source Chat Relay",
	}

	flag.VisitAll(func(f *flag.Flag) {
		// ignore our own flags
		if f.Name == "service" {
			return
		}

		// ignore flags with default value
		if f.Value.String() == f.DefValue {
			return
		}

		svcConfig.Arguments = append(svcConfig.Arguments, "-"+f.Name+"="+f.Value.String())
	})

	s, err := service.New(&program{}, svcConfig)

	if err != nil {
		exit(err)
	}

	if action != "" {
		exit(actionHandler(action, s))
	}

	exit(s.Run())
}

func exit(err error) {
	if err != nil {
		logrus.Fatal(err)
	}

	os.Exit(0)
}

func actionHandler(action string, s service.Service) error {
	if action != "status" {
		return service.Control(s, action)
	}

	code, _ := s.Status()

	switch code {
	case service.StatusUnknown:
		fmt.Println("Service is not installed.")
	case service.StatusStopped:
		fmt.Println("Service is not running.")
	case service.StatusRunning:
		fmt.Println("Service is running.")
	default:
		fmt.Println("Error: ", code)
	}

	return nil
}
