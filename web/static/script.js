document.addEventListener("DOMContentLoaded", () => {
  const interBubble = document.querySelector(".interactive");
  let curX = 0;
  let curY = 0;
  let tgX = 0;
  let tgY = 0;

  function move() {
    curX += (tgX - curX) / 20;
    curY += (tgY - curY) / 20;
    interBubble.style.transform = `translate(${Math.round(
      curX
    )}px, ${Math.round(curY)}px)`;
    requestAnimationFrame(() => {
      move();
    });
  }

  window.addEventListener("mousemove", (event) => {
    tgX = event.clientX;
    tgY = event.clientY;
  });

  move();
});

const musicContainer = document.querySelector(".music-container");
const playBtn = document.querySelector("#play");
const prevBtn = document.querySelector("#prev");
const nextBtn = document.querySelector("#next");
const audio = document.querySelector("#audio-player");
const progress = document.querySelector(".progress");
const progressContainer = document.querySelector(".progress-container");
const title = document.querySelector("#track-name");
const image = document.querySelector("#album-image");
const sound = document.querySelector(".sound-bar");
const soundContainer = document.querySelector(".sound-bar-container");

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

function setSound(e) {
  const width = soundContainer.clientWidth;
  const clickX = e.offsetX;
  const volumeLevel = clickX / width;

  audio.volume = volumeLevel;
  sound.style.width = `${volumeLevel * 100}%`;
}

function dragSound(e) {
  const width = soundContainer.clientWidth; 
  const dragX = e.offsetX; 
  const volumeLevel = dragX / width; 

  audio.volume = volumeLevel; 
  sound.style.width = `${volumeLevel * 100}%`;
}

soundContainer.addEventListener("mousedown", () => {
  soundContainer.addEventListener("mousemove", dragSound);
});

window.addEventListener("mouseup", () => {
  soundContainer.removeEventListener("mousemove", dragSound);
});

progressContainer.addEventListener("click", setProgress);
soundContainer.addEventListener("click", setSound);


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