package utils

import (
	"fmt"
	"os"
)

func ParseFlags() (Flags, []string) {
	var f Flags
	var paths []string

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		if len(arg) > 1 && arg[0] == '-' {
			for _, char := range arg[1:] {
				switch char {
				case 'l':
					f.LongFormat = true
				case 'a':
					f.ShowAll = true
				case 'R':
					f.Recursive = true
				case 'r':
					f.Reverse = true
				case 't':
					f.SortByTime = true
				default:
					fmt.Printf("my-ls: invalid option -- '%c'\n", char)
					os.Exit(1)
				}
			}
		} else {
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 {
		paths = append(paths, ".")
	}
	return f, paths
}
