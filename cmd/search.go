/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"log"
	"os"
	"strings"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A CLI with Gemini integrated in it",
	Long:  `This CLI is built using Go, Cobra library as a CLI toolkit, and Gemini API. `,
	Args:  cobra.MinimumNArgs(1), // Minimum 1 arg required
	Run: func(cmd *cobra.Command, args []string) {
		getResponse(args)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func getResponse(args []string) {

	// Creating a sentence out of a slice
	userArgs := strings.Join(args[0:], " ")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(userArgs))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Candidates[0].Content.Parts[0])
}
