import express from 'express'
import cors from 'cors';
import fetch from 'node-fetch'
import { createCsUrl } from '../common.js'
import env from '../environment.js'

const routes = express.Router()

async function getKubernetesStats() {
    const command = 'listKubernetesClusters'

    const url = createCsUrl(env.csBaseApiUrl, command, env.csApiKey, env.csApiSecret)

    return fetch(url)
        .then(res => res.json())
        .then(result => {
            console.log(result.listkubernetesclustersresponse);
        })
}


routes.get('/stats', cors(), async (req, res) => {

    await getKubernetesStats()

    res.status(200).json({ msg: 'ok' })
})

export default routes