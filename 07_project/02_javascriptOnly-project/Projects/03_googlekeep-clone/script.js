const addButton = document.querySelector("#add");
const updateData = () => {
  const textAreaData = document.querySelectorAll("textarea");
  const notes = [];
  textAreaData.forEach((note) => {
    return notes.push(note.value);
  });
  console.log(JSON.stringify(notes));
  console.log(notes);
  localStorage.setItem("notes", JSON.stringify(notes));
};
const addNewNote = (text = "") => {
  const note = document.createElement("div");
  note.classList.add("notes");
  const htmlData = `<div class="notes-btn">
    <button class="edit"><i class="fas fa-edit"></i></button>
    <button class="delete"><i class="fas fa-trash"></i></button>
    
</div>
<div class="main ${text ? "" : "hidden"}"></div>
    <textarea class="${text ? "hidden" : ""}"></textarea>`;
  note.insertAdjacentHTML("afterbegin", htmlData);
  document.body.appendChild(note);
  const editButton = note.querySelector(".edit");
  const deleteButton = note.querySelector(".delete");
  const mainDiv = note.querySelector(".main");
  const textArea = note.querySelector("textarea");
  deleteButton.addEventListener("click", () => {
    note.remove();
    updateData();
  });
  textArea.value = text;
  mainDiv.innerHTML = text;
  textArea.addEventListener("change", (event) => {
    const value = event.target.value;
    mainDiv.innerHTML = value;
    updateData();
  });
  editButton.addEventListener("click", () => {
    mainDiv.classList.toggle("hidden");
    textArea.classList.toggle("hidden");
  });
};
const notes = JSON.parse(localStorage.getItem("notes"));
if (notes) {
  notes.forEach((note) => addNewNote(note));
}
addButton.addEventListener("click", () => addNewNote());
