package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	conf "github.com/cakazies/go-graphql/application/config"
	"github.com/cakazies/go-graphql/routes"
	"github.com/cakazies/go-graphql/utils"
	"github.com/spf13/cobra"
)

var (
	cfgFile string // cek for file toml or viper save in this variable
)

// rootCmd call in this file in function execute
var rootCmd = &cobra.Command{
	Use:   "GO-Graphql",
	Short: "Tutorial golang with graphql",
	Long:  `tutorial golang with graphql and some plugins`,
	Run: func(cmd *cobra.Command, args []string) {
		routes.Route()
	},
}

func init() {
	// call first time for this repo declare viper, connect DB and splash
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
	// this splash i get in website : http://patorjk.com/software/taag/#p=display&f=Graffiti
}

// this function call in init() and this file
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

// Execute function call in firsttime in main.go
func Execute() {
	err := rootCmd.Execute()
	utils.FailError(err, "Error Execute RootCMD")
}
