package main

import (
	"errors"
	"fmt"
	"github.com/beaquant/utils/json_file"
	"github.com/beaquant/utils/wx"
	"github.com/urfave/cli/v2"
	"github.com/wuYin/logx"

	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func init() {

}

func main() {
	logger, err := logx.LoadLogger("logger-config-sample.json")
	if err != nil {
		panic(err)
	}
	defer logger.Close()
	configFile := "config-sample.json"
	app := &cli.App{
		Name:    "weibo_alert",
		Usage:   "monitor weibo user to send a notice!",
		Version: "v0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "config-sample.json",
				Usage: "configuration file with json",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				if strings.Contains(c.String("config"), ".json") {
					configFile = c.String("config")
					logger.Info("input config file:%s", configFile)
				} else {
					logger.Info("input config file isn't a json file")
					return errors.New("input config file isn't a json file")
				}
				return nil
			}
			return errors.New("input config file!")
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
		return
	}
	c := &Config{}
	json_file.Load(configFile, c)
	notify := wx.NewWxPush(c.Notifier.Url, c.Notifier.Key)

	client := NewClient(http.DefaultClient)

	tick := time.NewTicker(10 * time.Second)
	exitSignal := make(chan os.Signal, 1)
	sigs := []os.Signal{os.Interrupt, syscall.SIGILL, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGTERM}
	signal.Notify(exitSignal, sigs...)
	for {
		select {
		case <-exitSignal:
			tick.Stop()
			return
		case <-tick.C:
			for _, v := range c.WeiboId {
				err, _, weiboId, userName, content, _ := client.GetWeiboData(v.Uid)
				if err != nil {
					continue
				}
				logger.Warn(fmt.Sprintf("[%s]发表了微博 https://m.weibo.cn/detail/%s", userName, weiboId))
				text := fmt.Sprintf("[%s]发表了微博", userName)
				desp := fmt.Sprintf(`链接: https://m.weibo.cn/detail/%s

			预览: %s`, weiboId, content)

				notify.SendWxString(text, desp)
			}
		}
	}
}
