'use stict'

document.addEventListener("DOMContentLoaded", async function(){
    localStorage.removeItem("token")
    
    try{
        const response = await fetch("/revoke");

        if (!response.ok){
            throw new Error(`Server error: ${response.status}`);
        }

        window.location.href = "/"
    } catch (error) {
        console.error('Fetch error', error);
    }


})