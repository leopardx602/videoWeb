package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leopardx602/goTool"
)

type VideoInfo struct {
	name     string
	episode  []string
	synopsis string
}

type Video struct {
	Animation []VideoInfo
	Movie     []VideoInfo
	Series    []VideoInfo
}

func main() {
	workPath := "D:/downloads/video/"

	//var video Video

	video := map[string][]string{}
	videoList := make(map[string]VideoInfo)

	videoType, _ := ioutil.ReadDir(workPath)
	for _, vt := range videoType { // animaion
		videoName, _ := ioutil.ReadDir(workPath + "/" + vt.Name())
		for _, vn := range videoName { // one punch
			video[vt.Name()] = append(video[vt.Name()], vn.Name())
			var videoTmp VideoInfo
			videoTmp.name = vn.Name()
			videoEpisode, _ := ioutil.ReadDir(workPath + "/" + vt.Name() + "/" + vn.Name())

			for _, videoFile := range videoEpisode {
				fmt.Println(videoFile.Name())
				if strings.Contains(videoFile.Name(), ".mkv") {
					videoTmp.episode = append(videoTmp.episode, videoFile.Name())
				}
			}

			// get synopsis
			data, err := goTool.ReadJson(workPath + vt.Name() + "/" + vn.Name() + "/info.json")
			if err != nil {
				fmt.Println(err)
				videoList[vn.Name()] = videoTmp
				continue
			}
			videoTmp.synopsis = data["synopsis"].(string)
			videoTmp.name = data["name"].(string)
			videoList[vn.Name()] = videoTmp
		}

	}
	fmt.Println(video)
	fmt.Println("===============")
	fmt.Println(videoList)

	// // http server
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.GET("/animation", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"videoFolderList": video["animation"]})
	})
	router.GET("/movie", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"videoFolderList": video["movie"]})
	})
	router.GET("/series", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"videoFolderList": video["series"]})
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
		ctx.JSON(200, gin.H{"synopsis": videoList[videoName].synopsis, "name": videoList[videoName].name})
	})

	router.GET("/play/:videoPath", func(ctx *gin.Context) {
		videoPath := ctx.Param("videoPath")
		videoPath = strings.Replace(videoPath, "_", "/", -1)
		fmt.Println(workPath + videoPath)
		ctx.File(workPath + "/animation/" + videoPath)
	})

	// css javascript
	router.Static("/static", "./static")

	router.Run(":5000")
}
