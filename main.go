package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vagrant-wrapper",
	Short: "A Vagrant CLI wrapper",
	Long:  "A Vagrant CLI wrapper that generates a Vagrantfile based on user input.",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println(color.GreenString("Welcome to the Vagrant CLI wrapper!"))
		fmt.Println("Let's generate a Vagrantfile.")

		// Prompt for simplified or full configuration
		fmt.Println("Do you want a simplified configuration or full configuration?")
		fmt.Println("1. Simplified")
		fmt.Println("2. Full")
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "1" {
			generateSimplifiedConfig()
		} else if choice == "2" {
			generateFullConfig(reader)
		} else {
			fmt.Println(color.RedString("Invalid choice. Please try again."))
		}
	},
}

func generateSimplifiedConfig() {
	fmt.Println(color.CyanString("Generating simplified configuration..."))

	// Generate Vagrantfile content
	vagrantfileContent := `
Vagrant.configure("2") do |config|
  config.vm.box = "hashicorp/bionic64"
  # Add other simplified configuration options here
end
`
	writeVagrantfile(vagrantfileContent)
}

func generateFullConfig(reader *bufio.Reader) {
	fmt.Println(color.CyanString("Generating full configuration..."))

	// Prompt for Vagrantfile options
	fmt.Println("Please provide the following Vagrantfile options:")

	// Prompt for box name
	fmt.Print("Box name: ")
	boxName, _ := reader.ReadString('\n')
	boxName = strings.TrimSpace(boxName)

	// Prompt for other options...
	// Add more prompts for other Vagrantfile options

	// Generate Vagrantfile content
	vagrantfileContent := fmt.Sprintf(`
Vagrant.configure("2") do |config|
  config.vm.box = "%s"
  # Add other full configuration options here
end
`, boxName)

	writeVagrantfile(vagrantfileContent)
}

func writeVagrantfile(content string) {
	// Write Vagrantfile
	vagrantfile, err := os.Create("Vagrantfile")
	if err != nil {
		fmt.Println(color.RedString("Failed to create Vagrantfile:", err))
		return
	}
	defer vagrantfile.Close()

	_, err = vagrantfile.WriteString(content)
	if err != nil {
		fmt.Println(color.RedString("Failed to write Vagrantfile:", err))
		return
	}

	fmt.Println(color.GreenString("Vagrantfile generated successfully!"))
}

func main() {
	// Add timestamp to logs
	color.New(color.FgYellow).Println("Program started at", time.Now().Format("2006-01-02 15:04:05"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.RedString(err.Error()))
		os.Exit(1)
	}

	// Add timestamp to logs
	color.New(color.FgYellow).Println("Program ended at", time.Now().Format("2006-01-02 15:04:05"))
}
