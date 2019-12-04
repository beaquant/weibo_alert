package main

import (
	"encoding/csv"
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

var (
	SaveWeibos []SaveWeibo
)

func openFile(fileName string) (bool, *os.File) {
	var file *os.File
	var err1 error
	var isNew = false
	checkFileIsExist := func(fileName string) bool {
		var exist = true
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			exist = false
		}
		return exist
	}
	if checkFileIsExist(fileName) {
		file, err1 = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 666)
	} else {
		file, err1 = os.Create(fileName)
		isNew = true
	}
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "unable to write file on filehook %v", err1)
		panic(err1)
	}
	return isNew, file
}

func init() {
	SaveWeibos = make([]SaveWeibo, 0)
}

func main() {
	logger, err := logx.LoadLogger("logger-config-sample.json")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	isNew, weiboSaveFile := openFile("weiboSave.csv")
	defer weiboSaveFile.Close()
	weiboSaveFileCsv := csv.NewWriter(weiboSaveFile)
	if isNew {
		data := []string{"UTime", "用户名", "微博id", "文章"}
		weiboSaveFileCsv.Write(data)
		weiboSaveFileCsv.Flush()
	}

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
			if strings.Contains(c.String("config"), ".json") {
				configFile = c.String("config")
				logger.Info("input config file:%s", configFile)
			} else {
				logger.Info("input config file isn't a json file")
				return errors.New("input config file isn't a json file")
			}
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
		return
	}
	c := &Config{}
	err = json_file.Load(configFile, c)
	if err != nil {
		logger.Fatal(err)
		return
	}

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
				err, _, weiboId, userName, content, ts := client.GetWeiboData(v.Uid)
				if err != nil {
					continue
				}
				if len(SaveWeibos) > 0 && weiboId == SaveWeibos[len(SaveWeibos)-1].WeiboId {
					continue
				}
				SaveWeibos = append(SaveWeibos, SaveWeibo{
					Timestamp: ts.Unix(),
					WeiboId:   weiboId,
					Content:   content,
				})
				logger.Warn(fmt.Sprintf("[%s]发表了微博 https://m.weibo.cn/detail/%s", userName, weiboId))
				text := fmt.Sprintf("[%s]发表了微博", userName)
				desp := fmt.Sprintf(`链接: https://m.weibo.cn/detail/%s

			预览: %s`, weiboId, content)

				notify.SendWxString(text, desp)

				data := []string{
					fmt.Sprint(ts.Unix()),
					userName,
					weiboId,
					content,
				}

				weiboSaveFileCsv.Write(data) //写入数据
				weiboSaveFileCsv.Flush()

			}
		}
	}
}
