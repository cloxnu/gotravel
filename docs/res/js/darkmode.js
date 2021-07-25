let is_dark_mode = false;
let dark_mode_class_name = "dark-mode-off";

change_dark_mode(window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches);

window.matchMedia('(prefers-color-scheme: dark)').addListener(e => {
    const isDarkMode = e.matches;
    change_dark_mode(isDarkMode);
    application_to_html();
});

function dark_mode_clicked() {
    change_dark_mode(!is_dark_mode);
    application_to_html();
}

function change_dark_mode(isDarkMode) {
    is_dark_mode = isDarkMode;
    if (isDarkMode) {
        dark_mode_class_name = "dark-mode-on";
    }
    else {
        dark_mode_class_name = "dark-mode-off";
    }
}

function application_to_html() {
    document.body.classList.remove(is_dark_mode ? "dark-mode-off" : "dark-mode-on");
    document.body.classList.add(is_dark_mode ? "dark-mode-on" : "dark-mode-off");
}
