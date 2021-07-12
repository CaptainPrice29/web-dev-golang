//[&#xf2cb;] [&#xf2ca;] [&#xf2c9;]  [&#xf2c8;] [&#xf2c7;]
//  const elmt = document.getElementById("therma");
//  function blink() {
//    elmt.innerHTML = "&#xf2cb";
//    elmt.style.color = "green";
//    setTimeout(() => {
//      elmt.innerHTML = "&#xf2ca";
//      elmt.style.color = "greenyellow";
//    }, 1000);
//    setTimeout(() => {
//      elmt.innerHTML = "&#xf2c9";
//      elmt.style.color = "yellow";
//    }, 2000);
//    setTimeout(() => {
//      elmt.innerHTML = "&#xf2c8";
//      elmt.style.color = "pink";
//    }, 3000);
//    setTimeout(() => {
//      elmt.innerHTML = "&#xf2c7";
//      elmt.style.color = "rgb(156, 50, 52)";
//    }, 4000);
//  }
//  blink();
// const stop=setInterval(blink, 5000);
// setTimeout(() => {
//     clearInterval(stop)
// }, 10000);

//different way
const elmt = document.getElementById("therma");
function blink() {
  elmt.innerHTML = "&#xf2cb";
  elmt.style.color = "green";
  setTimeout(() => {
    elmt.innerHTML = "&#xf2ca";
    elmt.style.color = "greenyellow";
    setTimeout(() => {
      elmt.innerHTML = "&#xf2c9";
      elmt.style.color = "yellow";
      setTimeout(() => {
        elmt.innerHTML = "&#xf2c8";
        elmt.style.color = "pink";

        setTimeout(() => {
          elmt.innerHTML = "&#xf2c7";
          elmt.style.color = "rgb(156, 50, 52)";
        }, 1000);
      }, 1000);
    }, 1000);
  }, 1000);
}
blink();
const stop = setInterval(blink, 5000);
setTimeout(() => {
  clearInterval(stop);
}, 10000);
