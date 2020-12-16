package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/evdatsion/aphelion-dpos-bft/config"
	"github.com/evdatsion/aphelion-dpos-bft/libs/cli"
	tmflags "github.com/evdatsion/aphelion-dpos-bft/libs/cli/flags"
	"github.com/evdatsion/aphelion-dpos-bft/libs/log"
)

var (
	config = cfg.DefaultConfig()
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

func init() {
	registerFlagsRootCmd(RootCmd)
}

func registerFlagsRootCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().String("log_level", config.LogLevel, "Log level")
}

// ParseConfig retrieves the default environment configuration,
// sets up the Aphelion root and ensures that the root exists
func ParseConfig() (*cfg.Config, error) {
	conf := cfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	conf.SetRoot(conf.RootDir)
	cfg.EnsureRoot(conf.RootDir)
	if err = conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("Error in config file: %v", err)
	}
	return conf, err
}

// RootCmd is the root command for Aphelion core.
var RootCmd = &cobra.Command{
	Use:   "aphelion",
	Short: "Aphelion Core (BFT Consensus) in Go",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}
		config, err = ParseConfig()
		if err != nil {
			return err
		}
		if config.LogFormat == cfg.LogFormatJSON {
			logger = log.NewTMJSONLogger(log.NewSyncWriter(os.Stdout))
		}
		logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel())
		if err != nil {
			return err
		}
		if viper.GetBool(cli.TraceFlag) {
			logger = log.NewTracingLogger(logger)
		}
		logger = logger.With("module", "main")
		return nil
	},
}
