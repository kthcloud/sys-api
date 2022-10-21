import express from 'express'
import cors from 'cors';
import fetch from 'node-fetch'
import env from '../environment.js'
import { k8sClients } from '../common.js';

const routes = express.Router()

async function getTotalPodCount() {
    const clients = [k8sClients.sys, k8sClients.prod, k8sClients.dev]

    const podsPromises = clients.map(client => client.api.v1.pods.get())

    const pods = await Promise.all(podsPromises)
    const podCount = pods.map(podList => podList.body.items.length)

    return podCount.reduce((acc, podCount) => acc + podCount, 0)
}

async function getKubernetesStats() {
    const podCount = await getTotalPodCount()

    return {
        podCount: podCount
    }
}


routes.get('/stats', cors(), async (req, res) => {
    const k8sStatsPromise = getKubernetesStats()

    const [k8sStats] = await Promise.all([k8sStatsPromise])

    res.status(200).json({
        k8s: k8sStats
    })
})

export default routes