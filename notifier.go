package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var ScKey string

func (c *Client) PushToWechat(userId string, weiboId string, username string, content string, publishedAt time.Time) error {
	//log.Println(fmt.Sprintf("[%s]发表了微博 https://m.weibo.cn/detail/%s", username, weiboId))

	u, err := url.Parse(fmt.Sprintf("https://sc.ftqq.com/%s.send", ScKey))
	if err != nil {
		return err
	}

	q := u.Query()
	q.Add("text", fmt.Sprintf("[%s]发表了微博", username))
	q.Add("desp", fmt.Sprintf(`链接: https://m.weibo.cn/detail/%s

预览: %s`, weiboId, content))

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	_ = res.Body.Close()

	return nil
}
