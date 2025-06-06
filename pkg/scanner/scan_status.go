package scanner

import (
	"fmt"
	"sync"
)

type ScanStatus struct {
	mu    sync.Mutex
	Lines map[string]int // Platform -> line number
}

func (ss *ScanStatus) LineInit(line int, platform, msg string) {
	ss.Lines[platform] = line
	fmt.Println(msg)
}

func (ss *ScanStatus) Set(platform, msg string) {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	line := ss.Lines[platform]
	total_lines := len(ss.Lines)

	moveCursorUp(total_lines - line)
	clearLine()
	fmt.Print(msg)
	moveCursorDown(total_lines - line)
}

// Move cursor up relative to current line
func moveCursorUp(n int) {
	if n > 0 {
		fmt.Printf("\033[%dA", n)
	}
}

// Move cursor down relative to current line
func moveCursorDown(n int) {
	if n > 0 {
		fmt.Printf("\033[%dB", n)
	}
}

// Clear current line
func clearLine() {
	fmt.Print("\r\033[2K")
}
