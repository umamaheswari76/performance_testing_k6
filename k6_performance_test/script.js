import http from 'k6/http';
import { check} from 'k6';


export const options = {
    vus: 1,
    iterations: 100000,
}

export default function(){
    const url = 'http://localhost:8080/k6test';
    // Define the JSON request body
    const payload = JSON.stringify({
        message: 'Hello, world!',
    });
    const params = {};
    
    const response = http.post(url, payload, params);
    check(response, {
        'status is 200': (r) => r.status === 200,
    });
}
