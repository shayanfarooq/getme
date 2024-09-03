package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
)

const (
	ContextLength = 200 // Number of characters before and after the keyword
	WorkerCount   = 10  // Number of concurrent threads

	Reset        = "\033[0m"
	Yellow       = "\033[33m"
	Red          = "\033[31m"
	Cyan         = "\033[36m"
	Magenta      = "\033[35m"
)

// Extracts snippets around the keyword from the content
func extractSnippets(content string, keyword string) {
	keywordPattern := regexp.MustCompile(regexp.QuoteMeta(keyword))
	matches := keywordPattern.FindAllStringIndex(content, -1)

	for _, match := range matches {
		start := match[0] - ContextLength
		if start < 0 {
			start = 0
		}
		end := match[1] + ContextLength
		if end > len(content) {
			end = len(content)
		}

		snippet := content[start:end]
		coloredSnippet := highlightKeyword(snippet, keyword)
		fmt.Printf("%sSnippet:%s\n%s\n\n", Magenta, Reset, coloredSnippet)
	}
}

// Highlights the keyword in the snippet
func highlightKeyword(snippet string, keyword string) string {
	keywordPattern := regexp.MustCompile(regexp.QuoteMeta(keyword))
	return keywordPattern.ReplaceAllString(snippet, Red+"$0"+Reset)
}

// Fetches the content from the given URL
func fetchContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

// Worker function to process each URL
func worker(urls <-chan string, keyword string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		content, err := fetchContent(url)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("%s[URL] %s%s\n", Cyan, url, Reset)
		extractSnippets(content, keyword)
	}
}

// Main function to handle command-line arguments and process URL(s)
func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: getme.go <url> [keyword]")
		fmt.Println("Or: echo <urls> | getme.go <keyword>")
		return
	}

	keyword := os.Args[len(os.Args)-1]
	var urls []string

	if len(os.Args) == 3 {
		urls = append(urls, os.Args[1])
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	urlChan := make(chan string, WorkerCount)
	var wg sync.WaitGroup

	// Start worker pool
	for i := 0; i < WorkerCount; i++ {
		wg.Add(1)
		go worker(urlChan, keyword, &wg)
	}

	// Feed URLs to workers
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)

	wg.Wait()
}
