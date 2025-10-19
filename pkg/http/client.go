package pkg_http

import (
	"net/http"
	"time"

	logger "github.com/BMokarzel/stress-test/pkg/logger"
)

type Client struct {
	Logger *logger.Logger
	Client *http.Client
	URL    string
}

func New(url string) *Client {
	client := http.Client{
		Timeout: time.Second * 30,
	}

	return &Client{
		Client: &client,
		URL:    url,
	}
}

func (c *Client) Call() (int, error) {
	res, err := c.Client.Get(c.URL)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
