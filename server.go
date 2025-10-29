package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc"
    "sync"
    "time"
)

// ChatServer stores chat messages and exposes RPC methods.
// Methods must be exported.
type ChatServer struct {
    mu       sync.Mutex
    messages []Message
}

// SendMessage appends the incoming message and returns the full history.
func (s *ChatServer) SendMessage(msg Message, reply *[]Message) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    // ensure timestamp if client didn't set
    if msg.Timestamp.IsZero() {
        msg.Timestamp = time.Now()
    }

    s.messages = append(s.messages, msg)

    // return a copy to the client
    history := make([]Message, len(s.messages))
    copy(history, s.messages)
    *reply = history
    return nil
}

func main() {
    server := new(ChatServer)
    // register with net/rpc
    err := rpc.Register(server)
    if err != nil {
        log.Fatalf("rpc.Register error: %v", err)
    }

    ln, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatalf("Listen error: %v", err)
    }
    fmt.Println("✅ Chat server listening on :1234")

    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Println("Accept error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
