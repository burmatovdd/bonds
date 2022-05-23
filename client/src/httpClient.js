export function Post(url,sendData){
    return fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(sendData)
    }).then((response) => {
        return  response.json();
    }).then((data) => {
        return data;
    })
}

export function PostWithAuthorization(url,sendData,token){
    return fetch(url, {
        method: 'POST',
        headers: {
            'Authorization': token,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(sendData)
    }).then((response) => {
        return  response.json();
    }).then((data) => {
        return data;
    })
}

export function PostWithoutReturn(url,sendData,token){
    return fetch(url, {
        method: 'POST',
        headers: {
            'Authorization': token,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(sendData)
    })
}
