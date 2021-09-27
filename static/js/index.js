var videoFolderList = []
function getVideoFolderList(videoType){
  console.log(videoType)
  $.get("/videoList", function(data) {
    console.log(data)
    var text = ""
    Object.keys(data.data).forEach(key => {
      text += `<a href="/videoList/${data.data[key]}">`
      text += `<div class="cell btn">`
      text += `<div>`
      text += `<img class="cover" src="/play/${data.data[key]}_cover.jpg">`
      text += `</img>`
      text += `</div>`
      text += `<div class="videoTitle">`
      text += `${data.data[key]}`
      text += `</div>`
      text += `</div>`
      text += "</a>"
      console.log(data.data[key])
      videoFolderList.push(data.data[key])
    })
    $("#content").html(text)
  });
}



$(document).ready(function(){
  $("#show").click(()=>{
    $("#media").css("display", "block")
    $(".shadow").css("display", "block")
    var text = `<video controls>`
    text += `<source src="onePunch01.mp4">`
    text += `<source src="/play/D:/downloads/video/One punch/play/onePunch01.mkv">`
    text += `<track src="onePunch01.vtt" id="track01" kind="subtitles" label="ch" srclang="zh" default>`
    text += '</video>'
    $("#mediaBody").html(text)
  })

  $("#meidaClose").click(()=>{
    $("#media").css("display", "none")
    $(".shadow").css("display", "none")
    $("#mediaBody").html("")
  })
  $(".shadow").click(()=>{
    $("#media").css("display", "none")
    $(".shadow").css("display", "none")
    $("#mediaBody").html("")
  })



  /*$.get("/videoList", function(data) {
    console.log(data)
    Object.keys(data.data).forEach(key => {
      //text += `<div onclick="show_video('[ANK-Raws] ワンパンマン 02 (BDrip 1920x1080 HEVC-YUV420P10 FLAC SUP).mkv')">${data.data[key]}</div>`
      console.log(data.data[key])
    })
    //$("#one_punch").html(text)
  });*/

  

  function show_video(filePath){
    console.log(filePath)
    $.get(`/play/${filePath}`, function(data) {
      console.log(data)
    });
  }
  
  
})