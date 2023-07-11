const url = "http://worldtimeapi.org/api/timezone/Europe/Moscow"
const localhost = "http://localhost:5050"

const button = document.getElementsByClassName("sign_up_button")[0]

async function getData() {
    const response = await fetch(localhost)
    console.log(await response.text())
}

button.addEventListener("click", (event) => {
    getData()
})

//getData() 




