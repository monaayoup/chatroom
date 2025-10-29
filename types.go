package main

import "time"

// Message is the structure sent between client and server.
// Fields must be exported (capitalized) so net/rpc can encode them.
type Message struct {
    Sender    string
    Content   string
    Timestamp time.Time
}
