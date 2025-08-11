'use strict'

/**
 * This function send creds to server to create a new user and return the new user object
 */
const sentData = async function(username, password){
        try{
        const response = await fetch("/signin", {
            method: "POST", 
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                "username": username,
                "pwd": password})
        })

        if (!response.ok){
            throw new Error(`Server error: ${response.status}`);
        }

        const result = await response.json();

        return result
        } catch (error) {
            console.error('Fetch error', error)
            return null
        }
}

/**
 * This function handle signin form
 */
const signin = async function(e) {
    e.preventDefault();

    const formData = new FormData(e.target);
    const username = formData.get("username");
    const password = formData.get("password");

    

    console.log("password: ", password, "\nusername: ", username)

    const result = await sentData(username, password);

    if (!result) {
        console.error('Can\'t signin')
        return;
    }

    window.location.href = "/login"
}

/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const registerInteracriveElements = function(){
    let token = localStorage.getItem("token")

    if (token != null) {
        window.location.href = "/"
    }

    const form = document.getElementById('auth')

    form.addEventListener('submit', signin)
}

document.addEventListener("DOMContentLoaded", registerInteracriveElements)