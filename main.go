package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leopardx602/goTool"
)

type Video struct {
	name     string
	episode  []string
	synopsis string
}

func main() {
	workPath := "D:/downloads/video/"

	// get video folders
	var videoFolderList []string
	files, _ := ioutil.ReadDir(workPath)
	for _, f := range files {
		//fmt.Println(f.Name())
		videoFolderList = append(videoFolderList, f.Name())
	}
	fmt.Println(videoFolderList)

	// get videos
	videoList := make(map[string]Video)
	for _, folder := range videoFolderList {
		// get episode
		files2, _ := ioutil.ReadDir(workPath + folder)
		var videoTmp Video
		videoTmp.name = folder
		for _, video := range files2 {
			fmt.Println(video.Name())
			if strings.Contains(video.Name(), ".mkv") {
				videoTmp.episode = append(videoTmp.episode, video.Name())
			}
		}

		// get synopsis
		data, err := goTool.ReadJson(workPath + folder + "/info.json")
		if err != nil {
			fmt.Println(err)
			videoList[folder] = videoTmp
			continue
		}
		videoTmp.synopsis = data["synopsis"].(string)
		videoList[folder] = videoTmp
	}
	fmt.Println("==============")
	fmt.Println("videoList", videoList)

	// http server
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/animation", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"videoFolderList": videoFolderList})
	})
	router.GET("/movie", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"videoFolderList": videoFolderList})
	})
	router.GET("/series", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"videoFolderList": videoFolderList})
	})

	// router.GET("/videoList", func(ctx *gin.Context) {
	// 	fmt.Println(videoFolderList)
	// 	ctx.JSON(200, gin.H{"data": videoFolderList})
	// })

	router.GET("/videoList/:videoName", func(ctx *gin.Context) {
		videoName := ctx.Param("videoName")
		fmt.Println(videoName)
		ctx.HTML(200, "episode.html", gin.H{"episode": videoList[videoName].episode, "videoName": videoName})
	})

	router.GET("/videoList/:videoName/info", func(ctx *gin.Context) {
		videoName := ctx.Param("videoName")
		fmt.Println(videoName)
		ctx.JSON(200, gin.H{"data": videoList[videoName].synopsis})
	})

	router.GET("/play/:videoPath", func(ctx *gin.Context) {
		videoPath := ctx.Param("videoPath")
		videoPath = strings.Replace(videoPath, "_", "/", -1)
		fmt.Println(workPath + videoPath)
		ctx.File(workPath + videoPath)
	})

	// css javascript
	router.Static("/static", "./static")

	router.Run(":5000")
}
