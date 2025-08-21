'use strict'

document.addEventListener("DOMContentLoaded", function(){
    const token = localStorage.getItem("token")

    if (token != null) {
        window.location.href = "/"
    }

    const form = document.getElementById('auth')

    const signin = async function(e) {
        e.preventDefault();

        const formData = new FormData(e.target);
        const email = formData.get("email");
        const displayName = formData.get("display-name");
        const password = formData.get("password");

        try{
        const response = await fetch("/signin", {
            method: "POST", 
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                "email": email,
                "display_name": displayName,
                "pwd": password,
            })
        })

        if (!response.ok){
            throw new Error(`Server error: ${response.status}`);
        }

        const result = await response.json();

        window.location.href = "/login"
        } catch (error) {
            console.error('Fetch error', error)
            return null
        }
    }

    form.addEventListener('submit', signin)
})