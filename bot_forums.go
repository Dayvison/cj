package main

import (
	"fmt"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

// GetUserBio scrapes a user page and returns the bio text.
func (app App) GetUserBio(url string) (string, error) {
	var result string
	var err error
	var req *http.Request
	var response *http.Response

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return result, err
	}
	if response, err = app.httpClient.Do(req); err != nil {
		return result, err
	}

	path, err := xmlpath.Compile(`//*[@id="collapseobj_aboutme"]/div/ul/li[1]/dl/dd[1]`)
	if err != nil {
		return result, err
	}

	root, err := xmlpath.ParseHTML(response.Body)
	if err != nil {
		return result, err
	}
	result, ok := path.String(root)
	if !ok {
		return result, fmt.Errorf("xmlpath did not return a result")
	}

	return result, nil
}

// GetFirstTenUserVisitorMessages returns up to ten visitor messages from
func (app App) GetFirstTenUserVisitorMessages(url string) ([]string, error) {
	var result []string
	var err error
	var req *http.Request
	var response *http.Response

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return result, err
	}
	if response, err = app.httpClient.Do(req); err != nil {
		return result, err
	}

	path, err := xmlpath.Compile(`//*[@id="message_list"]/li/*/div[2]`)
	if err != nil {
		return result, err
	}

	root, err := xmlpath.ParseHTML(response.Body)
	if err != nil {
		return result, err
	}

	if !path.Exists(root) {
		return result, fmt.Errorf("xmlpath did not return a result")
	}

	messageBlock := path.Iter(root)
	for messageBlock.Next() {
		result = append(result, messageBlock.Node().String())
	}

	return result, nil
}