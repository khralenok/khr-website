'use strict'

document.addEventListener("DOMContentLoaded", function(){

    if (localStorage.getItem("token") != null) {
        console.log("You're logged in already")
        window.location.href = "/"
        return
    }

    const form = document.getElementById('auth');

    const login = async function(e) {
        e.preventDefault();

        const formData = new FormData(e.target);
        const email = formData.get("email");
        const password = formData.get("password");

        try{
            const response = await fetch("/login", {
                method: "POST", 
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({
                    "email": email,
                    "pwd": password,
                })
            });

            if (!response.ok){
                throw new Error(`Server error: ${response.status}`);
            }

            const data = await response.json();

            localStorage.setItem("token", data.token);

            window.location.href = "/";
        } catch (error) {
            console.error('Fetch error', error);
        }
    }

    form.addEventListener('submit', login);
});