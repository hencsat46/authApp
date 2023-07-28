const localhost = "http://localhost:3000/"

async function loginData() {

    let username = document.getElementsByName("username")[0].value
    let password = document.getElementsByName("password")[0].value

    const response = await fetch(localhost, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({username: username, password: password, type: "Sign in"}),
    });
    haha = await response.text()
    console.log(haha)
    
}