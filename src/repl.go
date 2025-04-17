package src

import (
	"bufio"
	"fmt"
	"os"
)

func InitRepl() map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Paginate the next page the PokeAPI",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Paginate the previous page the PokeAPI",
			callback:    commadMapBack,
		},
	}

}

type config struct {
	url      string
	commands *map[string]cliCommand
	next     string
	prev     string
	base     string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func StartRepl() {

	reader := bufio.NewScanner(os.Stdin)
	commands := InitRepl()

	cfg := &config{
		next: "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0",
		prev: "",
	}

	cfg.commands = &commands

	for {

		fmt.Print("Pokedex > ")
		reader.Scan()

		command := cleanInput(reader.Text())[0]

		if len(command) == 0 {
			continue
		}

		cmd, exist := commands[command]

		if !exist {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg)

		if err != nil {
			fmt.Println("Error on command execution:", err)
			continue
		}

	}

}
