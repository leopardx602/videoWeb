package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
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
	videoEpisodeList := make(map[string][]string)
	for _, folder := range videoFolderList {
		//files2, _ := filepath.Glob(workPath + folder + "/*.mkv")
		files2, _ := ioutil.ReadDir(workPath + folder)

		var videoTmp Video
		videoTmp.name = folder
		for _, video := range files2 {
			fmt.Println(video.Name())
			if strings.Contains(video.Name(), ".mkv") {
				videoEpisodeList[folder] = append(videoEpisodeList[folder], video.Name())
				videoTmp.episode = append(videoTmp.episode, video.Name())
			}
		}
		videoList[folder] = videoTmp
	}
	fmt.Println("videoEpisodeList", videoEpisodeList)
	fmt.Println("==============")
	fmt.Println("videoList", videoList)

	// get episode

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/videoList", func(ctx *gin.Context) {
		fmt.Println(videoFolderList)
		ctx.JSON(200, gin.H{"data": videoFolderList})
	})

	router.GET("/videoList/:videoName", func(ctx *gin.Context) {
		videoName := ctx.Param("videoName")
		fmt.Println(videoName)
		ctx.HTML(200, "episode.html", gin.H{"episode": videoEpisodeList[videoName], "videoName": videoName})
	})

	router.GET("/play/:videoPath", func(ctx *gin.Context) {
		videoPath := ctx.Param("videoPath")
		videoPath = strings.Replace(videoPath, "_", "/", -1)
		fmt.Println(workPath + videoPath)
		ctx.File(workPath + videoPath)
	})

	router.Static("/static", "./static")

	router.Run(":5000")
}
