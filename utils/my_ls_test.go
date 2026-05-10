package utils

import (
	"os"
	"testing"
)

// TestJoinPath ελέγχει αν η ένωση φακέλων και αρχείων γίνεται σωστά
func TestJoinPath(t *testing.T) {
	tests := []struct {
		dir      string
		file     string
		expected string
	}{
		{"folder", "file.txt", "folder/file.txt"},
		{"folder/", "file.txt", "folder/file.txt"},
		{"", "file.txt", "file.txt"},
	}

	for _, tt := range tests {
		result := joinPath(tt.dir, tt.file)
		if result != tt.expected {
			t.Errorf("joinPath(%s, %s) = %s; want %s", tt.dir, tt.file, result, tt.expected)
		}
	}
}

// TestParseFlags ελέγχει αν το πρόγραμμα καταλαβαίνει σωστά τα flags από το τερματικό
func TestParseFlags(t *testing.T) {
	// Προσομοιώνουμε τα ορίσματα του τερματικού
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }() // Επαναφορά στο τέλος

	os.Args = []string{"cmd", "-la", "my_folder"}
	flags, paths := ParseFlags()

	if !flags.LongFormat || !flags.ShowAll {
		t.Errorf("Expected flags -l and -a to be true")
	}

	if len(paths) != 1 || paths[0] != "my_folder" {
		t.Errorf("Expected path 'my_folder', got %v", paths)
	}
}
