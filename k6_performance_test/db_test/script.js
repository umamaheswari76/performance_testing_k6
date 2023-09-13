import http from 'k6/http';
import {check} from 'k6';


export const options = {
    vus : 100,
    iterations : 1000000,
}

export default function(){
    const url = 'http://localhost:8000/db_test'
    const payload = JSON.stringify({Token : 'test token',});
    const response = http.post(url, payload);
    check(response, {
        'status is 200': (r) => r.status === 200,
    });
}