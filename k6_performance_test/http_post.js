import http from 'k6/http';

export default function(){
    const url = 'http://test.k6.io/k6test';

    const reponse = http.get(url);

    check(response, {
        'status is 200': (r) => r.status === 200,
    });
}
