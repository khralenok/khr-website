'use strict'

/**
 * Sent request to add a like to specific post
 */
const likePost = async function(postId) {
    const url = "/like/" + postId

    try{
        const response = await fetch(url, {
            method: "POST",
        });

        if (!response.ok){
            return response.json().then(errorData => {
                throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
            });  
        }

        const data = await response.json();
    } catch(error) {
        console.error('Fetch error', error)
    }
}

/**
 * Sent request to remove a like from specific post
 */
const unlikePost = async function(postId) {
    const url = "/like/" + postId

    try{
        const response = await fetch(url, {
            method: "PUT",
        });

        if (response.status === 204){
            console.log('Like was erased')
        }
    } catch(error) {
        console.error('Fetch error', error)
    }
}



/**
 * Sent data about liked post to the server and increase likes counter(fake increasing, real one will take place on page reloading)
 */
const handleLikes = function(e){
    const btn = e.target

    const parts = btn.innerText.trim().split(" ")

    var emoji = parts[0]
    var amount = parseInt(parts[1])

    if (!btn.dataset.isChecked) {
        likePost(btn.dataset.postId)
        emoji = "❤️"
        amount ++
        btn.innerText = emoji + " " + amount
        btn.dataset.isChecked = "true"
    } else if (btn.dataset.isChecked == "true"){
        unlikePost(btn.dataset.postId)
        emoji = "🩶"
        amount --
        delete btn.dataset.isChecked
        btn.innerText = emoji + " " + amount
    }
}


/**
 * Get elements from the page and add corresponsing event listeners to them
 */
const registerLikeBtns = function(){
    const likeBtns = document.querySelectorAll('.like')

    likeBtns.forEach((btn) => {
        btn.addEventListener('click', handleLikes)
    })
}

document.addEventListener('DOMContentLoaded', registerLikeBtns)