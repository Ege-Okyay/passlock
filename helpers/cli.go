package helpers

import (
	"fmt"
	"sort"

	"github.com/Ege-Okyay/pass-lock/types"
)

func PrintHelp(commands map[string]types.Command) {
	PrintBanner("Passlock - Secure Key/Value Store")
	fmt.Println("USAGE:")
	fmt.Println("\tpasslock <command> [arguments]")

	fmt.Println("AVAILABLE COMMANDS:")
	for _, cmd := range commands {
		fmt.Printf("\t%-15s %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println("FLAGS:")
	fmt.Println("\t-h, --help")

	fmt.Println("EXAMPLES:")
	fmt.Println("\tpasslock setup \tInitialize the vault with a master password.")
	fmt.Println("\tpasslock set apiKey secret123 \tStore a new key-value pair.")
	fmt.Println("\tpasslock get apiKey \tRetrive the value for the specified key.")
	fmt.Println("\tpasslock delete apiKey \tDelete the specified key from the vault.")

	fmt.Println("\nUse 'passlock <command> --help' for more information about a command.")
}

func PrintCommandHelp(cmd types.Command) {
	fmt.Println("NAME:")
	fmt.Printf("\t%s\n", cmd.Name)

	fmt.Println("DESCRIPTION:")
	fmt.Printf("\t%s\n", cmd.Description)

	fmt.Println("USAGE:")
	fmt.Printf("\t%s\n", cmd.Usage)
}

func FindClosestCommands(commands map[string]types.Command, cmdName string, maxResults int) []string {
	str1 := []rune(cmdName)
	var commandDistances []types.CommandDistance

	for name := range commands {
		str2 := []rune(name)
		score := Levenshtein(str1, str2)

		commandDistances = append(commandDistances, types.CommandDistance{Name: name, Score: score})
	}

	sort.Slice(commandDistances, func(i, j int) bool {
		return commandDistances[i].Score < commandDistances[j].Score
	})

	var closestCommands []string
	for i := 0; i < min(maxResults, len(commandDistances)); i++ {
		closestCommands = append(closestCommands, commandDistances[i].Name)
	}

	return closestCommands
}

func HandleUnknownCommand(commands map[string]types.Command, cmdName string) {
	closestCommands := FindClosestCommands(commands, cmdName, 5)

	fmt.Printf("passlock: '%s' is not a passlock command. See 'passlock help'.\n", cmdName)
	if len(closestCommands) > 0 {
		fmt.Println("\nThe most similar commands are:")
		for _, command := range closestCommands {
			fmt.Printf("\t%s\n", command)
		}
	}
}

func IsHelpFlag(arg string) bool {
	return arg == "-h" || arg == "--help" || arg == "help"
}
