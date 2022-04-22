/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/borakasmer/exchange-cli/parser"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Bu Cli Tool ile, kur bilgileri anlık çekilir",
	Long: `
Herhangi bir tanımlama yapılmaz ise, Dolar "-d", eğer tanımlama yapılır ise 
Euro "-e" veya Sterlin "-s" kur bilgileri Doviz.com anlık olarak, Parse Edilerek ekrana basılır. 

**Doviz.com'da bir sorun olması durumunda, bu servis hizmet veremez!!'

For example:
 .exchange get => Dolar [2022-04-18 22:16:26] : 14.6505₺
 .exchange get -e => Euro [2022-04-18 22:16:49] : 15.8084₺
`,
	/*Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},*/
	Run: func(cmd *cobra.Command, args []string) {
		isDolar, _ := cmd.Flags().GetBool("dolar")
		isEuro, _ := cmd.Flags().GetBool("euro")
		isSterlin, _ := cmd.Flags().GetBool("sterlin")

		if isDolar {
			getDolar()
		} else if isEuro {
			getEuro()
		} else if isSterlin {
			getSterlin()
		} else {
			getDolar()
		}
	},
}

func getDolar() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("DOLAR")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)
	current_time := time.Now()
	fmt.Printf("Dolar [%s] : %s₺", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
}
func getEuro() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("EURO")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)

	current_time := time.Now()
	fmt.Printf("Euro [%s] : %s₺", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
}

func getSterlin() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("STERLİN")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)

	current_time := time.Now()
	fmt.Printf("Sterlin [%s] : %s₺", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("dolar", "d", false, "Get Dolar Currency")
	getCmd.Flags().BoolP("euro", "e", false, "Get Euro Currency")
	getCmd.Flags().BoolP("sterlin", "s", false, "Get Sterlin Currency")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
