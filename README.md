# my-ls

A custom implementation of the Unix `ls` command written in **Go**. This project replicates the behavior of the original system command, supporting various flags, file metadata retrieval, and recursive directory listing.

## Features (Supported Flags)

The `my-ls` utility incorporates the following essential flags:

* `-l`: **Long format** (displays file permissions, number of links, owner, group, size, and last modification time).
* `-a`: **All** (includes hidden files starting with `.` as well as the special `.` and `..` directories).
* `-R`: **Recursive** (lists all subdirectories and their contents recursively).
* `-r`: **Reverse** (reverses the order of the sorting).
* `-t`: **Time sort** (sorts entries by modification time, newest first).

## Installation & Requirements

To run this project, you need to have [Go](https://go.dev) installed on your system.

1. **Clone the repository:**
   ```bash
   git clone <your-repo-link>
   cd my-ls-1
   ```

2. **Run without building:**
   ```bash
   go run . [flags] [directory]
   ```

3. **Build the executable:**
   ```bash
   go build -o my-ls
   ./my-ls -la
   ```

## Usage Examples

- List the current directory in long format:
  `./my-ls -l`
- List all files (including hidden) recursively:
  `./my-ls -laR`
- Sort by time in reverse order:
  `./my-ls -tr`
- List a specific directory:
  `./my-ls -l /home/user/Documents`

## Unit Testing

The project includes unit tests to ensure the reliability of core logic such as path joining and flag parsing.

To execute the tests, run:
```bash
go test -v ./utils
```

## Technical Highlights
- **System Calls**: Uses the `syscall` package to retrieve low-level file metadata (blocks, UID, GID).
- **Custom Sorting**: Implements a sorting algorithm to handle alphabetical and time-based ordering.
- **Modularity**: Organized into packages (`main` and `utils`) for better maintainability and clean code practices.
- **No os/exec**: Built strictly using Go standard library packages (prohibiting the use of external shell commands).


