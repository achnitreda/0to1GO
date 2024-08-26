const dropdownToggle = document.querySelector(".dropdown-toggle");
const dropdownMenu = document.querySelector(".dropdown-menu");
const textFile = document.getElementById("option1");

dropdownToggle.addEventListener("click", function () {
    if (dropdownMenu.style.display === "flex") {
        dropdownMenu.style.display = "none";
    } else {
        dropdownMenu.style.display = "flex";
    }
});

window.addEventListener("click", (event) => {
    if (!event.target.matches(".dropdown-toggle")) {
        if (dropdownMenu.style.display === "flex") {
            dropdownMenu.style.display = "none";
        }
    }
});

textFile.addEventListener("click", (e) => {
    e.preventDefault();
    downloadResult("text");
});

function downloadResult(fileType) {
    const result = document.getElementById("res").innerText;
    const link = document.createElement("a");
    link.href = `/download?type=${fileType}&content=${encodeURIComponent(result)}`;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
}