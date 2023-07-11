const url = "http://worldtimeapi.org/api/timezone/Europe/Moscow"
const localhost = "http://localhost:5050"

async function getData() {
    const response = await fetch(localhost)
    console.log(await response.text())
}

getData()




