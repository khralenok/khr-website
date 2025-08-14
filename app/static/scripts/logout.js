'use stict'

const revokeSession = async function(){
        try{
        const response = await fetch("/revoke");

        if (!response.ok){
            throw new Error(`Server error: ${response.status}`);
        }

        const data = await response.json();

        return data
        } catch (error) {
            console.error('Fetch error', error);
        }
}

/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const logout = async function(){
    localStorage.removeItem("token")
    
    data = await revokeSession()

    if (data.error){
        console.log("Can't log out")
        console.log(data.error)
        return
    }

    console.log(data)


    //window.location.href = "/login"
}

document.addEventListener("DOMContentLoaded", logout)