import qs from 'qs'
import crypto from 'crypto'
import url from 'url'
import * as fs from 'fs'

function createCsUrl(baseUrl, command, apiKey, secretKey) {
    let queryDict = {
        'apiKey': apiKey,
        'command': command,
        'response': 'json',
    }

    let hmac = crypto.createHmac('sha1', secretKey);
    let orderedQuery = qs.stringify(queryDict, { encode: true }).replace(/\%5B(\D*?)\%5D/g, '.$1').replace(/\%5B(\d*?)\%5D/g, '[$1]').split('&').sort().join('&').toLowerCase();
    hmac.update(orderedQuery);
    const signature = hmac.digest('base64');

    queryDict['signature'] = signature

    let apiURL = url.parse(baseUrl)
    apiURL.path += qs.stringify(queryDict, { encode: true }).replace(/\%5B(\D*?)\%5D/g, '.$1');

    return `${apiURL.protocol}//${apiURL.host}${apiURL.path}`;

}

function __loadHosts() {
    const path = process.env.LANDING_HOSTS_PATH

    let hosts = []

    if (path) {
        hosts = JSON.parse(fs.readFileSync(process.env.LANDING_HOSTS_PATH))
    }

    console.log('Using hosts: ');
    console.log(hosts);

    return hosts
}

const hosts = __loadHosts()

export { createCsUrl, hosts }