package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func main() {
	htmlContent := ""
	server := "www.globo.com"
	// tag := "p"

	conf := &tls.Config{
		//InsecureSkipVerify: true,
	}

	// Establish a TLS connection
	conn, err := tls.Dial("tcp", server+":443", conf)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(time.Second * 5)) // Set a 5-second timeout

	request := "GET / HTTP/1.1\r\n"
	host := "Host: " + server + "\r\n"
	userAgent := "User-agent: curl/7.81.0\r\n"
	accept := "Accept: */*\r\n"

	fmt.Fprintf(conn, request+host+userAgent+accept+"\r\n\r\n")

	reader := bufio.NewReader(conn)

	// Read and process headers
	headers := make(map[string]string)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		if len(line) == 0 {
			break // End of headers
		}

		// Process headers
		parts := strings.SplitN(string(line), ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	// Handle body based on headers
	if length, ok := headers["Content-Length"]; ok {
		// Read fixed number of bytes
		len, _ := strconv.Atoi(length)
		body := make([]byte, len)
		_, err := io.ReadFull(reader, body)
		if err != nil {
			fmt.Println("Error reading body:", err)
			return
		}
		// fmt.Println(string(body))
		htmlContent = htmlContent + string(body)
	} else if headers["Transfer-Encoding"] == "chunked" {
		// Handle chunked transfer encoding
		for {
			// Read the size of the chunk
			sizeStr, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading chunk size:", err)
				return
			}
			sizeStr = strings.Trim(sizeStr, "\r\n")
			size, err := strconv.ParseInt(sizeStr, 16, 64)
			if err != nil {
				fmt.Println("Error parsing chunk size:", err)
				return
			}
			if size == 0 {
				break // End of body
			}

			// Read the chunk
			chunk := make([]byte, size)
			_, err = io.ReadFull(reader, chunk)
			if err != nil {
				fmt.Println("Error reading chunk:", err)
				return
			}
			// fmt.Print(string(chunk))
			htmlContent = htmlContent + string(chunk)

			// Read the trailing \r\n after the chunk
			_, err = reader.Discard(2)
			if err != nil {
				fmt.Println("Error discarding chunk end:", err)
				return
			}
		}
	}

	// content, err := FindStringInTag(htmlContent, tag)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("Content inside <%s>: %s\n", tag, content)

	contents, err := FindContentByClass(htmlContent, "tooltip-vitrine-btn-ofertas")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, content := range contents {
		fmt.Println(content)
	}
}
