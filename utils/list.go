package utils

import (
	"fmt"
	"os"
	"syscall"
)

type Flags struct {
	LongFormat bool
	Recursive  bool
	ShowAll    bool
	Reverse    bool
	SortByTime bool
}

// Αντικαθιστά το filepath.Join
func joinPath(dir, file string) string {
	if len(dir) == 0 {
		return file
	}
	// Αν ο φάκελος δεν τελειώνει ήδη σε /, το προσθέτουμε
	if dir[len(dir)-1] == '/' {
		return dir + file
	}
	return dir + "/" + file
}

// Δικός μας Bubble Sort για να αντικαταστήσει το package sort
func sortEntries(entries []os.DirEntry, sortByTime bool) {
	n := len(entries)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			shouldSwap := false

			if sortByTime {
				infoI, _ := entries[j].Info()
				infoJ, _ := entries[j+1].Info()
				// Αν ο χρόνος του J είναι μεταγενέστερος του I, αλλάζουμε (νεότερα πρώτα)
				if infoI.ModTime().Before(infoJ.ModTime()) {
					shouldSwap = true
				}
			} else {
				// Αλφαβητική ταξινόμηση (A-Z)
				if entries[j].Name() > entries[j+1].Name() {
					shouldSwap = true
				}
			}

			if shouldSwap {
				entries[j], entries[j+1] = entries[j+1], entries[j]
			}
		}
	}
}

func ListDirectory(path string, f Flags) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("my-ls: %s: No such file or directory\n", path)
		return
	}

	var filtered []os.DirEntry

	for _, entry := range entries {
		name := entry.Name()

		if !f.ShowAll && len(name) > 0 && name[0] == '.' {
			continue
		}
		filtered = append(filtered, entry)
	}

	// 2. Ταξινόμηση
	sortEntries(filtered, f.SortByTime)

	// 3. Αντιστροφή (-r)
	if f.Reverse {
		for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
			filtered[i], filtered[j] = filtered[j], filtered[i]
		}
	}

	// 4. Εμφάνιση του Path στην αναδρομή
	if f.Recursive {
	}

	// 5. Υπολογισμός TOTAL (μόνο αν έχει -l)
	if f.LongFormat && (len(filtered) > 0 || f.ShowAll) {
		var total int64

		for _, entry := range filtered {
			info, _ := entry.Info()
			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				total += stat.Blocks / 2
			}
		}

		if f.ShowAll {
			dot, _ := os.Stat(path)
			dotDot, _ := os.Stat(joinPath(path, ".."))
			if s, ok := dot.Sys().(*syscall.Stat_t); ok {
				total += s.Blocks / 2
			}
			if s, ok := dotDot.Sys().(*syscall.Stat_t); ok {
				total += s.Blocks / 2
			}
		}
		fmt.Printf("total %d\n", total)
	}

	// 6. ΕΚΤΥΠΩΣΗ

	printSpecial := func() {
		if f.ShowAll {
			dot, _ := os.Stat(path)
			dotDot, _ := os.Stat(joinPath(path, ".."))
			if f.LongFormat {
				PrintLongFormat(dot, ".")
				PrintLongFormat(dotDot, "..")
			} else {
				fmt.Print(".  ..  ")
			}
		}
	}

	if !f.Reverse {
		printSpecial()
	}

	for _, entry := range filtered {
		if f.LongFormat {
			info, _ := entry.Info()
			PrintLongFormat(info, entry.Name())
		} else {
			fmt.Print(entry.Name(), "  ")
		}
	}

	if f.Reverse {
		printSpecial()
	}

	if !f.LongFormat {
		fmt.Println()
	}

	// 7. ΑΝΑΔΡΟΜΗ (-R)
	if f.Recursive {
		for _, entry := range filtered {
			if entry.IsDir() {
				name := entry.Name()
				if name == "." || name == ".." {
					continue
				}
				newPath := joinPath(path, name)
				fmt.Printf("\n%s:\n", newPath)
				ListDirectory(newPath, f)
			}
		}
	}
}
