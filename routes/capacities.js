import express from 'express'
import cors from 'cors';
import fetch from 'node-fetch'
import { createCloudstackUrl } from '../common.js'
import env from '../environment.js'


const routes = express.Router()

function convertToGB(bytes) {
    return Math.round(bytes / 1073741824);
}

async function getCapacities() {
    const command = 'listCapacity'

    const url = createCloudstackUrl(env.cloudstack.api.url, command, env.cloudstack.api.key, env.cloudstack.api.secret)

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
                cpuCores: {
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