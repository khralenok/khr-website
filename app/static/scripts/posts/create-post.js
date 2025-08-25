'use srict'

document.addEventListener('DOMContentLoaded', function(){
    const form = document.getElementById('workshop')
    const attachmentTypeHandler = document.getElementById('attachement-type')

    const newPost = async function(e){
        e.preventDefault()

        const formData = new FormData(e.target);

        for (const pair of formData.entries()) {
            console.log(pair[0] + " : " + pair[1]);
        }
                
/*         try{
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

            window.location.href = "/";
        } catch(error) {
            console.error('Fetch error', error)
        } */
    }

    const attachmentInputHandler = function(){
        optPart = document.getElementById('optional-part')

        if (optPart){
            optPart.remove()
        }

        let markup =""

        switch(this.value){
            case "image":
            markup = `
                <div id="optional-part" class="mol v">
                    <label for="image">Select an image:</label>
                    <input type="file" id="image" name="image" accept="image/*"/>
                </div>`
            break

            case "carousel":
            markup = `
                <div id="optional-part" class="mol v">
                    <label for="images">Select images:</label>
                    <input type="file" id="images" name="images" accept="image/*" multiple/>
                </div>`
            break

            case "youtube":
            markup = `
                <div id="optional-part" class="mol v">
                    <label for="video-url">Input Youtube video URL:</label>
                    <input type="url" id="video-url" name="video-url" />
                </div>`
            break    
        }

        attachmentTypeHandler.insertAdjacentHTML('afterend', markup)
    }

    form.addEventListener('submit', newPost)
    attachmentTypeHandler.addEventListener('change', attachmentInputHandler)
})