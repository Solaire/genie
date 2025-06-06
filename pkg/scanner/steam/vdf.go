package steam

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type VdfNode map[string]interface{}

func parseVdf(r io.Reader) (VdfNode, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var (
		stack []VdfNode
		keys  []string
	)

	root := VdfNode{}
	current := root

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		switch line {
		case "{":
			child := VdfNode{}
			last_key := keys[len(keys)-1]
			current[last_key] = child
			stack = append(stack, current)
			current = child
			keys = keys[:len(keys)-1]
		case "}":
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected }")
			}
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			parts := parseLine(line)
			if len(parts) == 2 {
				key := parts[0]
				val := parts[1]
				current[key] = val
			} else if len(parts) == 1 {
				keys = append(keys, parts[0])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return root, nil
}

func parseLine(line string) []string {
	var result []string
	for {
		start := strings.Index(line, "\"")
		if start == -1 {
			break
		}
		end := strings.Index(line[start+1:], "\"")
		if end == -1 {
			break
		}

		result = append(result, line[start+1:start+1+end])
		line = line[start+1+end+1:]
	}
	return result
}
