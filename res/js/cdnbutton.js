function cdnButtonOnClick() {
    const cdn = getQueryVariable('cdn');
    if (cdn === 'false') {
        const newURL = "./?" + window.location.search.substring(1).replace(/\&?cdn=false/g, "");
        window.location.href = newURL;
    } else {
        if (cdn) {
            window.location.href = "./?" + window.location.search.substring(1).replace(/cdn=.*($|&)/g, "cdn=false");
        } else {
            window.location.href = "./?" + window.location.search.substring(1) + "&cdn=false";
        }

    }
}

function setCDNLink() {
    const cdn = getQueryVariable('cdn');
    if (cdn === 'false') {
        const newURL = "./?" + window.location.search.substring(1).replace(/\&?cdn=false/g, "");
        document.getElementById("cdn-a").href = newURL;
    } else {
        if (cdn) {
            document.getElementById("cdn-a").href = "./?" + window.location.search.substring(1).replace(/cdn=.*($|&)/g, "cdn=false");
        } else {
            document.getElementById("cdn-a").href = "./?" + window.location.search.substring(1) + "&cdn=false";
        }
    }
}

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
