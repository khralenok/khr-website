'use srict'

/**
 * Handle new comment creation request. 
 * Contain logic for getting data from the form which called this function, and async function that sent this data to server.
 */
const handleNewReply = function(e){
    e.preventDefault()

    const sentData = async function(input) {
        const url = "/reply/" + input.get("comment_id"); 

        try{
            const response = await fetch(url, {
                method: "POST",
                body: JSON.stringify({
                    "content": input.get("content"),
                }),
            });

            if (!response.ok){
                return response.json().then(errorData => {
                    throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
                });  
            }

            const data = await response.json();
            console.log(data)
            window.location.href = "/comment/" + input.get("comment_id");
        } catch(error) {
            console.error('Fetch error', error)
        }
    }

    formData = new FormData(document.getElementById('workshop'));
    const urlParams =  new URLSearchParams(window.location.search);
    const input = new Map();

    input.set("content", formData.get("content"));
    input.set("comment_id", urlParams.get("comment_id"))

    sentData(input)
}

/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const registerInteracriveElements = function(){
    const form = document.getElementById('workshop')

    form.addEventListener('submit', handleNewReply)
}

document.addEventListener('DOMContentLoaded', registerInteracriveElements)