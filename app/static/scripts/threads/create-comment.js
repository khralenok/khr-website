'use srict'

const newComment = function(){
    const form = document.getElementById('workshop')

    form.addEventListener('submit', function(e){
        e.preventDefault()

        const sentData = async function(input) {
            const url = "/comment/" + input.get("post_id"); 

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
                window.location.href = "/post/" + input.get("post_id");
            } catch(error) {
                console.error('Fetch error', error)
            }
        }

        formData = new FormData(document.getElementById('workshop'));
        const urlParams =  new URLSearchParams(window.location.search);
        const input = new Map();

        input.set("content", formData.get("content"));
        input.set("post_id", urlParams.get("post_id"))

        sentData(input)
    })
}

document.addEventListener('DOMContentLoaded', newComment)