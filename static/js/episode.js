
function showEpisode(videoName, episode){
  var text = ""
  Object.keys(episode).forEach(key => {
    text += `<div class="episodeNumber btn" onclick="showVideo('${videoName}','${episode[key]}')">${key}</div>`
    console.log(episode[key])
  })
  $("#episodeList").html(text)
}
