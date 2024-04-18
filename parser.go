package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var listOfLinks []Link
	linkNodesList := listLinkNodes(doc)

	for _, node := range linkNodesList {
		fmt.Println(&node)
		listOfLinks = append(listOfLinks, buildLinkFromNode(node))
	}
	fmt.Println(listOfLinks)
	return listOfLinks, nil
}

func buildLinkFromNode(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
		}
	}
	ret.Text = "buildTextFromNode(n)"
	return ret
}

func buildTextFromNode(n *html.Node) string {
	fmt.Println(html.TextNode)
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		ret += buildTextFromNode(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func listLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var list []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		list = append(list, listLinkNodes(c)...)
	}
	return list
}

func dfs(n *html.Node, pad string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + n.Data + ">"
	}
	fmt.Println(pad + msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, pad+"      ")
	}
}
