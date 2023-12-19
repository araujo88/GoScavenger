# GoScavenger

This repository contains a simple web scraper written in Go. The scraper is designed to connect to a specified website using HTTPS, retrieve HTML content, and extract specific data based on HTML class names.

## Features

- Connects to websites using HTTPS.
- Reads and processes HTTP response headers.
- Handles both fixed `Content-Length` and `Transfer-Encoding: chunked` responses.
- Extracts content from HTML based on class names, ID and HTML tags.

## Files in the Repository

- `main.go`: Contains the main function that drives the web scraping process.
- `scraper.go`: Includes the `FindStringInTag`, `FindContentByID` and `FindContentByClass` functions, which are used for parsing HTML and extracting content.

## Getting Started

To use this scraper, you need to have Go installed on your machine. [Download and install Go](https://golang.org/dl/) if you haven't already.

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/araujo88/GoScavenger.git
cd GoScavenger
```

### Usage

1. Open `main.go`.
2. Modify the `server` variable to specify the website you want to scrape.
3. Optionally, adjust the request headers according to your requirements.
4. Run the scraper:

```bash
go run .
```

The output will be printed to the console.

## Contributing

Contributions to improve this simple web scraper are welcome. Feel free to fork the repository and submit pull requests.

## License

This project is licensed under the GPL License - see the [LICENSE](LICENSE) file for details.
