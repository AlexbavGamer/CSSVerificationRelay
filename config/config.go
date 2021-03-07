package github.com/AlexbavGamer/CSSVerificationRelay/config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var Config Config_t

var Path string

func ParseString() {
	if _, err := toml.DecodeFile(Path, &Config); err != nil {
		logrus.WithField("error", err).Fatal("Unable to parse config")
	}
}
