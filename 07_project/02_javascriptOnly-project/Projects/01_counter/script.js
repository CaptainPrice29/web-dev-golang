{
  /* <div class="card" id="bikeing">
            <i class="fas fa-biking fa-5x"></i>
            <div class="counter" data-target="10000"></div>
            <p>BIKING</p>
        </div>
        <div class="card" id="running">
            <i class="fas fa-running fa-5x"></i>
            <div class="counter" data-target="4000"></div>
            <p>RUNNING</p>
        </div>
        <div class="card" id="skating">
            <i class="fas fa-skating fa-5x"></i>
            <div class="counter" data-target="2000"></div>
            <p>SKATING</p>
        </div> */
}
const counters = document.querySelectorAll(".counter");

counters.forEach((counter) => {
  counter.innerHTML = 0;
  function increment() {
    const target = counter.getAttribute("data-target");
    const initialValue = Number(counter.innerHTML);
    let incr = target / 100;
    if (initialValue < target) {
      counter.innerHTML = `${Math.round(initialValue + incr)}`;
      setTimeout(increment, 20);
    } else {
        counter.innerHTML=target
      console.log("stop");
      return;
    }
    console.log("start");
  }
  increment();
});
