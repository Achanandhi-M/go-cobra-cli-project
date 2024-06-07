/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
    Use:   "random",
    Short: "Get a random quote",
    Long: `This command fetches a random quote from the quotable.io API
and displays it in the console.`,
    Run: func(cmd *cobra.Command, args []string) {
        getRandomQuote()
    },
}

func init() {
    rootCmd.AddCommand(randomCmd)
}

type Quote struct {
    ID      string `json:"_id"`
    Content string `json:"content"`
    Author  string `json:"author"`
}

func getRandomQuote() {
	fmt.Println("Fetching a random quote 😊")
    url := os.Getenv("QUOTES_API_URL")
    if url == "" {
        fmt.Println("Error: QUOTES_API_URL environment variable is not set")
        return
    }

    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Printf("Error creating request: %v\n", err)
        return
    }

    req.Header.Set("Accept", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("Error fetching quote: %v\n", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error: received status code %d\n", resp.StatusCode)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %v\n", err)
        return
    }

    var quote Quote
    err = json.Unmarshal(body, &quote)
    if err != nil {
        fmt.Printf("Error unmarshaling JSON: %v\n", err)
        return
    }

	fmt.Printf("Quote ✨ \"%s\"\nBy🤵: %s\n", quote.Content, quote.Author)

}
