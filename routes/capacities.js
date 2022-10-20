import express from 'express'
import cors from 'cors';
import fetch from 'node-fetch'
import qs from 'qs'
import crypto from 'crypto'
import url from 'url'

const routes = express.Router()

const LANDING_CS_BASE_API_URL = 'https://dashboard.kthcloud.com/client/api?'
const LANDING_CS_API_KEY = '0wfmb6fRbHewRxui8B2eEEHcIRci6E3sVq-5ErqSqTzcnY2xo6aicOMywalgAH03g-_hy2Al14mGFaLJwW52UQ'
const LANDING_CS_SECRET_KEY = '5YPuTviUMXGcTAvwLOqSbGkc3nZvR_1DlORKMtETi9TThfCdprV6nyfdyxXDei7EqM0EkeoaazfFd_AU2JCE5A'

function createUrl(baseUrl, command, apiKey, secretKey) {
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

function convertToGB(bytes) {
    return Math.round(bytes / 1073741824);
}

async function getCapacities() {
    const command = 'listCapacity'

    const baseUrl = LANDING_CS_BASE_API_URL
    const apiKey = LANDING_CS_API_KEY
    const secret = LANDING_CS_SECRET_KEY

    const url = createUrl(baseUrl, command, apiKey, secret)

    return fetch(url)
        .then(res => res.json())
        .then(result => {
            const capacities = result.listcapacityresponse.capacity

            const cpuCapacity = capacities.find(item => item.name === 'CPU_CORE')
            const memoryCapacity = capacities.find(item => item.name === 'MEMORY')


            return {
                ram: {
                    used: convertToGB(memoryCapacity.capacityused),
                    total: convertToGB(memoryCapacity.capacitytotal)
                },
                cpu: {
                    used: cpuCapacity.capacityused,
                    total: cpuCapacity.capacitytotal
                }
            }
        })
}


routes.get('/capacities', cors(), async (req, res) => {

    let capacities = await getCapacities()

    res.status(200).json(capacities)
})

export default routes