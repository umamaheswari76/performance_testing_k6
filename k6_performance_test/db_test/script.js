import http from 'k6/http';
import {check} from 'k6';


const options = {
    vus : 1,
    iterations : 1
}

export default function(){
    const url = 'http://localhost:8000/db_test'
    const payload = JSON.stringify({Token : 'test token',});
    const response = http.post(url, payload);
    check(response, {
        'status is 200': (r) => r.status === 200,
    });
}