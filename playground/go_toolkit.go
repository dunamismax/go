// go_toolkit.go
// A Nord dark–themed Go CLI toolkit for development on MacOS.
// Features include: navigating to a Go project directory, building a project,
// running the built executable in a new Terminal window, testing, formatting, and more.

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	// Nord theme ANSI color codes (truecolor)
	nord0  = "\033[38;2;46;52;64m"    // #2E3440
	nord1  = "\033[38;2;59;66;82m"    // #3B4252
	nord2  = "\033[38;2;67;76;94m"    // #434C5E
	nord3  = "\033[38;2;76;86;106m"   // #4C566A
	nord4  = "\033[38;2;216;222;233m" // #D8DEE9
	nord5  = "\033[38;2;229;233;240m" // #E5E9F0
	nord6  = "\033[38;2;236;239;244m" // #ECEFF4
	nord7  = "\033[38;2;143;188;187m" // #8FBCBB
	nord8  = "\033[38;2;136;192;208m" // #88C0D0
	nord9  = "\033[38;2;129;161;193m" // #81A1C1
	nord10 = "\033[38;2;94;129;172m"  // #5E81AC
	reset  = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		fmt.Print(nord4 + "Enter your choice: " + reset)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			changeDirectory(reader)
		case "2":
			buildProject(reader)
		case "3":
			runProject()
		case "4":
			runTests()
		case "5":
			runFmt()
		case "6":
			runModTidy()
		case "7":
			runGoRun(reader)
		case "8":
			about()
		case "9":
			fmt.Println(nord7 + "Exiting go_toolkit. Happy coding!" + reset)
			os.Exit(0)
		default:
			fmt.Println(nord10 + "Invalid choice, please try again." + reset)
		}
		fmt.Println() // spacing between operations
	}
}

// printMenu displays the main menu using Nord colors.
func printMenu() {
	clearScreen()
	fmt.Println(nord7 + "========================================" + reset)
	fmt.Println(nord8 + "         go_toolkit - Go CLI Toolkit        " + reset)
	fmt.Println(nord7 + "========================================" + reset)
	fmt.Println(nord4 + "1. Change Working Directory (Go Project)" + reset)
	fmt.Println(nord4 + "2. Build Go Project" + reset)
	fmt.Println(nord4 + "3. Run Built Go Project (New Terminal)" + reset)
	fmt.Println(nord4 + "4. Run Go Tests" + reset)
	fmt.Println(nord4 + "5. Format Code (go fmt)" + reset)
	fmt.Println(nord4 + "6. Tidy Modules (go mod tidy)" + reset)
	fmt.Println(nord4 + "7. Run a Go File (go run)" + reset)
	fmt.Println(nord4 + "8. About go_toolkit" + reset)
	fmt.Println(nord4 + "9. Exit" + reset)
	fmt.Println(nord7 + "========================================" + reset)
}

// clearScreen clears the terminal (works on macOS/Unix systems).
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// changeDirectory lets the user change the current working directory.
func changeDirectory(reader *bufio.Reader) {
	fmt.Print(nord4 + "Enter the path to your Go project: " + reset)
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)
	err := os.Chdir(path)
	if err != nil {
		fmt.Println(nord10 + "Error changing directory: " + err.Error() + reset)
	} else {
		fmt.Println(nord7 + "Directory changed to " + path + reset)
	}
}

// buildProject compiles the provided Go file into an executable named "built_app".
func buildProject(reader *bufio.Reader) {
	fmt.Print(nord4 + "Enter the main Go file to build (e.g., main.go): " + reset)
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	cmd := exec.Command("go", "build", "-o", "built_app", fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(nord10 + "Build failed: " + err.Error() + reset)
	} else {
		fmt.Println(nord7 + "Build successful. Executable created: built_app" + reset)
	}
}

// runProject opens a new Terminal window on macOS and runs the "built_app" executable.
func runProject() {
	if _, err := os.Stat("./built_app"); os.IsNotExist(err) {
		fmt.Println(nord10 + "Executable 'built_app' not found. Please build the project first." + reset)
		return
	}
	// AppleScript to open Terminal and run the executable
	script := `tell application "Terminal" to do script "./built_app"`
	cmd := exec.Command("osascript", "-e", script)
	err := cmd.Run()
	if err != nil {
		fmt.Println(nord10 + "Failed to run project in new Terminal: " + err.Error() + reset)
	} else {
		fmt.Println(nord7 + "Project is running in a new Terminal window." + reset)
	}
}

// runTests executes "go test ./..." in the current directory.
func runTests() {
	cmd := exec.Command("go", "test", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(nord10 + "Tests encountered an error: " + err.Error() + reset)
	}
	fmt.Println(nord7 + string(output) + reset)
}

// runFmt executes "go fmt ./..." to format the code.
func runFmt() {
	cmd := exec.Command("go", "fmt", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(nord10 + "go fmt encountered an error: " + err.Error() + reset)
	}
	// go fmt usually outputs nothing if no files changed.
	if len(output) > 0 {
		fmt.Println(nord7 + string(output) + reset)
	} else {
		fmt.Println(nord7 + "Code formatted successfully." + reset)
	}
}

// runModTidy runs "go mod tidy" to clean up the module dependencies.
func runModTidy() {
	cmd := exec.Command("go", "mod", "tidy")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(nord10 + "go mod tidy encountered an error: " + err.Error() + reset)
	}
	if len(output) > 0 {
		fmt.Println(nord7 + string(output) + reset)
	} else {
		fmt.Println(nord7 + "Modules are tidy." + reset)
	}
}

// runGoRun runs a specific Go file using "go run".
func runGoRun(reader *bufio.Reader) {
	fmt.Print(nord4 + "Enter the Go file to run (e.g., main.go): " + reset)
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	cmd := exec.Command("go", "run", fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(nord10 + "Error running the file: " + err.Error() + reset)
	}
}

// about displays information about the toolkit.
func about() {
	fmt.Println(nord7 + "\n go_toolkit is an all-in-one CLI toolkit for Go developers." + reset)
	fmt.Println(nord7 + " Features include:" + reset)
	fmt.Println(nord4 + "  - Navigating to a Go project directory" + reset)
	fmt.Println(nord4 + "  - Building and running Go projects (with new Terminal window on macOS)" + reset)
	fmt.Println(nord4 + "  - Running tests, formatting code, and tidying modules" + reset)
	fmt.Println(nord4 + "  - Running individual Go files" + reset)
	fmt.Println(nord4 + " Designed with the Nord color theme for a sleek development experience." + reset)
}