package console

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const BANNER string = `
                                 | |            
  __ _  ___ _ __ ___   __ _ _ __ | |_ _ __ __ _ 
 / _' |/ _ \ '_ ' _ \ / _' | '_ \| __| '__/ _' |
| (_| |  __/ | | | | | (_| | | | | |_| | | (_| |
 \__, |\___|_| |_| |_|\__,_|_| |_|\__|_|  \__,_|
    | |                                         
    |_|
`

func ShowBanner() {
	fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb86c")).Render(BANNER))
	fmt.Println()
}
