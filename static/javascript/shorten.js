const URL = document.querySelector('#copyText')
URL.href = URL.textContent




function copy() {
  navigator.clipboard.writeText(URL.href)
    .then(() => alert('已經複製縮短後的網誌 !'))
    .catch(error => console.log(error))
}