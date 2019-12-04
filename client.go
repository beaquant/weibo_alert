package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client          *http.Client
	userName        map[string]string
	userContainedId map[string]string
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		client: httpClient,
	}

	return c
}

func (c *Client) SendNewWeiboRequest(u string) ([]byte, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	req.Header.Add("Referer", "https://m.weibo.cn/")

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_ = res.Body.Close()

	return bodyByte, nil
}
