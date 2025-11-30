# keylogger

# This repository contains a demo keylogger project written in Go, built strictly for educational and testing purposes.
# It includes:

# A Go keylogger client that captures keystrokes on the machine it runs on

# A simple web server that receives and displays the logged keystrokes

# A demo warning page (instead of consent flow) that clearly states the project is for demonstration only



# Project Structure

# /cmd
#  /client        → keylogger client
#  /server        → HTTP server that collects & shows logs
# go.mod
# README.md


# Features

# Captures keystrokes using Go
# Sends logs to your server endpoint
# Displays logs live on a web page
# Includes a visible demo warning message on the webpage
# Lightweight and minimal—easy to learn from


# Installation

# Clone the repository:

# git clone https://github.com/Mustafa-nq/keylogger.git
# cd keylogger


# Initialize Go modules (if not already done):

# go mod tidy

# ▶ Running the Server

# Start the HTTP server first:

# cd cmd/server
# go run .


# The server:

# Listens for keystroke logs
# Displays them on a webpage
# Shows a demo keylogger warning message

# Open in browser:

# http://localhost:8080

# ⌨ Running the Keylogger Client

# In another terminal:

# cd cmd/client
# go run .


# The client will:

# Start capturing keystrokes
# Send them to the server’s /log endpoint

# WEB UI

# The webpage automatically displays:

# Logged keystrokes
# A red banner stating:
# “⚠️ This is a demo keylogger project — for educational use only.”

# CUSTOMIZING

# You can modify:

# The endpoint URL
# The UI layout
# Log storage mechanism
# Client behavior (buffering, intervals, etc.)