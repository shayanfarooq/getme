# URL Keyword Extractor

A Go command-line tool to search for keywords within content fetched from URLs and display snippets around the keyword with highlights.

## Overview

`getme.go` is a simple utility that fetches the content of one or more URLs and searches for a specified keyword within that content. It then extracts and highlights snippets around the keyword, making it easy to locate occurrences in the fetched content.

## Features

- Fetches content from a list of URLs.
- Searches for a keyword within the content.
- Displays snippets with the keyword highlighted.
- Supports concurrent processing with worker threads.

## Usage

### Command-Line Arguments

You can use `getme.go` in two ways:

1. **Single URL Mode:**

   ```
   go run getme.go <url> <keyword>
This fetches the content from the specified URL and searches for the keyword.

Multiple URLs Mode:

You can also use getme.go to process a list of URLs provided via standard input:

sh
Copy code
echo <urls> | go run getme.go <keyword>
Here, <urls> should be a list of URLs (one per line), and <keyword> is the keyword to search for.

Example
```
echo "https://example.com/page1" > urls.txt
echo "https://example.com/page2" >> urls.txt
cat urls.txt | go run getme.go "searchTerm"
```
Configuration
ContextLength: Number of characters before and after the keyword to include in the snippet. (Default: 200)
WorkerCount: Number of concurrent threads to use for fetching URLs. (Default: 10)
You can modify these constants directly in the code if needed.

Colors
The snippets are highlighted using ANSI escape codes:

Reset: Resets color formatting.
Yellow: Used for keywords (optional, not used in the provided code).
Red: Highlights the keyword.
Cyan: Indicates the URL being processed.
Magenta: Marks the beginning of a snippet.
Dependencies
This tool relies on standard Go libraries:

bufio
fmt
io/ioutil
net/http
os
regexp
sync
Building and Running
To build and run the program, ensure you have Go installed, then:

Clone this repository:

```git clone https://github.com/<username>/url-keyword-extractor.git```
Navigate to the project directory:

cd url-keyword-extractor
Build and run:

```go build -o getme getme.go```
./getme <url> <keyword>
Or run directly without building:

go run getme.go <url> <keyword>
License
This project is licensed under the MIT License. See the LICENSE file for details.

Contributing
Feel free to submit issues or pull requests. Contributions and suggestions are welcome!

Replace `<username>` with your GitHub username and adjust other placeholder details as needed. This README file provides a clear and concise overview of the tool, its usage, and setup instructions.
