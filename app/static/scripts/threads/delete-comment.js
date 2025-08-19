'use srict'

document.addEventListener('DOMContentLoaded', function(){
    const deleteBtn = document.getElementById('delete')
    const backwardUrl = deleteBtn.dataset.postUrl

    const deleteComment = async function(){
        const url = "/comment/delete/" + document.getElementById('delete').dataset.commentId;

        try{
            const response = await fetch(url, {
                    method: "PUT",
            });

            if (!response.ok){
                    return response.json().then(errorData => {
                        throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
                    }); 
            }

            window.location.href = backwardUrl;
            } catch(error) {
                console.error('Fetch error', error)
            }
    }

    deleteBtn.addEventListener('click', deleteComment)
})