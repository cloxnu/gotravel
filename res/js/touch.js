add_touch();

function add_touch() {
    for (let element of document.getElementsByClassName("touch")) {
        element.addEventListener("touchstart", function (e) {
            e.stopPropagation();
            for (let otherElement of document.getElementsByClassName("touch")) {
                if (otherElement !== element) {
                    otherElement.classList.remove("on-touch");
                }
            }
            element.classList.add("on-touch");
        });
    }
}
