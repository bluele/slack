package main

import (
	"fmt"
	"github.com/bluele/slack"
	"path/filepath"
	"strings"
)

// Please change these values to suit your environment
const (
	token             = "your-api-token"
	channelName       = "general"
	uploadFilePath    = "./assets/test.txt"
	uploadPicFilePath = "./assets/Lenna.png"
)

func main() {
	api := slack.New(token)
	channel, err := api.FindChannelByName(channelName)
	if err != nil {
		panic(err)
	}

	err = api.FilesUpload(&slack.FilesUploadOpt{
		Filepath: uploadFilePath,
		Filetype: "text",
		Filename: filepath.Base(uploadFilePath),
		Title:    "upload test1",
		Channels: []string{channel.Id},
	})
	if err != nil {
		panic(err)
	}

	err = api.FilesUpload(&slack.FilesUploadOpt{
		Filetype: "text",
		Content:  strings.Repeat("a", 10000),
		Title:    "upload test2",
		Channels: []string{channel.Id},
	})
	if err != nil {
		panic(err)
	}

	err = api.FilesUpload(&slack.FilesUploadOpt{
		Filepath: uploadPicFilePath,
		Title:    "upload test3",
		Filename: filepath.Base(uploadPicFilePath),
		Channels: []string{channel.Id},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Completed file upload.")
}
