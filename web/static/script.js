const musicContainer = document.querySelector(".music-container");
const playBtn = document.querySelector("#play");
const prevBtn = document.querySelector("#prev");
const nextBtn = document.querySelector("#next");
const audio = document.querySelector("#audio-player");
const progress = document.querySelector(".progress");
const progressContainer = document.querySelector(".progress-container");
const title = document.querySelector("#track-name");
const image = document.querySelector("#album-image");

let trackList = [];
let currentTrackIndex = -1;

function setProgress(e) {
  const width = this.clientWidth;
  const clickX = e.offsetX;
  const duration = audio.duration;

  audio.currentTime = (clickX / width) * duration;
}

function updateProgress(e) {
  const { duration, currentTime } = e.srcElement;
  const progressPercent = (currentTime / duration) * 100;
  progress.style.width = `${progressPercent}%`;
}

progressContainer.addEventListener("click", setProgress);

function getRandomAudio() {
  fetch("/getRandomAudio")
    .then((response) => response.json())
    .then((data) => {
      const trackData = {
        audioSrc: data.audiodownload,
        name: data.name,
        artist: data.artist_name,
        albumImage: data.album_image,
      };

      trackList.push(trackData);
      currentTrackIndex = trackList.length - 1;

      loadTrack(trackData);
    })
    .catch((error) => console.error("Error:", error));
}

document.addEventListener("DOMContentLoaded", function () {
  getRandomAudio();
});

function loadTrack(trackData) {
  document.getElementById("audio-player").src = trackData.audioSrc;
  document.getElementById("track-name").textContent = trackData.name;
  document.getElementById("artist-name").textContent = trackData.artist;
  document.getElementById("album-image").src = trackData.albumImage;

  pauseSong();
}

function playSong() {
  musicContainer.classList.add("play");
  playBtn.querySelector("i.fas").classList.remove("fa-play");
  playBtn.querySelector("i.fas").classList.add("fa-pause");

  audio.play();
}

function pauseSong() {
  musicContainer.classList.remove("play");
  playBtn.querySelector("i.fas").classList.add("fa-play");
  playBtn.querySelector("i.fas").classList.remove("fa-pause");

  audio.pause();
}

playBtn.addEventListener("click", () => {
  const isPlaying = musicContainer.classList.contains("play");

  if (isPlaying) {
    pauseSong();
  } else {
    playSong();
  }
});

nextBtn.addEventListener("click", () => {
  if (currentTrackIndex < trackList.length - 1) {
    currentTrackIndex++;
    loadTrack(trackList[currentTrackIndex]);
  } else {
    getRandomAudio(); 
  }
});

prevBtn.addEventListener("click", () => {
  if (currentTrackIndex > 0) {
    currentTrackIndex--;
    loadTrack(trackList[currentTrackIndex]);
  }
});

audio.addEventListener("timeupdate", updateProgress);
audio.addEventListener("ended", function () {
  getRandomAudio();
  audio.onloadeddata = () => {
    playSong();
  };
});
