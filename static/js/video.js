function showVideo(videoName ,videoEpisode){
  console.log(videoEpisode)
  $("#media").css("display", "block")
  $(".shadow").css("display", "block")
  var text = `<video controls>`
  text += `<source src="onePunch01.mp4">`
  text += `<source src="/play/${videoName}_${videoEpisode}">`
  text += `<track src="onePunch01.vtt" id="track01" kind="subtitles" label="ch" srclang="zh" default>`
  text += '</video>'
  $("#mediaBody").html(text)
}
$(document).ready(function(){

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

  
})