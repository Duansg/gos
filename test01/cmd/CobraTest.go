package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/contrib/server/http"
	"github.com/spf13/cobra"
	"os"
)

var (
	// dataDirPath save all answer application data in this directory. like config file, upload file...
	dataDirPath string
	// dumpDataPath dump data path
	dumpDataPath string
	// plugins needed to build in answer application
	buildWithPlugins []string
	// build output path
	buildOutput string
	// This config is used to upgrade the database from a specific version manually.
	// If you want to upgrade the database to version 1.1.0, you can use `answer upgrade -f v1.1.0`.
	upgradeVersion string
	// The fields that need to be set to the default value
	configFields []string
)
var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "Duansg",
		Short: "Duansg open source Q&A community.",
		Long:  `Duansg open source Q&A community.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ssss")
		},
	}

	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Answer is a minimalist open source Q&A community.",
		Long:  `Answer is a minimalist open source Q&A community.`,
		Run: func(cmd *cobra.Command, args []string) {
			str, _ := cmd.Flags().GetString("with")
			str1, _ := cmd.Flags().GetString("witha")
			fmt.Println("config file path: " + str)
			fmt.Println("config file path: " + str1)
			fmt.Println("Duansg is starting..........................")
			newApplication()
		},
	}
)

func newApplication() {
	r := gin.New()
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })
	app := pacman.NewApp(
		pacman.WithName("Duansg"),
		pacman.WithVersion("1.0.0"), pacman.WithServer(http.NewServer(r, "127.0.0.1:8080")))
	app.Run(context.Background())
}

func init() {
	showCmd.Flags().String("with", "", "这是晓断的测试参数")
	showCmd.Flags().String("witha", "", "这是晓断的测试参数")
	rootCmd.AddCommand(showCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
