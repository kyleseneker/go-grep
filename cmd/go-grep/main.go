package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	colorRed   = "\033[31m"
	colorBold  = "\033[1m"
	colorReset = "\033[0m"
)

func main() {
	// Initialize CLI application
	app := &cli.App{
		Name:  "go-grep",
		Usage: "search for patterns in files",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "r",
				Aliases: []string{"recursive"},
				Usage:   "recurse through directories",
			},
			&cli.BoolFlag{
				Name:    "v",
				Aliases: []string{"invert"},
				Usage:   "invert match",
			},
			&cli.BoolFlag{
				Name:    "i",
				Aliases: []string{"ignore-case"},
				Usage:   "ignore case distinctions",
			},
			&cli.BoolFlag{
				Name:    "c",
				Aliases: []string{"count"},
				Usage:   "print only a count of matching lines per FILE",
			},
			&cli.BoolFlag{
				Name:    "n",
				Aliases: []string{"line-number"},
				Usage:   "prefix each line of output with the 1-based line number within its input file",
			},
			&cli.BoolFlag{
				Name:    "q",
				Aliases: []string{"quiet", "silent"},
				Usage:   "suppress normal output",
			},
		},
		Action: func(c *cli.Context) error {
			// Check if enough arguments are provided
			if c.NArg() < 2 {
				fmt.Println("Usage: go-grep [-r] [-v] [-i] [-c] [-n] [-q] <pattern> <filename or directory>")
				return cli.Exit("", 2)
			}

			// Extract pattern and path(s) from command-line arguments
			pattern := c.Args().First()
			paths := c.Args().Slice()[1:]

			// Parse CLI flags
			recursive := c.Bool("r")
			invert := c.Bool("v")
			caseInsensitive := c.Bool("i")
			count := c.Bool("c")
			lineNumber := c.Bool("n")
			quiet := c.Bool("q")

			var re *regexp.Regexp
			var err error

			// Compile regex with case insensitivity if -i is specified
			if caseInsensitive {
				re, err = regexp.Compile("(?i)" + pattern)
			} else {
				re, err = regexp.Compile(pattern)
			}

			// Handle regex compilation errors
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error compiling regex pattern %s: %v\n", pattern, err)
				return cli.Exit("", 2)
			}

			matched := false

			// Iterate over each path provided
			for _, path := range paths {
				if recursive {
					// Perform recursive search through directories
					err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
						if err != nil {
							fmt.Fprintf(os.Stderr, "Error accessing path %s: %v\n", path, err)
							return err
						}
						if !info.IsDir() {
							// Search each file for matches
							if searchFile(path, re, invert, count, lineNumber, quiet, recursive) {
								matched = true
								if quiet {
									return filepath.SkipDir // Skip remaining files in this directory
								}
							}
						}
						return nil
					})
					// Handle errors during recursive directory traversal
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error walking the path %s: %v\n", path, err)
						return cli.Exit("", 2)
					}
				} else {
					// Perform search in a single file or directory without recursion
					if searchFile(path, re, invert, count, lineNumber, quiet, recursive) {
						matched = true
						if quiet {
							return cli.Exit("", 0)
						}
					}
				}
			}

			// Exit with appropriate status code based on whether matches were found
			if matched {
				return cli.Exit("", 0)
			} else {
				return cli.Exit("", 1)
			}
		},
	}

	// Run the CLI application
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running app: %v\n", err)
	}
}

// searchFile searches for the regex pattern in the specified file
// and prints matching lines. Returns true if matches were found.
func searchFile(fileName string, re *regexp.Regexp, invert, count, lineNumber, quiet, recursive bool) bool {
	// Access file information
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error accessing %s: %v\n", fileName, err)
		return false
	}

	// Skip directories, as we only want to search files
	if fileInfo.IsDir() {
		return false
	}

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", fileName, err)
		return false
	}
	defer file.Close()

	// Initialize variables for line number and match count
	matched := false
	reader := bufio.NewReader(file)
	lineNum := 0
	matchCount := 0

	// Read each line from the file
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", fileName, err)
			}
			break
		}
		lineNum++
		line = strings.TrimRight(line, "\n") // Remove trailing newline
		match := re.MatchString(line)
		// Process matching lines based on flags
		if (match && !invert) || (!match && invert) {
			matched = true
			if quiet {
				return true
			}
			if count {
				matchCount++
			} else {
				if lineNumber {
					if recursive {
						fmt.Printf("%s:%d:", fileName, lineNum)
					} else {
						fmt.Printf("%d:", lineNum)
					}
				} else {
					if recursive {
						fmt.Printf("%s:", fileName)
					}
				}
				// Highlight matching parts of the line
				highlightedLine := re.ReplaceAllStringFunc(line, func(match string) string {
					return colorBold + colorRed + match + colorReset
				})
				fmt.Println(highlightedLine)
			}
		}
	}

	// Print match count if -c flag is set and matches were found
	if count && matchCount > 0 {
		fmt.Printf("%s:%d\n", fileName, matchCount)
	}

	return matched
}
