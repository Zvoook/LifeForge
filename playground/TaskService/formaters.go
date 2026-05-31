package main

import "fmt"

type Color string

const (
	Red        Color = "\033[31m"
	Green      Color = "\033[32m"
	Yellow     Color = "\033[33m"
	Blue       Color = "\033[36m"
	ResetColor Color = "\033[0m"
)

const Menu = "╔══════════════════════════════════════╗ \n║          LifeForge Task CLI          ║ \n╠══════════════════════════════════════╣ \n║ 1. Create task                       ║ \n║ 2. Show all tasks                    ║ \n║ 3. Show tasks by area                ║ \n║ 4. Find task by ID                   ║ \n║ 5. Complete task                     ║ \n║ 6. Change task priority              ║ \n║ 7. Delete task                       ║ \n║ 8. Show dashboard                    ║ \n║ 0. Exit                              ║ \n╚══════════════════════════════════════╝"
const ClearScreenCommand = "\033[2J\033[H"
const pressEnterRequire = "\nPress Enter to continue...\n"

func printMenu() {
	fmt.Println(Menu)
}

func clearScreen() {
	fmt.Print(ClearScreenCommand)
	fmt.Print("\n")
}

func (c *CLI) waitForEnter() {
	c.readLine(pressEnterRequire)
}

func printSuccess(prompt string) {
	fmt.Printf("%s✓ %s%s\n", Green, ResetColor, prompt)
}

func printError(err error) {
	if err != nil {
		fmt.Printf("%s✗ Error: %s%v\n", Red, ResetColor, err)
	}
}

func printInfo(prompt string) {
	fmt.Printf("%s[i] %s%s\n", Blue, ResetColor, prompt)
}
