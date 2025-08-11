'use strict'


/**
 * This function take user's creds and return JWT token for user. 
 */
const getJWT = async function(username, password){
        try{
        const response = await fetch("/login", {
            method: "POST", 
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify({
                "username": username,
                "pwd": password})
        });

        if (!response.ok){
            throw new Error(`Server error: ${response.status}`);
        }

        const data = await response.json();

        return data.token
        } catch (error) {
            console.error('Fetch error', error);
        }
}

/**
 * This function handle login request. 
 */
const login = async function(e) {
    e.preventDefault();

    const formData = new FormData(e.target);
    const username = formData.get("username");
    const password = formData.get("password");

    const token = await getJWT(username, password);

    if (!token) {
        console.error('Can\'t login');
        return;
    };
    
    localStorage.setItem("token", token);
    window.location.href = "/";
}

/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const registerInteracriveElements = function(){
    let token = localStorage.getItem("token")

    if (token != null) {
        window.location.href = "/"
    }

    const form = document.getElementById('auth');

    console.log(form);

    form.addEventListener('submit', login);
}

document.addEventListener("DOMContentLoaded", registerInteracriveElements);