package main

import (
	"strings"

	"golang.org/x/net/html"
)

// FindStringInTag finds all occurrence of a specified tag and returns its inner content.
func FindStringInTag(htmlContent string, tagName string) ([]string, error) {
	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var results []string

	// Traverse the HTML nodes
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			if n.FirstChild != nil {
				results = append(results, n.FirstChild.Data)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	// Start traversing from the root node
	traverse(doc)

	return results, nil
}

// FindContentByID searches for the content inside the first HTML element with a specific id.
func FindContentByID(htmlContent string, id string) (string, error) {
	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	// Traverse the HTML nodes
	var traverse func(*html.Node) string
	traverse = func(n *html.Node) string {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					if n.FirstChild != nil {
						return n.FirstChild.Data
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result := traverse(c)
			if result != "" {
				return result
			}
		}
		return ""
	}

	// Start traversing from the root node
	return traverse(doc), nil
}

// FindContentByClass searches for the content inside all HTML elements with a specific class.
func FindContentByClass(htmlContent string, className string) ([]string, error) {
	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var results []string

	// Traverse the HTML nodes
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "class" {
					classes := strings.Fields(attr.Val)
					for _, c := range classes {
						if c == className {
							if n.FirstChild != nil {
								// Add the content of the first child node to results
								results = append(results, n.FirstChild.Data)
							}
							break // Found the class, no need to check other classes
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	// Start traversing from the root node
	traverse(doc)

	return results, nil
}
