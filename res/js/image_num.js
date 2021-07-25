let image_num = get_default_num();
const image_nums = [1, 2, 4];
const display_image_num = {
    "1": "L",
    "2": "M",
    "4": "S"
}

applicate_image_num();

function change_image_num() {
    image_num = image_nums[(image_nums.indexOf(image_num) + 1) % image_nums.length];
    applicate_image_num();
}

function applicate_image_num() {
    document.getElementById("image-num").innerText = display_image_num[image_num.toString()];
    image_nums.map(value => "thumbnail" + value.toString()).forEach(function (value) {
        document.getElementById("content").classList.remove(value);
    });

    if (image_num === 1)
        document.getElementById("content").classList.remove("thumbnail");
    else {
        document.getElementById("content").classList.add("thumbnail");
        document.getElementById("content").classList.add("thumbnail" + image_num.toString());
    }
}

function get_default_num() {
    const width = document.documentElement.clientWidth;
    if (width < 500)
        return 2;
    else
        return 4;
}
