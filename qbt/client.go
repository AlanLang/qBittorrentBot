package qbt

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"

	wrapper "github.com/pkg/errors"
)

//ErrBadResponse means that qbittorrent sent back an unexpected response
var ErrBadResponse = errors.New("received bad response")

//Client creates a connection to qbittorrent and performs requests
type Client struct {
	http *http.Client
	URL  string
	Jar  http.CookieJar
}

//NewClient creates a new client connection to qbittorrent
func NewClient(url string) *Client {
	client := &Client{}

	// ensure url ends with "/"
	if url[len(url)-1:] != "/" {
		url += "/"
	}

	client.URL = url

	// create cookie jar
	client.Jar, _ = cookiejar.New(nil)
	client.http = &http.Client{
		Jar: client.Jar,
	}
	return client
}

//Login logs you in to the qbittorrent client
//returns the current authentication status
func (client *Client) Login(username string, password string) error {
	credentials := make(map[string]string)
	credentials["username"] = username
	credentials["password"] = password

	resp, err := client.post("api/v2/auth/login", credentials)
	if err != nil {
		return err
	} else if resp.Status != "200 OK" { // check for correct status code
		return wrapper.Wrap(ErrBadResponse, "couldnt log in")
	}
	return nil
}

//Sync returns the server state and list of torrents in one struct
func (client *Client) Sync(rid string) (Sync, error) {
	var s Sync

	params := make(map[string]string)
	params["rid"] = rid

	resp, err := client.get("api/v2/sync/maindata", params)
	if err != nil {
		return s, err
	}
	json.NewDecoder(resp.Body).Decode(&s)
	return s, nil
}

//post will perform a POST request with no content-type specified
func (client *Client) post(endpoint string, opts map[string]string) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for k, v := range opts {
		_ = bodyWriter.WriteField(k, v)
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest("POST", client.URL+endpoint, bodyBuf)
	if err != nil {
		return nil, wrapper.Wrap(err, "failed to build request")
	}
	req.Header.Add("content-type", contentType)

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, wrapper.Wrap(err, "failed to perform request")
	}
	defer resp.Body.Close()
	return resp, nil

}

//get will perform a GET request with no parameters
func (client *Client) get(endpoint string, opts map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", client.URL+endpoint, nil)
	if err != nil {
		return nil, wrapper.Wrap(err, "failed to build request")
	}

	// add optional parameters that the user wants
	if opts != nil {
		query := req.URL.Query()
		for k, v := range opts {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, wrapper.Wrap(err, "failed to perform request")
	}
	return resp, nil
}
