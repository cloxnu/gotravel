
const classification_items = {
    all: "all",
    travel: "travel",
    life: "life",
    inspiration: "inspiration"
};
let current_classification = classification_items.all;

const layout_items = {
    grid: "grid",
    list: "list"
};
const cover_style_items = {
    grid: "cover-style-2",
    list: "cover-style-3"
}
let current_layout = layout_items.grid;

function class_btn_clicked(name) {
    if (name === current_classification)
        return ;
    for (let ele of document.getElementsByClassName("class-btn")) {
        ele.classList.remove("active");
    }
    document.getElementById(`class-btn-${name}`).classList.add("active");
    current_classification = name;

    document.getElementById('blog-div').classList.remove(`class-filter-${classification_items.travel}`)
    document.getElementById('blog-div').classList.remove(`class-filter-${classification_items.life}`)
    document.getElementById('blog-div').classList.remove(`class-filter-${classification_items.inspiration}`)

    if (name !== 'all') document.getElementById('blog-div').classList.add(`class-filter-${name}`)
}

function layout_btn_clicked(name) {
    if (name === current_layout)
        return ;
    for (let ele of document.getElementsByClassName("layout-btn")) {
        ele.classList.remove("active");
    }
    document.getElementById(`layout-btn-${name}`).classList.add("active");
    current_layout = name;

    if (name === layout_items.grid) {
        let eles = document.getElementsByClassName("cover-style-3");
        Array.from(eles).forEach(function (ele) {
            ele.classList.remove("cover-style-3");
            ele.classList.add("cover-style-2");
        });
    } else {
        let eles = document.getElementsByClassName("cover-style-2");
        Array.from(eles).forEach(function (ele) {
            ele.classList.remove("cover-style-2");
            ele.classList.add("cover-style-3");
        });
    }
}


window.onload = function () {
    loaded()
}

