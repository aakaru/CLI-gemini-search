# CLI-Gemini-Search

## Overview
**CLI-Gemini-Search** is a Command Line Interface (CLI) tool built using Go, the Cobra CLI toolkit, and the Gemini API. This tool allows users to interact with the Gemini model directly from the command line to generate text-based responses from provided input queries.

## Features
- Built using Go and Cobra library for CLI structure and handling.
- Integrates with Google's **Gemini** model for text generation.
- Accepts user input from the command line and returns responses using the Gemini model.

## Installation

### Prerequisites
- Go (1.19 or later) installed on your machine.
- Cobra CLI toolkit installed: `go get -u github.com/spf13/cobra/cobra`
- A valid Gemini API key from Google.
  
### Clone the Repository
To get started, clone this repository using:
```bash
git clone https://github.com/yourusername/CLI-gemini-search.git
```
###Set up Environment Variables
You'll need to set your Gemini API key as an environment variable:
```bash
export GEMINI_API_KEY=your_api_key_here
```
###Build the Application
Navigate to the project directory and run:
```bash
go build
```

##Usage
Once you've successfully built the application, you can use it by running the following command:
```bash
./CLI-gemini-search search "your_query_here"
```
For example:
```bash
./CLI-gemini-search search "Tell me a joke"
```
##Project Structure
```bash
.
├── cmd
│   └── search.go      # Contains the core logic for the search command
├── main.go            # Entry point for the CLI application
├── go.mod             # Go module file
└── go.sum             # Go dependencies file
```
##Example
```bash
$ ./CLI-gemini-search search "Provide me a breif description of concurrency in Go in 100 words."
```
Output:
```bash
Concurrency in Go is achieved through goroutines, lightweight threads managed by the Go runtime. Goroutines run concurrently, allowing multiple tasks to execute seemingly simultaneously. Channels provide a safe and efficient way for goroutines to communicate and synchronize, preventing data races. This approach allows for efficient utilization of system resources and writing programs that handle multiple tasks gracefully. 
```
##Contributing
Contributions are welcome! If you'd like to improve this tool, feel free to submit a pull request or open an issue.

##License
This project is licensed under the MIT License. See the LICENSE file for more details.


Feel free to adjust the repository URL or any other details as needed!

