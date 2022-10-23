import express from 'express'
import cors from 'cors';
import fetch from 'node-fetch'
import { createCloudstackUrl } from '../common.js'
import env from '../environment.js'
import { hosts } from '../common.js';

const routes = express.Router()

function convertToGB(bytes) {
    return Math.round(bytes / 1073741824);
}

function getLocalCapacity(ip, port) {
    return fetch(`http://${ip}:${port}/capacities`,)
        .then(res => res.json())
        .then(capacities => capacities.gpu.count)
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
                cpuCores: {
                    used: cpuCapacity.capacityused,
                    total: cpuCapacity.capacitytotal
                }
            }
        })
}

async function getGpuCapacity() {
    // Fetch data from every host
    const gpuCountPromises = hosts.reduce((acc, elem) => acc + getLocalCapacity(elem.ip, elem.port), 0)
    const gpuCount = Promise.all(gpuCountPromises)

    return {
        count: gpuCount
    }
}

routes.get('/capacities', cors(), async (req, res) => {


    let cloudStackCapacities = getCloudstackCapacities()
    let gpuCapacity = getGpuCapacity()


    cloudStackCapacities = await cloudStackCapacities
    gpuCapacity = await gpuCapacity

    const result = {
        cpu: cloudStackCapacities.cpu,
        ram: cloudStackCapacities.ram,
        gpu: gpuCapacity
    }

    res.status(200).json(cloudStackCapacities)
})

export default routes