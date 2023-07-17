const url = "http://worldtimeapi.org/api/timezone/Europe/Moscow"
const localhost = "http://localhost:5050"

//const button = document.getElementsByClassName("sign_up_button")[0]



async function getData() {
    const response = await fetch(localhost)
    haha = await response.text()
    console.log(haha)
    
}


//я сру


// button.addEventListener("click", (event) => {
//     (async () => {
//         const res = await getData()
//         console.log(res)
//     })();
// })

getData() 




