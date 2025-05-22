// Global reference to the folder where right-click happened
let currentFolder = null;

document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll(".folder-toggle").forEach(toggle => {
    toggle.addEventListener("click", function () {
      const nested = this.parentElement.querySelector(".nested");
      nested.classList.toggle("active");
    });

    // Right-click
    toggle.addEventListener("contextmenu", function (e) {
      e.preventDefault();
      currentFolder = this.parentElement.querySelector(".nested");

      // Show context menu
      const menu = document.getElementById("context-menu");
      menu.style.display = "block";
    });
  });
});

document.addEventListener("click",() =>{
  document.getElementById("context-menu").style.display = "none";

});

function contextAction(type) {
  if (!currentFolder) return;

  const name = prompt(`Enter ${type} name:`);
  if (!name) return;

  if (type === "folder") {
    const li = document.createElement("li");

    const span = document.createElement("span");
    span.className = "folder-toggle";
    span.textContent = "> " + name;

    const nested = document.createElement("ul");
    nested.className = "nested active";

    span.addEventListener("click", function () {
      nested.classList.toggle("active");
    });

    span.addEventListener("contextmenu", function (e) {
      e.preventDefault();
      currentFolder = nested;

      const menu = document.getElementById("context-menu");

      menu.style.display = "block";
    });

    li.appendChild(span);
    li.appendChild(nested);
    currentFolder.appendChild(li);
    currentFolder.classList.add("active");
  } else if (type === "file") {
    const li = document.createElement("li");
    li.textContent = name;
    currentFolder.appendChild(li);
    currentFolder.classList.add("active");
  }

  document.getElementById("context-menu").style.display = "none";
}
