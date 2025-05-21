# Real-time Data Stream

## Overview

The **Real-time Data Stream** project is a WebSocket-based application that enables real-time communication. Users can send and receive messages instantly, benefiting from features like username support, message history, and a clean user interface. This project demonstrates the use of WebSockets for creating dynamic and interactive applications.

---

## Features

- **Real-time Messaging**: Instantly send and receive messages.
- **Username Support**: Attribute messages to users by allowing them to set a username.
- **Message History**: Display previously sent messages to new users upon connection.
- **User-friendly Interface**: A clean and intuitive UI for seamless interaction.
- **Notification System**: Visual notifications for new incoming messages.

---

## Technologies Used

- **Go**: Backend server built using Go and the Gorilla WebSocket package.
- **HTML/CSS/JavaScript**: Frontend interface created with standard web technologies.
- **Python**: For serving static files.

---

## Prerequisites

Ensure the following are installed:

- **Go** (1.16 or later)
- **Python** (for static file serving, if desired)
- **Git** (optional, for cloning the repository)

---

## Required Packages

Install the necessary Go packages to run the application:

1. Open a terminal and navigate to the project directory.
2. Run the following command:

   ```bash
   go get github.com/gorilla/websocket
   ```

---

## Setup

### Clone the Repository
Clone the project repository (if hosted on GitHub):

```bash
git clone https://github.com/yourusername/real-time-data-stream.git](https://github.com/joyo11/RealTimeDataPipeline
```

### Run the Backend Server
Start the Go WebSocket server:

```bash
go run main.go
```

### Serve the Frontend (Optional)
To serve static files using Python:

```bash
python3 -m http.server 8000
```

---

## How to Use

1. **Enter Your Username**:
   - Use the username input field to set your desired username.

2. **Send Messages**:
   - Type your message in the input box.
   - Click the "Send" button or press Enter to send your message.

3. **Real-time Updates**:
   - All sent messages appear instantly for all connected users.

---

## Example Interaction

```plaintext
[System]: Welcome to the Real-time Chat!
[User1]: Hello, everyone!
[User2]: Hi, User1! How are you?
```

- **New User**:
  Upon joining, the system displays the message history to the new user.

---

## Code Structure

### Backend (Go)
- **WebSocket Server**: Manages real-time communication between clients.
- **Message Handling**: Broadcasts messages to all connected users.

### Frontend (HTML/CSS/JavaScript)
- **User Interface**: Provides a clean and responsive design for users.
- **WebSocket Client**: Handles sending and receiving messages via WebSockets.

---

## Contact
For any questions or inquiries, feel free to reach out:
- **Email**: [shafay11august@gmail.com](mailto:shafay11august@gmail.com)

---

### Notes
This project is designed for educational purposes and demonstrates WebSocket implementation in Go. Future enhancements could include:
- Persistent chat storage using a database.
- Authentication for secure communication.
- Advanced UI features like typing indicators and emojis.

