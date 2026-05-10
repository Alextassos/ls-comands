package main

import (
	"fmt"
	"os"

	"myls/utils"
)

func main() {
	myFlags, myPaths := utils.ParseFlags()

	for _, path := range myPaths {
		// Χρησιμοποιούμε Lstat για να μην "πηδάει" τα links
		info, err := os.Lstat(path)
		if err != nil {
			fmt.Printf("my-ls: %s: No such file or directory\n", path)
			continue
		}

		// Αν είναι φάκελος, τρέξε τη ListDirectory
		if info.IsDir() {
			// Αν έχουμε πολλά paths, τύπωσε το όνομα του φακέλου (όπως το ls)
			if len(myPaths) > 1 {
				fmt.Printf("%s:\n", path)
			}
			utils.ListDirectory(path, myFlags)

			// Αν έχουμε κι άλλα paths, άφησε μια κενή γραμμή ενδιάμεσα
			if len(myPaths) > 1 {
				fmt.Println()
			}
		} else {
			// Αν είναι μεμονωμένο αρχείο
			if myFlags.LongFormat {
				utils.PrintLongFormat(info, path)
			} else {
				fmt.Println(path)
			}
		}
	}
}
