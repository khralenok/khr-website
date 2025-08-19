'use srict'

document.addEventListener('DOMContentLoaded', function(){
    const replies = document.getElementById('replies-box')

    const deleteReply = async function(e){
        if (!e.target.dataset.deleteReply){
            return
        }

        const btn = e.target
        const replyId = parseInt(btn.dataset.replyId.trim())

        console.log("Deletion of reply with ID: \"" + replyId + "\" was requested")

        const url = "/reply/delete/" + replyId;

        try{
            const response = await fetch(url, {
                    method: "PUT",
            });

            if (!response.ok){
                    return response.json().then(errorData => {
                        throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
                    }); 
            }

            window.location.reload();
            } catch(error) {
                console.error('Fetch error', error)
            }
    }

    replies.addEventListener('click', deleteReply)
})