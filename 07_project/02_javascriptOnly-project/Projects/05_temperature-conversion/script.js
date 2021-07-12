const temp = document.getElementById("temp");
const units = document.getElementById("unit");
const result = document.getElementById("result");
function calculate() {
  if (units.options[units.selectedIndex].value == "celsius") {
    let res = (temp.value - 32) * (5 / 9);
    result.innerHTML = res + "°C";
  } else {
    let res = temp.value * (9 / 5) + 32;
    result.innerHTML = res + "°F";
  }
}
