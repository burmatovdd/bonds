import *as storage from "./storage";

function request(method,url,data){
    return fetch(url,{
        method: method,
        headers:{
            'Authorization': storage.getToken(),
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then((response) => {
        if (response.headers.has('Content-Type') && response.headers.get('Content-Type').indexOf('application/json') !== -1){
            return  response.json();
        }else{
            return response.body;
        }
    })
}

export function Post(url,sendData){
    return request('POST', url,sendData);
}
