'use srict'

/**
 * Handle post editing request. 
 * Contain logic for getting data from form which called this function, and async function that sent this data to server.
 */
const handleThePost = function(e){
    e.preventDefault()

    const sentData = async function(input) {
        const formData = new FormData();
        formData.append("content", input.get("content"))

        if (input.get("image") && input.get("image").type.startsWith("image/")) {
           formData.append("image", input.get("image"))
        }

        const url = "/post/" + document.getElementById("post-workshop").dataset.postId;

        try{
            const response = await fetch(url, {
                method: "PUT",
                body: formData,
            });

            if (!response.ok){
                return response.json().then(errorData => {
                    throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
                }); 
            }

            const data = await response.json();

            console.log(data)
            window.location.href = "/";
        } catch(error) {
            console.error('Fetch error', error)
        }
    }

    formData = new FormData(document.getElementById('post-workshop'));
    const input = new Map();

    input.set("content", formData.get("post-content"));

    if (formData.get("post-image")) {
        input.set("image", formData.get("post-image"));
    }

    console.log(input)

    sentData(input)
}

/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const registerInteracriveElements = function(){
    const form = document.getElementById('post-workshop')

    form.addEventListener('submit', handleThePost)
}

document.addEventListener('DOMContentLoaded', registerInteracriveElements)