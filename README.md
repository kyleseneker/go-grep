# go-grep

`go-grep` is a command-line tool written in Go for searching patterns in files and directories. It supports various options to customize the search behavior, similar to the UNIX `grep` command-line utility.

This project was originally created as  part of a coding challenge from [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-grep).

## Features

- Search for a regular expression pattern in one or more files or directories.
- Recursive search through directories (`-r`).
- Invert the match to exclude lines that match the pattern (`-v`).
- Ignore case distinctions in the pattern (`-i`).
- Print only a count of matching lines per file (`-c`).
- Prefix each line of output with the 1-based line number within its input file (`-n`).
- Suppress normal output, useful for scripts (`-q`).

## Installation

To install the `go-grep` tool, you need to have Golang installed on your system. Follow these steps:

1. Clone the repository:

    ```sh
    git clone https://github.com/kyleseneker/go-grep.git
    ```

1. Navigate to the project directory:

    ```sh
    cd go-grep
    ```

1. Build the program:

    ```sh
    make build
    ```

## Usage

```sh
go-grep [-r] [-v] [-i] [-c] [-n] [-q] <pattern> <filename or directory>
```

### Options

- `-r`, `--recursive`: Recursively search through directories.
- `-v`, `--invert`: Invert the match to exclude matching lines.
- `-i`,`--ignore-case`: Ignore case distinctions in the pattern.
- `-c`, `--count`: Print only a count of matching lines per file.
- `-n`, `--line-number`: Prefix each line of output with the 1-based line number within its input file.
- `-q`, `--quiet`, `--silent`: Suppress normal output, useful for scripts.

### Special Regex Characters

- `^`: Matches the beginning of a line.
- `$`: Matches the end of a line.
- `\d`: Matches any digit.
- `\w`: Matches any word character.

## Examples

Perform a basic search:

```sh
$ ./go-grep J examples/rockbands.txt
Judas Priest
Bon Jovi
Junkyard
```

Perform a recursive search:

```sh
$ ./go-grep -r Nirvana *
examples/rockbands.txt:Nirvana
examples/test-subdir/BFS1985.txt:Since Bruce Springsteen, Madonna, way before Nirvana
examples/test-subdir/BFS1985.txt:On the radio was Springsteen, Madonna, way before Nirvana
examples/test-subdir/BFS1985.txt:And bring back Springsteen, Madonna, way before Nirvana
examples/test-subdir/BFS1985.txt:Bruce Springsteen, Madonna, way before Nirvana
```

Invert matched pattern:

```sh
$ ./grep-clone -r Nirvana * | grep -v Madonna
examples/rockbands.txt:Nirvana
```

Perform a case-insensitive search:

```sh
$ ./go-grep -i A examples/rockbands.txt | wc -l
      58
```

Perform pattern matching for searching only the beginning or end of a line:

```sh
$ ./go-grep ^A examples/rockbands.txt
AC/DC
Aerosmith
Accept
April Wine
Autograph
```

```sh
$ ./go-grep na$ examples/rockbands.txt
Nirvana
```

## Output Formatting

- Matching parts of lines are highlighted in red and bold by default, similar to how grep highlights matches in terminals.
- Output includes the file name if `-r` (recursive) option is used.

## Exit Codes

The `go-grep` utility exits with one of the following values:

- `0`: One or more lines were selected.
- `1`: No lines were selected.
- `2`: An error occurred.

## Performance

[hyperfine](https://github.com/sharkdp/hyperfine) is used to perform benchmarks.

To run the pre-defined benchmark:

```sh
make benchmark
```

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
