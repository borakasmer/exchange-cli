/*
Copyright © 2022 BORA KASMER <bora@borakasmer.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Bu Cli Tool ile, kur bilgileri anlık çekilir",
	Long: `
Herhangi bir tanımlama yapılmaz ise, Dolar "-d", eğer tanımlama yapılır ise 
Euro "-e" veya Sterlin "-s" kur bilgileri Doviz.com anlık olarak, Parse Edilerek ekrana basılır. 

**Doviz.com'da bir sorun olması durumunda, bu servis hizmet veremez!!'

For example:
 .exchange get => Dolar [2022-04-18 22:16:26] : 14.6505₺
 .exchange get -e => Euro [2022-04-18 22:16:49] : 15.8084₺
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.exchange.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
