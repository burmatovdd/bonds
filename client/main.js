let tmpl = {
    year : null,
}

let bond;
const SendData = async () => {
    tmpl.year = document.getElementById("input").value

    let sendUrl = "http://localhost:8080/result";

    await fetch(sendUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(tmpl)
    }).then((response) => {
        return response.json();
    })
        .then((data) => {
            console.log(data);
            bond = JSON.parse(JSON.stringify(data));
            console.log(bond);
        });

}




