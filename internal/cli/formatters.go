package cli

import (
	"fmt"
	"strings"

	"github.com/Zvoook/lifeforge/internal/task"
)

type Color string

const (
	Red        Color = "\033[31m"
	Green      Color = "\033[32m"
	Yellow     Color = "\033[33m"
	Blue       Color = "\033[36m"
	ResetColor Color = "\033[0m"
)

const Menu = `
╔════════════════════════════════════════════════════╗
║                  LifeForge Task CLI                ║
║        Build your day. Track your progress.        ║
╠════════════════════════════════════════════════════╣
║ View                                               ║
║   1. Show all tasks                                ║
║   2. Show tasks by area                            ║
║   3. Show tasks by status                          ║
║   4. Find task by ID                               ║
╠════════════════════════════════════════════════════╣
║ Create & Edit                                      ║
║   5. Create task                                   ║
║   6. Edit task                                     ║
║   7. Complete task                                 ║
╠════════════════════════════════════════════════════╣
║ Planning                                           ║
║   8. Forge daily plan                              ║
║   9. Show dashboard                                ║
╠════════════════════════════════════════════════════╣
║ Deleting                                           ║
║  10. Delete task                                   ║
║  11. Clear all tasks                               ║
╠════════════════════════════════════════════════════╣
║   0. Exit                                          ║
╚════════════════════════════════════════════════════╝
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

func PrintSuccess(prompt string) {
	fmt.Printf("%s✓ %s%s\n", Green, prompt, ResetColor)
}

func PrintError(err error) {
	if err != nil {
		fmt.Printf("%s✗ Error: %v%s\n", Red, err, ResetColor)
	}
}

func PrintInfo(prompt string) {
	fmt.Printf("%s[i] %s%s\n", Blue, prompt, ResetColor)
}

func printTasksTable(tasks []task.Task) {
	if len(tasks) == 0 {
		PrintInfo("No tasks found")
		return
	}

	const (
		idWidth       = 4
		areaWidth     = 12
		titleWidth    = 30
		statusWidth   = 13
		priorityWidth = 8
		minutesWidth  = 8
	)

	border := "+" +
		strings.Repeat("-", idWidth+2) + "+" +
		strings.Repeat("-", areaWidth+2) + "+" +
		strings.Repeat("-", titleWidth+2) + "+" +
		strings.Repeat("-", statusWidth+2) + "+" +
		strings.Repeat("-", priorityWidth+2) + "+" +
		strings.Repeat("-", minutesWidth+2) + "+"

	fmt.Println(border)
	fmt.Printf(
		"| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		idWidth, "ID",
		areaWidth, "Area",
		titleWidth, "Title",
		statusWidth, "Status",
		priorityWidth, "Priority",
		minutesWidth, "Minutes",
	)
	fmt.Println(border)

	for _, task := range tasks {
		fmt.Printf(
			"| %-*d | %-*s | %-*s | %-*s | %-*d | %-*d |\n",
			idWidth, task.ID,
			areaWidth, task.Area.String(),
			titleWidth, trimText(task.Title, titleWidth),
			statusWidth, task.Status.String(),
			priorityWidth, task.Priority,
			minutesWidth, task.EstimatedMinutes,
		)
	}

	fmt.Println(border)
}

func trimText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}

	if maxLength <= 3 {
		return text[:maxLength]
	}

	return text[:maxLength-3] + "..."
}

const DailyPlanHat = `
╔════════════════════════════════════════════════════════════╗
║                    TODAY'S FORGE PLAN                      ║
║              Recommended tasks for your focus              ║
╠════════════════════════════════════════════════════════════╣
║ Available time:   %8d min                             ║
║ Planned time:     %8d min                             ║
║ Free time:        %8d min                             ║
║ Tasks selected:   %8d                                 ║
╚════════════════════════════════════════════════════════════╝

`

func detectMainFocus(plan []task.Task) string {
	counter := make(map[task.Area]int)
	for _, task := range plan {
		counter[task.Area] += task.EstimatedMinutes
	}

	var areaWithMaxTime task.Area
	maxTime := 0

	for area, areaTotalTime := range counter {
		if areaTotalTime > maxTime {
			maxTime = areaTotalTime
			areaWithMaxTime = area
		}
	}

	return areaWithMaxTime.String()
}

func printForgeDailyPlan(plan []task.Task, timeLimit int, totalTime int) {
	fmt.Printf(DailyPlanHat, timeLimit, totalTime, timeLimit-totalTime, len(plan))
	printTasksTable(plan)

	fmt.Printf("\nMain focus: %s\n", detectMainFocus(plan))
	fmt.Println("Recommendation: Start with the first task and do not switch context.")
}
