`ccwc` is a command-line utility written in Go for analyzing text files. It calculates the number of bytes, lines, words, and
characters in a file or from standard input. Inspired by Unix tools like `wc`, it provides flexible options for text analysis.
Features
- Count **bytes**, **lines**, **words**, or **characters** in a file.
- Supports input via **file paths** or **piped input**.
- Simple and intuitive flags for specific operations.
Installation
1. Clone the repository:
 ```bash
 git clone https://github.com/your-username/ccwc.git
 cd ccwc
 ```
1. Build the binary:
 ```bash
 go build -o ccwc
 ```
1. (Optional) Add it to your `PATH` for global use:
 ```bash
ccwc CLI Tool - README
 mv ccwc /usr/local/bin/
 ```
Usage
### Analyze a File
- Count **bytes**:
 ```bash
 ccwc -c filename.txt
 ```
- Count **lines**:
 ```bash
 ccwc -l filename.txt
 ```
- Count **words**:
 ```bash
 ccwc -w filename.txt
 ```
- Count **characters**:
 ```bash
 ccwc -m filename.txt
 ```
ccwc CLI Tool - README
### Analyze Piped Input
- Count **lines** in piped input:
 ```bash
 cat filename.txt | ccwc -l
 ```
- Count **words** in piped input:
 ```bash
 echo "This is a test" | ccwc -w
 ```
### Default Output
If no flags are specified, the default output includes all metrics:
```bash
ccwc filename.txt
```
Example output:
```text
342190 7145 5200
```
(Output format: `bytes lines words`)
Flags
ccwc CLI Tool - README
| Flag | Shorthand | Description |
|-----------|-----------|-----------------------------------|
| `--bytes` | `-c` | Number of bytes in the file |
| `--lines` | `-l` | Number of lines in the file |
| `--words` | `-w` | Number of words in the file |
| `--chars` | `-m` | Number of characters in the file |
Examples
- Count lines in a file:
 ```bash
 ccwc -l sample.txt
 ```
- Count words from a piped input:
 ```bash
 echo "Hello world!" | ccwc -w
 ```