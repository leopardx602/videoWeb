
function showEpisode(videoName, episode){
  var text = ""
  Object.keys(episode).forEach(key => {
    text += `<div class="episodeNumber btn" onclick="showVideo('${videoName}','${episode[key]}')">${parseInt(key, 10) + 1}</div>`
    console.log(episode[key])
  })
  $("#episodeList").html(text)

  $.get(`/videoList/${videoName}/info`, function(data) {
    $("#synopsis").html(data.data)
  })
}




