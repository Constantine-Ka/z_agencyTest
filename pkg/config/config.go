package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	VP *viper.Viper
	L  *logrus.Logger
}
