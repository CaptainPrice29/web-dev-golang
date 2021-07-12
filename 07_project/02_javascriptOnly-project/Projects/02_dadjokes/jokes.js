// const joke = document.querySelector("#joke");
// const jokebtn = document.querySelector("#jokebtn");
// const generatejokes = () => {
//   var requestOptions = {
//     headers: {
//       Accept: "application/json",
//     },
//   };

//   fetch("https://icanhazdadjoke.com/", requestOptions)
//     .then((response) => response.json())
//     .then((result) => (jokes.innerHTML = result.joke))
//     .catch((error) => console.log("error", error));
// };

// jokebtn.addEventListener("click", generatejokes);
// generatejokes();

//modification
// function api() {
 
//   fetch("https://icanhazdadjoke.com/",{headers:{Accept:'application/json'}})
//     .then((response) => response.json())
//     .then((result) => (document.querySelector("#jokes").innerHTML = result.joke))
//     .catch((error) => console.log("error", error));
// }
// api();
// document.querySelector("#jokebtn").addEventListener("click", api);


//==>using async await
const api = async () => {
  try {
    const res = await fetch("https://icanhazdadjoke.com/", { headers: { Accept: 'application/json' } })
    const data = await res.json()
    document.querySelector('#jokes').innerHTML=data.joke
  } catch (error) {
    console.log(error);
    
  }
}
api()
document.querySelector('#jokebtn').addEventListener('click',api)
