'use srict'

document.addEventListener('DOMContentLoaded', function(){
    const form = document.getElementById('workshop')

    const editPost = async function(e){
        e.preventDefault()
        const formData = new FormData(e.target);
        const url = "/post/" + e.target.dataset.postId;

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

    form.addEventListener('submit', editPost)
})