// HTMLが完全に読み込まれたあとに以下の処理が実行される
window.onload = function () {
    bottonClick()
}

function bottonClick() {
    // htmlのbuttonのデータを変数に格納
    const button = document.querySelector('button');
    

    // button要素が押された時の動作を定義
    button.addEventListener('click', () => {
        // htmlのcontentの中の要素を変更している
        fetch('http://localhost:8080')
            .then(response => response.json())
            .then(data => insertTweetHTML(data));
    });
}

function insertTweetHTML(data) {
    // containerのcontentの中の要素を変更している
    const container = document.getElementById('container')
    const insertHtmlName = `<div id="user">${JSON.stringify(data.name)}</div>`
    const insertHtmlContent = `<div id="user"> Click count: ${JSON.stringify(data.message)} </div>`
    container.insertAdjacentHTML('afterbegin', insertHtmlName + insertHtmlContent);
}