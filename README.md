# go-grep

A simple implementation of the `grep` command in Go, supporting basic pattern matching, recursive search, case-insensitive search, and more.

This project was originally created as part of a coding challenge from [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-grep).

## Features

- Match lines in a file or directory tree using a regular expression.
- Support for the following command-line options:
  - `-r`: Recursively search directories.
  - `-v`: Invert the match, excluding lines that match the pattern.
  - `-i`: Case-insensitive search.
- Special regex characters supported:
  - `^`: Matches the start of a line.
  - `$`: Matches the end of a line.
  - `\d`: Matches any digit.
  - `\w`: Matches any word character.

## Installation

To install the `wc` tool, you need to have Golang installed on your system. Follow these steps:

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
Usage: ./grep-clone [-r] [-v] [-i] <pattern> <filename or directory>
```

Perform a basic search:

```sh
./go-grep "pattern" filename.txt
```

Perform a recursive search:

```sh
./go-grep -r "pattern" filename.txt
```

Invert matched pattern:

```sh
./grep-clone -v "pattern" filename.txt
```

Perform a case-insensitive search:

```sh
./go-grep -i "pattern" filename.txt
```

You can combine any of these options:

```sh
./go-grep -r -i -v "pattern" filename.txt
```

### Special Regex Characters

- `^`: Matches the beginning of a line.
- `$`: Matches the end of a line.
- `\d`: Matches any digit.
- `\w`: Matches any word character.

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
