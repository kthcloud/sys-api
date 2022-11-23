import express from 'express'
import cors from 'cors';
import fetch from 'node-fetch'
import { createCloudstackUrl } from '../common.js'
import env from '../environment.js'
import { hosts } from '../common.js';
import { clearInvalid } from '../lib/utils.js';

const routes = express.Router()

function convertToGB(bytes) {
    return Math.round(bytes / 1073741824);
}

async function getLocalCapacity(ip, port) {
    return fetch(`http://${ip}:${port}/capacities`,)
        .then(res => res.json())
        .catch(err => {
            console.error(`Failed to fetch hosts\' capacities. Details: ${err}`);
            return undefined
        })
}

async function getCloudstackCapacities() {
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
                cpuCore: {
                    used: cpuCapacity.capacityused,
                    total: cpuCapacity.capacitytotal
                }
            }
        })
}

async function getGpuCapacity() {
    // Fetch data from every host
    const capPromises = hosts.map(host => getLocalCapacity(host.ip, host.port))

    // Then wait for all of it
    const capacitiesPerHost = clearInvalid(await Promise.all(capPromises))

    const gpuCount = capacitiesPerHost.reduce((acc, capacites) => acc + capacites.gpu.count, 0)

    return {
        total: gpuCount
    }
}

routes.get('/capacities', cors(), async (req, res) => {
    let cloudStackCapacities = getCloudstackCapacities()
    let gpuCapacity = getGpuCapacity()

    cloudStackCapacities = await cloudStackCapacities
    gpuCapacity = await gpuCapacity

    const result = {
        cpuCore: cloudStackCapacities.cpuCore,
        ram: cloudStackCapacities.ram,
        gpu: gpuCapacity
    }

    res.status(200).json(result)
})

export default routes