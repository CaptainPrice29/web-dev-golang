function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(showPosition);
  } else {
    document.getElementById("weathercontainer").innerHTML = "Geolocation is not supported by this browser.";
  }
}

function showPosition(position) {
  const lat = position.coords.latitude;
  const lon = position.coords.longitude;
  async function weather() {
    try {
      const ap = await fetch(
        `https://api.weatherbit.io/v2.0/current?lat=${lat}&lon=${lon}&key=b266fc76c148406ebbc63794eccd351e`,
        { headers: { Accept: "application/json" } }
      );
      console.log(ap);
      const app = await ap.json();
      const job = app.data[0];
      document.getElementById("icon").src = `/icons/${job.weather.icon}.png`;
      document.getElementById("weather").innerHTML = job.weather.description;
      document.getElementById("temp").innerHTML = `${job.temp}℃`;
      document.getElementById(
        "feelslike"
      ).innerHTML = `Feels Like ${job.app_temp}℃`;
      document.getElementById("aqi").innerHTML = "AQI" + " " + job.aqi;
      document.getElementById("city").innerHTML = job.city_name;

      console.log(app);
    } catch (error) {
      console.log(error);
    }
  }
  weather();
}
getLocation();

const file = document.querySelector('#profilepic');
file.addEventListener('change', (e) => {
  // Get the selected file
  const [file] = e.target.files;
  // Get the file name and size
  const { name: fileName, size } = file;
  // Convert size in bytes to kilo bytes
  const fileSize = (size / 1000).toFixed(2);
  // Set the text content
  const fileNameAndSize = `${fileName} - ${fileSize}KB`;
  document.querySelector('.file-name').textContent = fileNameAndSize;
});
