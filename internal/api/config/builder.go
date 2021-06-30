package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

const (
	kafkaAddr   = "kafka-addr"
	port        = "port"
	logLevel    = "log-level"
	databaseURL = "databaseURL"
)

//Flags .
type Flags struct {
	KafkaAddr   string
	LogLevel    string
	Port        string
	databaseURL string
}

//APIBuilder .
type APIBuilder struct {
	*Flags
}

//AddFlags .
func AddFlags(flags *pflag.FlagSet) {
	flags.StringP(port, "p", "7070", "[optional] Custom port for accessing this services. Default: 7070")
	flags.StringP(logLevel, "l", "info", "[optional] Sets the Log Level to one of seven (trace, debug, info, warn, error, fatal, panic). Default: info")
	flags.StringP(databaseURL, "d", "postgresql://user2020:pass2020@postgres:5432/stop-analyzing-api", "[optional] Sets the database connection string. Default: postgresql://user2020:pass2020@postgres:5432/stop-analyzing-api")
}

//Init .
func (b *APIBuilder) Init(v *viper.Viper) *APIBuilder {
	flags := new(Flags)

	flags.Port = v.GetString(port)
	flags.LogLevel = v.GetString(logLevel)
	flags.databaseURL = v.GetString(databaseURL)

	os.Setenv("DATABASE_URL", flags.databaseURL)

	flags.check()
	b.Flags = flags
	return b
}

func (flags *Flags) check() {
	logrus.Infof("Flags: '%v'", flags)

	requiredFlags := []struct {
		value string
		name  string
	}{}
	// {
	// 	{flags.KafkaAddr, kafkaAddr},
	// }

	var errMsg string

	for _, flag := range requiredFlags {
		if flag.value == "" {
			errMsg += fmt.Sprintf("\n\t%v", flag.name)
		}
	}

	if errMsg != "" {
		errMsg = "The following flags are missing: " + errMsg
		panic(errMsg)
	}
}
