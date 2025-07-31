'use srict'

/**
 * Handle new post creation request. 
 * Contain logic for getting data from form which called this function, and async function that sent this data to server.
 */
const deleteThePost = async function(){
    const url = "/post/delete/" + document.getElementById('delete').dataset.postId;

    try{
        const response = await fetch(url, {
                method: "PUT",
        });

        if (!response.ok){
                return response.json().then(errorData => {
                    throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
                }); 
        }

        window.location.href = "/";

        } catch(error) {
            console.error('Fetch error', error)
        }
}

/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const registerInteracriveElements = function(){
    const deleteBtn = document.getElementById('delete')

    deleteBtn.addEventListener('click', deleteThePost)
}

document.addEventListener('DOMContentLoaded', registerInteracriveElements)