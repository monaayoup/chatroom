# 🗨️ Simple Chatroom (Go + RPC)

A lightweight chatroom implemented in **Go (Golang)** using **net/rpc** for remote procedure calls.  
This project demonstrates client-server communication, message persistence in memory, and simple Docker integration.

---

## 📘  Overview

> **: Simple Chatroom**

### Requirements
1. **Client**
   - Dials the RPC of the coordinating server.
   - Calls a remote procedure on the server to send messages.
   - Fetches the chat history using another remote procedure.

2. **Server**
   - Stores every message in a growing in-memory list.
   - Returns the full chat history whenever a client sends a new message.

### Notes
- The client runs continuously until you type `exit` or press `Ctrl+C`.
- Uses `bufio.Reader` instead of `fmt.Scan` to handle multi-word input.
- Handles reconnection if the server goes down.

---
working chatroom : https://drive.google.com/drive/folders/15oOfH59SfE66aes7NHgJxL-YOpVfWYCD?usp=drive_link



