package gdl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Error struct {
	Response *http.Response
	Message  string `json:"error,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Request failed for %v, with code %v and message %v \n", e.Response.Request.URL.String(), e.Response.StatusCode, e.Message)
}


func (c *GDLClient) NewRequest(method, urlStr string, body string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	q := u.Query()
	
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	return req, nil
}
 
func (c *GDLClient) Perform(req *http.Request) ([]byte, error) {

	req.Header.Add("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return ioutil.ReadAll(resp.Body)
	case 400:
		return nil, ErrBadRequest
	case 403:
		return nil, ErrUnauthorized
	case 404:
		return nil, ErrNotFound
	case 500:
		return nil, ErrGDLInternal
	default:
		return nil, ErrUnsupportedStatusCode
	}
}

  