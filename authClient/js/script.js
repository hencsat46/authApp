
let btn = document.getElementsByClassName("sign_up_button")[0]


btn.addEventListener("click", (event) => {
    const response = fetch("http://127.0.0.1:5050")
    console.log(response)
    
})




