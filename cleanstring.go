package cleanstring

import (
	"bufio"
	"strings"
	"unicode"
)

// Given a string input, this method will:
// - Remove any leading blank/whitespace-only lines.
// - Strip a prefix consisting of any amount of whitepsace followed by a pipe ("|") character.
// - Remove any trailing blank/whitespace-only lines.
func Get(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	output := strings.Builder{}
	hasSeenNonWhitespaceLine := false
	lastLineWasNonWhitespace := false
	startOfTrailingWhitespace := -1
	for scanner.Scan() {
		line := scanner.Text()
		isLineWhitespace, prefixLength := parseLine(line)
		if isLineWhitespace {
			if !hasSeenNonWhitespaceLine {
				// Do not write prefix lines which are only whitespace.
			} else {
				if lastLineWasNonWhitespace {
					// This might be the start of trailing whitespace lines.
					// Minus one because we would have written an additional newline.
					startOfTrailingWhitespace = output.Len() - 1
				}
				// Write the line as output.
				output.WriteString(line)
				// Add an extra newline since Scan strips newlines.  We'll strip the trailing newline.
				output.WriteString("\n")
			}
			lastLineWasNonWhitespace = false
		} else {
			hasSeenNonWhitespaceLine = true
			// Strip prefix if it exists.
			if prefixLength > 0 {
				line = line[prefixLength:]
			}
			// Write every non-whitespace line.
			output.WriteString(line)
			// Add an extra newline since Scan strips newlines.  We'll strip the trailing newline.
			output.WriteString("\n")
			lastLineWasNonWhitespace = true
			startOfTrailingWhitespace = -1
		}
	}
	// There are trailing lines we should strip.
	if startOfTrailingWhitespace != -1 {
		return output.String()[0:startOfTrailingWhitespace]
	}
	// Otherwise return everything minus the last newline character.
	return output.String()[0:(output.Len() - 1)]
}

func parseLine(line string) (isWhitespaceOnly bool, prefixLength int) {
	prefixLength = 0
	isWhitespaceOnly = true
	for _, c := range line {
		if unicode.IsSpace(c) {
			if isWhitespaceOnly {
				prefixLength += 1
			}
		} else {
			if c == '|' && isWhitespaceOnly {
				prefixLength += 1
			}
			isWhitespaceOnly = false
			break
		}
	}
	return isWhitespaceOnly, prefixLength
}