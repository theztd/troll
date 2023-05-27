import http from 'k6/http';
import { group, sleep } from 'k6';
export const options = {"stages":[{"duration":"300s","target":25},{"duration":"15s","target":0}],"thresholds":{"http_req_failed":["rate\u003c0.01"],"http_req_duration":["p(99) \u003c 1000"]}};
export default () => {

    group("pomalu", (_) => {
        http.get(`http://debug-troll.fejk.net/v1/pomalu/`, {"headers":{"X-Auth-Key":""}});

    })

    group("hodne-pomalu", (_) => {
        http.get(`http://debug-troll.fejk.net/v1/hodne-pomalu/?wait=500`, {"headers":{"X-Auth-Key":""}}).json();
        http.get(`http://debug-troll.fejk.net/v1/hodne-pomalu/?wait=10`, {"headers":{"X-Auth-Key":""}});
    })


    http.get(`http://debug-troll.fejk.net/v1/hodne-moc-pomalu/?wait=5000`, {"headers":{"X-Auth-Key":""}}).json();
    http.get(`http://debug-troll.fejk.net/v1/hodne-moc-pomalu/?wait=1000`, {"headers":{"X-Auth-Key":""}});

    sleep(0.1);
};

