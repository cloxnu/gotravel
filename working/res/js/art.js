function getQueryVariable(v) {
    const query = window.location.search.substring(1);
    const vars = query.split('&');
    for (let i = 0; i < vars.length; i++) {
        const pair = vars[i].split('=');
        if (pair[0] === v) {
            return pair[1];
        }
    }
    return false;
}

let art = getQueryVariable('art');
let cdn = function () {
    const v = getQueryVariable('cdn');
    return v !== 'false';
}();

if (cdn) {
    document.getElementById("cdn-btn").classList.add("green");
    // document.getElementById("cdn-line").classList.add("bg-green");
    document.getElementById("cdn-point").classList.add("bg-green");
    // document.getElementById("cdn-div").classList.add("bg-green");
    document.getElementById("cdn-line").style.opacity = "0";
    document.getElementById("cdn-point").style.opacity = "1";
} else {
    document.getElementById("cdn-btn").classList.add("red");
    document.getElementById("cdn-line").classList.add("bg-red");
    // document.getElementById("cdn-div").classList.add("bg-red");
    document.getElementById("cdn-point").style.opacity = "0";
    document.getElementById("cdn-line").style.opacity = "1";
}

let blog_json;
let blog_title;
let blog_content;
let blog_associated;
let blog_cover;

if (art) {
    read_info(art, function (json) {
        blog_json = json;
        blog_title = blog_json["title"];
        blog_cover = blog_json["cover"].length === 0 ? "" : content(art + "/" + blog_json["cover"]);
        if (cdn) {
            request(content(art + "/" + blog_json["content"]), function (content) {
                blog_content = content;
                load_page();
            }, function (e) {
                console.error(`${art} read failed with cdn turn on: ${e}.`);
            });
        } else {
            read(art + "/" + blog_json["content"], function (content) {
                blog_content = content;
                load_page();
            }, function (e) {
                console.error(`${art} read failed with cdn turn off: ${e}.`);
            });
        }
        if (blog_json["associated"]) {
            read_infos(blog_json["associated"], function (json) {
                blog_associated = json;
                load_associated();
            }, function (e) {
                console.error(`associated art read failed: ${e}.`);
            });
        }
    }, function (e) {
        console.error(`${art} has not found: ${e}.`);
        window.location.href = "404.php";
    });
}


function load_page() {
    const with_cdn = content(art) + "/";
    const without_cdn = "content/" + art + "/";

    const baseURL = function () {
        if (cdn) {
            return with_cdn;
        } else {
            return without_cdn;
        }
    }();

    let heading_count = 0;
    let renderer = {
        html(html) {
            if (html.includes("-cdn")) {
                return html.replace("src=\"", `src=\"${without_cdn}`);
            }
            return html.replace("src=\"", `src=\"${baseURL}`);
        },
        heading(text, level) {
            heading_count ++;
            if (level === 2 || level === 4)
                all_heading.push({text, level, id: `title-anchor-${heading_count}`, pos: 0});
            // return false;
            return `<h${level} id="title-anchor-${heading_count}">${text}</h${level}>`;
        }
    };

    marked.use({ renderer });
    document.getElementById("content").innerHTML = marked(blog_content, {
        baseUrl: baseURL
    });

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

    if (blog_cover.length === 0) {
        document.getElementById("cover-img").style.display = "none";
        document.getElementById("cover-shadow").style.display = "none";
        document.getElementById("head").style.height = "30vh";
        document.getElementById("head").classList.remove("dark-mode-on");
    } else {
        document.getElementById("cover-img").style.opacity = "1";
        document.getElementById("cover-img").src = blog_cover;
    }

    document.getElementById("title").innerText = blog_title;
    document.title = blog_title + " - 一只狗的火星旅行计划";
    document.getElementById("class").innerText = blog_json["class"];

    document.getElementById("story-date").innerText = blog_json["story_date"];

    add_title_link();
    add_touch();

    load_once(load_item.content);
}

function load_associated() {
    document.getElementById("associated-div").style.display = "block";
    document.getElementById("associated-blogs").innerHTML = generate_blogs_html(blog_associated, "all", "cover-style-2");

    add_transition_animation();
}


function load_handler() {
    if (window.location.hash)
        window.location.href = window.location.hash;
}


let has_loaded = {
    content: false,
    page: false
}
const load_item = {
    content: 'content',
    page: 'page'
}
function load_once (which) {
    has_loaded[which] = true;

    if (Object.values(has_loaded).every(Boolean))
        loaded(true, load_handler);
}

window.onload = function () {
    load_once(load_item.page);
}
