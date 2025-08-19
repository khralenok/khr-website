'use srict'

document.addEventListener('DOMContentLoaded', function(){
    const form = document.getElementById('workshop')
    const url = "/comment/" + form.dataset.commentId

    const editComment = async function(e){
        e.preventDefault()

        const input = new FormData(form);

        try{
            const response = await fetch(url, {
                method: "PUT",
                body: JSON.stringify({
                    "content": input.get("content").trim(), 
                }),
            });

            if (!response.ok){
                return response.json().then(errorData => {
                    throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
                }); 
            }

            const data = await response.json();

            console.log(data)
            window.location.href = url;
        } catch(error) {
            console.error('Fetch error', error)
        }
    }

    form.addEventListener('submit', editComment)
})