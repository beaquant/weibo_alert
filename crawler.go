package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (c *Client) GetWeiboData(uid string) (error, string, string, string, string, time.Time) {
	userName, isOk1 := c.userName[uid]
	userContainedId, isOk2 := c.userContainedId[uid]
	if !isOk1 || !isOk2 {
		username, containerId, err := c.getUsernameAndWeiboContainerId(uid)
		if err != nil {
			return err, "", "", "", "", time.Now()
		}
		userName = username
		c.userName[uid] = username
		c.userContainedId[uid] = containerId
	}

	u := fmt.Sprintf("https://m.weibo.cn/api/container/getIndex?type=uid&value=%s&containerid=%s", uid, userContainedId)
	res, err := c.SendNewWeiboRequest(u)
	if err != nil {
		panic(err)
	}

	var weibo Weibo
	err = json.Unmarshal(res, &weibo)
	if err != nil {
		panic(err)
	}

	for _, value := range weibo.Data.Cards {
		if value.Mblog.CreatedAt == "刚刚" {
			weiboId := value.Mblog.ID
			content := value.Mblog.Text
			err = c.PushToWechat(uid, weiboId, userName, content, time.Now().Local())
			return err, uid, weiboId, userName, content, time.Now().Local()
		}
	}
	return errors.New("no news"), "", "", "", "", time.Now().Local()
}

func (c *Client) getUsernameAndWeiboContainerId(userId string) (username string, containerId string, err error) {
	u := fmt.Sprintf("https://m.weibo.cn/api/container/getIndex?type=uid&value=%s", userId)

	res, err := c.SendNewWeiboRequest(u)
	if err != nil {
		return "", "", err
	}

	container := Container{}
	err = json.Unmarshal(res, &container)
	if err != nil {
		return "", "", err
	}

	if container.Ok != 1 {
		return "", "", errors.New(container.Msg)
	}

	username = container.Data.UserInfo.ScreenName

	for _, value := range container.Data.TabsInfo.Tabs {
		if value.Title == "微博" {
			containerId = value.Containerid
		}
	}

	//log.Println(username, containerId)

	return username, containerId, nil
}
