function onReady() {
    const sayHelloBtn = document.querySelector('#sayHelloBtn')
    sayHelloBtn.addEventListener('click', () => {
        alert('Hola Mundo')
    })
}


(function ready(fn) {
    if (document.attachEvent ? document.readyState === "complete" : document.readyState !== "loading") {
        fn()
    } else {
        document.addEventListener('DOMContentLoaded', fn)
    }
})(onReady)