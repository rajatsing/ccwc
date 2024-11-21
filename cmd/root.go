/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "number of bytes in a file",
	Long:  `input a file path and find out number of bytes in that file`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		bytesInfile, _ := cmd.Flags().GetBool("bytes")
		lines, _ := cmd.Flags().GetBool("lines")
		words, _ := cmd.Flags().GetBool("words")
		chars, _ := cmd.Flags().GetBool("chars")

		var data []byte
		var err error
		var fileName string

		pipeInput := os.Stdin
		pipeInputInfo, err := pipeInput.Stat()
		size := pipeInputInfo.Size()

		if size > 0 {
			data, err = io.ReadAll(pipeInput) // read from the pipe
			if err != nil {
				panic(err)
			}
		} else {
			if len(args) == 0 {
				fmt.Println("please enter filename")
				os.Exit(1)
			}
			fileName = args[0]
			data, err = os.ReadFile(fileName) // this reads the file in bytes
			if err != nil {
				panic(err)
			}
		}

		fileSize := len(data)
		linesinFile := strings.Count(string(data), "\n")   // splits data into each line
		wordsInFile := strings.Fields(string(data))        // splits the data in words
		totalChars := utf8.RuneCountInString(string(data)) // counts runes/everything

		if bytesInfile {
			fmt.Println("file size in bytes:", fileSize)
		} else if lines {
			fmt.Println("number of lines:", linesinFile)
		} else if words {
			fmt.Println("number of words:", len(wordsInFile))
		} else if chars {
			fmt.Println("total number of characters are:", totalChars)
		} else {
			fmt.Println(fileSize, linesinFile, len(wordsInFile))
		}
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("bytes", "c", false, "number of bytes in a file")
	rootCmd.Flags().BoolP("lines", "l", false, "number of lines in a file")
	rootCmd.Flags().BoolP("words", "w", false, "number of words in a file")
	rootCmd.Flags().BoolP("chars", "m", false, "number of chars in a file")

}
