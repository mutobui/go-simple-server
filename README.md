Sure, here’s a revised `README.md` focused on step-by-step instructions for testing the WebSocket server and client.

### README.md

```markdown
# Simple Server

This repository contains a simple WebSocket server implemented in Golang that generates unique big integers for each WebSocket connection. A JavaScript client is also provided to connect to the server and receive these unique big integers.

## Requirements

- Golang 1.20+
- Web browser (for running the JavaScript client)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/Simple-Server.git
cd Simple-Server
```

### 2. Install Dependencies

Ensure you have the Gorilla WebSocket package installed:

```bash
go get github.com/gorilla/websocket
```

### 3. Run the Server

```bash
go run server.go
```

The server will start on port 8080.

## Usage

### 1. Open the Client HTML File

Open the `websocket_test.html` file in a web browser.

### 2. Establish a WebSocket Connection

- Click the "Connect" button to establish a WebSocket connection.
- You should see a message indicating that the connection is established.

### 3. Send a Message

- Click the "Send Message" button to send a message to the server.
- The server will respond with a unique big integer, which will be displayed on the webpage.

### 4. Check the Browser Console

- Open the developer tools in your web browser (usually `F12` or `Ctrl+Shift+I`) and go to the console tab.
- You will see detailed logs of the connection and message events.

### 5. Verify Server Logs

- Check the server logs in your terminal to see messages indicating new connection attempts, established connections, received messages, sent unique numbers, and connection closures.

## Directory Structure

```plaintext
.
├── server.go
└── websocket_test.html
```

- `server.go`: The Go server code.
- `websocket_test.html`: The HTML and JavaScript client code.
```

This version of the `README.md` file provides clear, step-by-step instructions for testing the WebSocket server and client without including the code or license details. Good luck with your interview!# go-simple-server
