
let all_heading = [];

function generate_nav_html() {
    let nav_html = "<ul>";
    let is_last_h2 = true;
    for (let heading_item of all_heading) {
        if (heading_item.level === 2) {
            if (!is_last_h2) {
                nav_html += "</ul>";
                is_last_h2 = true;
            }
            nav_html += `<li id="nav-${heading_item.id}" class="h2 touch" onclick="scroll_to('${heading_item.pos - 90}')"><span>${heading_item.text}</span></li>`;
        } else if (heading_item.level === 4) {
            if (is_last_h2) {
                nav_html += "<ul>";
                is_last_h2 = false;
            }
            nav_html += `<li id="nav-${heading_item.id}" class="h4 touch" onclick="scroll_to('${heading_item.pos - 90}')"><span>${heading_item.text}</span></li>`;
        }
    }
    if (!is_last_h2)
        nav_html += "</ul>";
    nav_html += "</ul>";
    return nav_html;
}

function scroll_to(pos) {
    window.scrollTo({
        top: pos,
        left: 0,
        behavior: "smooth"
    });
}

function update_heading_pos() {
    for (let heading_item of all_heading) {
        heading_item.pos = document.getElementById(heading_item.id).offsetTop;
        document.getElementById(`nav-${heading_item.id}`).onclick = function () {
            scroll_to(`${heading_item.pos - 90}`);
        }
        // heading_item.pos = document.getElementById(heading_item.id).getBoundingClientRect().top;
    }
    scroll_handler();
    console.log("heading position updated.");
    console.log(all_heading);
}

function update_current_heading(heading) {
    document.getElementById("nav-title").innerHTML = heading.text;
    for (let ele of document.getElementsByClassName("nav-active")) {
        ele.classList.remove("nav-active");
    }
    document.getElementById(`nav-${heading.id}`).classList.add("nav-active");
}

function listen_document_height_change(){
    let lastHeight = document.body.clientHeight, newHeight;
    (function run(){
        newHeight = document.body.clientHeight;
        if (lastHeight !== newHeight)
            update_heading_pos();
        lastHeight = newHeight;

        if (document.body.onElementHeightChangeTimer)
            clearTimeout(document.body.onElementHeightChangeTimer);

        document.body.onElementHeightChangeTimer = setTimeout(run, 200);
    })()
}

function scroll_handler() {
    // calculate last_heading
    const heading_expose_height = window.pageYOffset + 100;
    let passed_heading = all_heading.filter(item => item.pos < heading_expose_height);
    if (passed_heading.length === 0)
        passed_heading = all_heading;
    let last_heading;
    let min_distance;
    for (let heading_item of passed_heading) {
        if (min_distance === undefined || Math.abs(heading_item.pos - heading_expose_height) < min_distance) {
            min_distance = Math.abs(heading_item.pos - heading_expose_height);
            last_heading = heading_item;
        }
    }
    update_current_heading(last_heading);
}

function nav_clicked() {
    document.getElementById('nav-list-div').classList.toggle('display');
    document.getElementById('nav-shelter').classList.toggle('display')
}
