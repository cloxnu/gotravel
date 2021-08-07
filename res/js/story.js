function load_page() {
    // detect all_heading
    document.querySelectorAll('h2, h4').forEach((ele) => {
        let level = ele.localName === 'h2' ? 2 : 4
        all_heading.push({text: ele.innerText, level: level, id: ele.id, pos: 0})
    })

    // thumbnail image onclick
    let all_img = document.querySelectorAll("#content img");
    all_img.forEach(function (img) {
        img.onclick = function () {
            if (document.getElementById("content").classList.contains("thumbnail")) {
                for (let ele of document.getElementsByClassName("checked")) {
                    if (ele !== this)
                        ele.classList.remove("checked");
                }
                this.classList.toggle("checked");
            }
        }
    })

    if (all_heading.length !== 0) {
        document.getElementById("nav-list-div").innerHTML = generate_nav_html();
        update_heading_pos();
        listen_document_height_change();
        window.addEventListener('scroll', scroll_handler);
    } else {
        document.getElementById("nav-div").style.display = "none";
    }

    // if (blog_cover.length === 0) {
    //     document.getElementById("cover-img").style.display = "none";
    //     document.getElementById("cover-shadow").style.display = "none";
    //     document.getElementById("head").style.height = "30vh";
    //     document.getElementById("head").classList.remove("dark-mode-on");
    // } else {
    //     document.getElementById("cover-img").style.opacity = "1";
    //     document.getElementById("cover-img").src = blog_cover;
    // }

    add_title_link();
    add_touch();

    loaded(true, load_handler)
}


function load_handler() {
    if (window.location.hash)
        window.location.href = window.location.hash;
}



window.onload = function () {
    load_page()
}
