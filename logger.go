package overlay

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// debug prints its input to a debug file, but only when a DEBUG environment variable has been set.
func debug(lines ...string) {
	ts := strconv.Itoa(int(time.Now().UnixMilli())) + "\n"
	s := strings.Join(lines, "\n")

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", ts+s)
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
}