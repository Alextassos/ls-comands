package utils

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func PrintLongFormat(info os.FileInfo, name string) {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return
	}

	mode := info.Mode().String()

	// Owner
	userName := strconv.FormatUint(uint64(stat.Uid), 10)
	if u, err := user.LookupId(userName); err == nil {
		userName = u.Username
	}

	// Group
	groupName := strconv.FormatUint(uint64(stat.Gid), 10)
	if g, err := user.LookupGroupId(groupName); err == nil {
		groupName = g.Name
	}

	modTime := info.ModTime().Format("Jan _2 15:04")

	fmt.Printf("%s %3d %-8s %-8s %8d %s %s\n",
		mode, stat.Nlink, userName, groupName, info.Size(), modTime, name)
}
