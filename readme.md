# Real-time Data Stream

## Overview

The **Real-time Data Stream** project is a simple WebSocket-based application that allows users to send and receive messages in real-time. This project demonstrates the power of WebSockets for creating interactive applications, showcasing features like message history, user attribution, and a clean user interface.

## Features

- **Real-time Messaging**: Users can send and receive messages instantly.
- **Username Support**: Users can set a username, allowing messages to be attributed to them.
- **Message History**: Previous messages are displayed to new users when they connect.
- **User-friendly Interface**: A clean and attractive UI designed for easy interaction.
- **Notification System**: Visual notifications for incoming messages.

## Technologies Used

- **Go**: The backend server is built using Go and the Gorilla WebSocket package.
- **HTML/CSS/JavaScript**: The frontend interface is created using standard web technologies.

## Prerequisites

Make sure you have the following installed:

- **Go** (1.16 or later)
- **Python** (for serving static files, if desired)
- **Git** (optional, for cloning the repository)

## Required Packages

To ensure the application runs smoothly, install the necessary Go packages:

1. Open a terminal and navigate to the project directory.
2. Run the following command:

   ```bash
   go get github.com/gorilla/websocket
3.go run main.go
4.python3 -m http.server 8000


Usage
Enter your username in the input field.
Type your message in the message input box.
Click the "Send" button or press Enter to send your message.
All messages sent by users will be displayed in real-time.