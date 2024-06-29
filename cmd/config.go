package cmd

import (
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath string = initConfigPath()

func initConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	// use viper to write config to ~/.coze/config.yml
	configDir := usr.HomeDir + "/.coze"
	// if config dir isn't exsit, create it
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.Mkdir(configDir, 0755); err != nil {
			panic(err)
		}
	}
	return configDir + "/config.yml"
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the bot",
	Long:  `Configure the bot with your OpenAI API key and other settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Your configuration logic here
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			panic(err)
		}
		botid, err := cmd.Flags().GetString("botid")
		if err != nil {
			panic(err)
		}
		userArg, err := cmd.Flags().GetString("user")
		if err != nil {
			panic(err)
		}
		viper.SetConfigFile(configPath)
		viper.Set("token", token)
		viper.Set("botid", botid)
		viper.Set("user", userArg)
		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	},
}

func init() {
	// Add flags and configuration options to ConfigCmd
	ConfigCmd.PersistentFlags().String("token", "", "Your OpenAI API token")
	ConfigCmd.PersistentFlags().String("botid", "7385882762747936785", "Your OpenAI bot ID")
	ConfigCmd.PersistentFlags().String("user", "29032201862555", "Your OpenAI user")
	ConfigCmd.PersistentFlags().Bool("stream", false, "Enable streaming mode")
}
