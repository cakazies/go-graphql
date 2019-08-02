package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	conf "github.com/local/go-graphql/application/config"
	"github.com/local/go-graphql/routes"
	"github.com/local/go-graphql/utils"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "GO-Graphql",
	Short: "Tutorial golang with graphql",
	Long:  `tutorial golang with graphql and some plugins`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("go-graphql is avaible running")
		routes.Route()
	},
}

func init() {
	cobra.OnInitialize(splash, InitViper, conf.Connect)
}

func splash() {
	fmt.Println(`
________ ________              __________________    _____ __________  ___ ___ ________  .____     
/  _____/ \_____  \            /  _____/\______   \  /  _  \\______   \/   |   \\_____  \ |    |    
/   \  ___  /   |   \   ______ /   \  ___ |       _/ /  /_\  \|     ___/    ~    \/  / \  \|    |    
\    \_\  \/    |    \ /_____/ \    \_\  \|    |   \/    |    \    |   \    Y    /   \_/.  \    |___ 
\______  /\_______  /          \______  /|____|_  /\____|__  /____|    \___|_  /\_____\ \_/_______ \
		\/         \/                  \/        \/         \/                \/        \__>       \/
	`)
	// http://patorjk.com/software/taag/#p=display&f=Graffiti
}

// InitViper from file toml
func InitViper() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
	log.Println("Using Config File: ", viper.ConfigFileUsed())
}

// Execute from Cobra Firsttime
func Execute() {
	err := rootCmd.Execute()
	utils.FailError(err, "Error Execute RootCMD")
}
