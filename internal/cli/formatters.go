package cli

import "fmt"

type Color string

const (
	Red        Color = "\033[31m"
	Green      Color = "\033[32m"
	Yellow     Color = "\033[33m"
	Blue       Color = "\033[36m"
	ResetColor Color = "\033[0m"
)

const Menu = `
╔════════════════════════════════════════╗
║          LifeForge Task CLI            ║
╠════════════════════════════════════════╣
║  1. Show all tasks                     ║
║  2. Show tasks by area                 ║
║  3. Show tasks by status               ║
║  4. Find task by ID                    ║
║  5. Create task                        ║
║  6. Change task priority               ║
║  7. Complete task                      ║
║  8. Delete task                        ║
║  9. Show dashboard                     ║
║  10. Clear all tasks                   ║
║  0. Exit                               ║
╚════════════════════════════════════════╝
`

const ClearScreenCommand = "\033[2J\033[H"
const pressEnterRequire = "\nPress Enter to continue...\n"

func printMenu() {
	fmt.Printf("%s%s%s\n", Yellow, Menu, ResetColor)
}

func clearScreen() {
	fmt.Print(ClearScreenCommand)
	fmt.Print("\n")
}

func (c *CLI) waitForEnter() {
	c.readLine(pressEnterRequire)
}

func printSuccess(prompt string) {
	fmt.Printf("%s✓ %s%s\n", Green, prompt, ResetColor)
}

func printError(err error) {
	if err != nil {
		fmt.Printf("%s✗ Error: %s%v\n", Red, err, ResetColor)
	}
}

func printInfo(prompt string) {
	fmt.Printf("%s[i] %s%s\n", Blue, prompt, ResetColor)
}
