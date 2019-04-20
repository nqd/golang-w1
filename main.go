package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nqd/lab1/shortener"
)

func main() {
	var cfgAlias, cfgURL string
	var runPort int
	var rootList bool

	shorten, err := shortener.NewShortener("./record.yaml")
	if err != nil {
		log.Fatalln(err)
		os.Exit(-1)
	}

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Start a HTTP server for shortening",
		Long: `start a HTTP server for shortening.
You will be redirected to the original page.`,
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
			fmt.Println("Running at port", runPort)
		},
	}
	cmdRun.Flags().IntVarP(&runPort, "port", "p", 3000, "HTTP server port")

	var cmdConfigure = &cobra.Command{
		Use:   "configure",
		Short: "append to the shorten list",
		Long: `append to the shorten list.
Configure create new redirect item.`,
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			err := shorten.Add(cfgAlias, cfgURL)
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Println("done!")
		},
	}

	cmdConfigure.Flags().StringVarP(&cfgAlias, "alias", "a", "", "alias key")
	cmdConfigure.Flags().StringVarP(&cfgURL, "url", "u", "", "url for the alias")

	// cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			if rootList {
				shorten.List()
			}
		},
	}
	rootCmd.Flags().BoolVarP(&rootList, "list", "l", false, "url for the alias")

	rootCmd.AddCommand(cmdRun, cmdConfigure)

	rootCmd.Execute()
}
