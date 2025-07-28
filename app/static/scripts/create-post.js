'use srict'

const handleNewPost = function(e){
    const sentData = async function(content, image) {
        const formData = new FormData();
        formData.append("content", content)
        formData.append("image", image)


        try{
            const response = await fetch("workshop/post", {
                method: "POST",
                body: formData,
            })

            if (!response.ok){
                throw new Error(`Server error: ${response.status}`);   
            }

            const data = await response.json()
            console.log(data)
        } catch(error) {
            console.error('Fetch error', error)
        }
    }

    e.preventDefault()

    formData = new FormData(document.getElementById('post-workshop'));

    const content = formData.get("post-content");
    const image = formData.get("post-image");

    if (!image.type.startsWith("image/")) {
        console.log("Input file should be an image")
        return
    }

    sentData(content, image)
}

const registerTheForm = function(){
    const form = document.getElementById('post-workshop')

    form.addEventListener('submit', handleNewPost)
}

document.addEventListener('DOMContentLoaded', registerTheForm)