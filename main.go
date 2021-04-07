// updater_api project main.go
// simple updater v0.1
// by @szradev

package main

import (
	"fmt"
	//"strconv"
	"gopkg.in/ini.v1"
	//"time"

	"github.com/gin-gonic/gin"
)

var err error

type Entry struct {
	DateTime string `json:"date_time"`
	FileName string `json:"filename"`
	Id       string `json:"id"`
	RomType  string `json:"romtype"`
	Size     int    `json:"size"`
	Url      string `json:"url:`
	Version  string `json:"version"`
}

var Response []Entry

func main() {
	//	data, err = ini.Load("my.ini")
	fmt.Println("Welcome UPDATER API v01.")

	r := gin.Default()

	update := Entry{}
	update.DateTime = "dupa"
	update.FileName = "dupa"
	update.Id = "dupa"
	update.RomType = "dupa"
	update.Size = 1111111
	update.Url = "dupa"
	update.Version = "dupa"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": []Entry{update},
		})
	})
	r.Run(":1111")
}
