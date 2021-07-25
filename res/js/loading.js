function loaded(first_show=true, handler=function () {}) {
    document.body.classList.remove("loading");
    document.getElementById("loading-shelter").style.display = "none";
    window.setTimeout(function () {
        document.getElementById("loading-div").style.opacity = "0";
        handler();
    }, 500);
    if (first_show)
        firebase_init();
}

function loading() {
    document.getElementById("loading-div").style.opacity = "1";
    document.body.classList.add("loading");
    window.setTimeout(function () {
        document.getElementById("loading-shelter").style.display = "block";
    }, 500);
}

let isFirstShow = true;
window.onpageshow = function () {
    if (isFirstShow) {
        isFirstShow = false;
    } else {
        loaded(false);
    }
}
