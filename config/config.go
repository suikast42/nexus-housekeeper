package config

import (
	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
	flag "github.com/spf13/pflag"
	"os"
	"strings"
)

var conf = koanf.Conf{
	Delim:       ".",
	StrictMerge: true,
}
var k = koanf.NewWithConf(conf)

func LoadConfig(log *zerolog.Logger) AppConfig {
	// Use the POSIX compliant pflag lib instead of Go's flag lib.
	f := flag.NewFlagSet("nexushk", flag.ContinueOnError)
	f.Usage = func() {
		log.Fatal().Msgf("Usage:\n%s", f.FlagUsages())
	}

	// Path to one or more config files to load into koanf along with some config params.
	f.StringSlice("config", []string{"config/mock.hcl"}, "path to one or more .hcl config files")

	cFiles, _ := f.GetStringSlice("config")
	for _, c := range cFiles {
		if err := k.Load(file.Provider(c), hcl.Parser(true)); err != nil {
			log.Warn().Msgf("error loading config file %s. Error: %s", c, err)
		}
	}

	// Parse the flags
	f.String("loglevel", "info", "log level of the application")

	f.String("nexusServer.scheme", "https", "http or https")
	f.String("nexusServer.host", "nexus.cloud.private", "host of the nexus server")
	f.String("nexusServer.port", "443", "port of the grafana server")
	f.String("nexusServer.username", "nouser", "api user for the nexus server")
	f.String("nexusServer.password", "nopass", "api pass for the nexus server")
	f.Int("nexusServer.keepImages.default", 1, "How many images should be left after deletion for a version of an image")
	err := f.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	// Load environment variables and merge into the loaded config.
	// "NEXUS_HK" is the prefix to filter the env vars by.
	// "." is the delimiter used to represent the key hierarchy in env vars.
	// The (optional, or can be nil) function can be used to transform
	// the env var names, for instance, to lowercase them.
	if err := k.Load(env.Provider("NEXUS_HK", ".", func(s string) string {
		return strings.Replace(strings.TrimPrefix(s, "NEXUS_HK"), "_", ".", -1)
	}), nil); err != nil {
		log.Fatal().Err(err).Msgf("error loading property from environment variable")
	}

	// The bundled posflag.Provider takes a flagset from the spf13/pflag lib.
	// Passing the Koanf instance to posflag helps it deal with default command
	// line flag values that are not present in conf maps from previously loaded
	// providers.
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		log.Fatal().Err(err).Msgf("error loading config")
	}
	config := AppConfig{}
	unmarshalConf := koanf.UnmarshalConf{Tag: "koanf", FlatPaths: false}
	if err := k.UnmarshalWithConf("", &config, unmarshalConf); err != nil {
		log.Fatal().Err(err).Msgf("error UnmarshalWithConf to AppConfig")
	}

	level, err := zerolog.ParseLevel(config.LogLevel)

	if err != nil {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Warn().Msgf("Can't parse supported loglevel %s use default %s", config.LogLevel, zerolog.InfoLevel)
		config.LogLevel = zerolog.InfoLevel.String()
	} else {
		zerolog.SetGlobalLevel(level)
	}
	return config

}
