const url = "http://worldtimeapi.org/api/timezone/Europe/Moscow"
const localhost = "http://localhost:3000/"

//const button = document.getElementsByClassName("sign_up_button")[0]



async function getData() {

    let username = document.getElementsByName("username")[0].value
    let password = document.getElementsByName("password")[0].value
    
    //console.log(JSON.stringify({username: password}))

    const response = await fetch(localhost, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({username: username, password: password}),
    });
    haha = await response.text()
    console.log(haha)
    
}





// button.addEventListener("click", (event) => {
//     (async () => {
//         const res = await getData()
//         console.log(res)
//     })();
// })

//getData() 




