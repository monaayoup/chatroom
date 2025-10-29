package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
	"time"
)

const serverAddr = "localhost:1234"

// connect tries to dial the RPC server with backoff.
func connect() *rpc.Client {
    var client *rpc.Client
    var err error
    backoff := time.Second

    for {
        client, err = rpc.Dial("tcp", serverAddr)
        if err == nil {
            return client
        }
        fmt.Printf("Unable to connect to server (%v). Retrying in %s...\n", err, backoff)
        time.Sleep(backoff)
        if backoff < 10*time.Second {
            backoff *= 2
        }
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    nameRaw, _ := reader.ReadString('\n')
    name := strings.TrimSpace(nameRaw)
    if name == "" {
        name = "Anonymous"
    }

    fmt.Println("Connecting to chat server...")
    client := connect()
    defer func() {
        if client != nil {
            client.Close()
        }
    }()
    fmt.Println("Connected. Type messages (type 'exit' to quit).")

    for {
        fmt.Print("> ")
        textRaw, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Read error:", err)
            break
        }
        text := strings.TrimSpace(textRaw)
        if text == "" {
            continue
        }
        if text == "exit" {
            fmt.Println("Exiting.")
            break
        }

        msg := Message{
            Sender:    name,
            Content:   text,
            Timestamp: time.Now(),
        }

        var history []Message
        callErr := client.Call("ChatServer.SendMessage", msg, &history)
        if callErr != nil {
            // Try reconnecting and re-sending once
            fmt.Println("Error sending message:", callErr)
            fmt.Println("Attempting to reconnect...")
            client.Close()
            client = connect()
            fmt.Println("Reconnected. Re-sending message...")
            callErr = client.Call("ChatServer.SendMessage", msg, &history)
            if callErr != nil {
                fmt.Println("Failed again:", callErr)
                fmt.Println("Please try again later or press Ctrl+C to quit.")
                continue
            }
        }

        fmt.Println("\n--- Chat History ---")
        for _, m := range history {
            ts := m.Timestamp.Format("2006-01-02 15:04:05")
            fmt.Printf("[%s] %s: %s\n", ts, m.Sender, m.Content)
        }
        fmt.Println("--------------------\n")
    }
}