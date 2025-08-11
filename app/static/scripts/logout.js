'use stict'

/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const logout = function(){
    localStorage.removeItem("token")
    window.location.href = "/login"
}

document.addEventListener("DOMContentLoaded", logout)