# Anonymous Chat

This project is a minimal anonymous chat server written in Go. The server allows users to connect via WebSocket and exchange messages in real time.

## Features

- Multiple user connections
- Anonymous message exchange
- Easy to use

## Technologies

- Go
- WebSocket (using the [Anonymous Chat](https://anonchat.com/) library)

## Installation

To use this project, you need to have Go installed. Then follow these steps:

1. Clone the repository:
   ```bash
   git clone <your-repository-URL>
   cd <repository-name>

2. Install dependencies:
```bash
go get -u github.com/gorilla/websocket

3. Run the server:
```bash
go run main.go
