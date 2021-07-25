add_transition_animation();

function add_transition_animation() {
    for (let ele of document.getElementsByClassName("transition")) {
        ele.addEventListener("click", function (e) {
            e.preventDefault();
            loading();
            window.setTimeout(function () {
                window.location.href = ele.href;
            }, 500);
        })
    }
}
