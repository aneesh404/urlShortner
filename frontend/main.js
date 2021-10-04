const inputBox = document.querySelector(".input-link");
const submitBtn = document.querySelector("#submit");
const outputBox = document.querySelector(".output-link");

submitBtn.addEventListener("click",function(){
    let url_data = {
        url:inputBox.value
    };
    fetch("/create-smourl",{
        headers: {
            'Accept': 'text/html',
            'Content-Type': 'text/html'
        },
        method: "POST"
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            console.log(result)
        });
    }).catch((error) => {
        console.log(error)
    });
})                                          