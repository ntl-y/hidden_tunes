document.addEventListener("DOMContentLoaded", function() {
    function getRandomAudio() {
        fetch('/getRandomAudio')
            .then(response => response.json())
            .then(data => {
                document.getElementById('audio-player').src = data.audiodownload;
                document.getElementById('track-name').textContent = data.name;
                document.getElementById('artist-name').textContent = data.artist_name;
                document.getElementById('album-name').textContent = data.album_name;
                document.getElementById('album-image').src = data.album_image;
            })
            .catch(error => console.error('Error:', error));
    }

    getRandomAudio();
});
