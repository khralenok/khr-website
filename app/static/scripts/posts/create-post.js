'use srict'

document.addEventListener('DOMContentLoaded', function(){
    const form = document.getElementById('workshop')

    const newPost = async function(e){
        e.preventDefault()

        const formData = new FormData(e.target);
                
        try{
            const response = await fetch("/post", {
                method: "POST",
                body: formData,
            });

            if (!response.ok){
                return response.json().then(errorData => {
                    throw new Error(`Server error: ${response.status} - ${errorData.message || response.statusText}`);
                });  
            }

            const data = await response.json();

            //console.log(data)

            window.location.href = "/";
        } catch(error) {
            console.error('Fetch error', error)
        }
    }

    form.addEventListener('submit', newPost)
})