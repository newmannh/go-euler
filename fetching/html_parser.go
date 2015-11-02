package fetching

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

const (
	siteURL = "https://projecteuler.net"
)

func FetchProblem(probNum int) (string, error) {
	reader, err := getProblemHTTP(probNum)
	if err != nil {
		return "", err
	}
	return parseHTTP(reader)
}

func parseHTTP(httpContent io.Reader) (string, error) {
	doc, err := html.Parse(httpContent)
	if err != nil {
		return "", err
	}
	var f func(*html.Node) (string, error)
	f = func(n *html.Node) (string, error) {
		if containsProblem(n) {
			// Do something with n...
			return nodeToString(n), nil
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if parsedChild, err := f(c); parsedChild != "" {
				return parsedChild, err
			}
		}
		return "", fmt.Errorf("Unable to find problem.")
	}
	return f(doc)
}

func nodeToString(node *html.Node) string {
	switch node.Type {
	case html.ElementNode:
		chilrens := ""
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			chilrens = fmt.Sprintf("%s%s", chilrens, nodeToString(c))
		}
		return fmt.Sprintf("<%s>%s</%s>", node.Data, chilrens, node.Data)
	case html.TextNode:
		return node.Data
	default:
		return ""
	}
}

func containsProblem(node *html.Node) bool {
	if node.Type == html.ElementNode && node.Data == "div" {
		for _, attr := range node.Attr {
			key, val := attr.Key, attr.Val
			if key == "class" && val == "problem_content" {
				return true
			}
			if key == "role" && val == "problem" {
				return true
			}
		}
	}
	return false
}

func getProblemHTTP(probNum int) (io.Reader, error) {
	resp, err := http.Get(getURL(probNum))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func getURL(probNum int) string {
	return fmt.Sprintf("%s/problem=%d", siteURL, probNum)
}
