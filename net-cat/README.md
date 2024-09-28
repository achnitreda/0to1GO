# TCP-Chat Application
## Overview

This project is a TCP-based chat server that allows multiple clients to connect, communicate, and exchange messages in real-time. The server handles up to two clients concurrently and manages messaging, user registration, message history, and broadcasting messages to all connected clients.

## Features
- Real-time Chat: Clients can send messages to one another via the server.
- Username Prompting: Users are prompted to enter a unique username upon connection.
- Welcome Message: Upon joining, users receive a welcome message with an ASCII banner.
- Message Broadcast: Messages are broadcasted to all connected users except the sender.
- Message History: New clients receive the full chat history upon joining.
- Client Management: The server manages a maximum of two clients simultaneously.
- Graceful Exit: Clients leaving the chat are announced to all others.

## Usage
Clone the repository:

```bash
git clone https://learn.zone01oujda.ma/git/mfir/net-cat.git
```

```bash
cd net-cat
Run the server:
```

```bash
go run .
```
Clients can connect to the server by specifying the IP and port.


## Code Explanation
- HandleConnection: Manages the connection lifecycle for each client. It handles username registration, message processing, and broadcasting. Mutex locks are used to ensure thread safety for shared resources like the clients map and messageHistory.

- broadcast: Sends messages to all connected clients except the sender. Mutex is locked during broadcasting to avoid concurrent access issues.

- printable: Helper function that ensures only printable ASCII characters are accepted in messages.

### Authors
- Mohamed Fri
- ABDELFATTAH ELIDRISSI
- REDA ACHNIT