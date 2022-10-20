import express from 'express'
import fetch from 'node-fetch';
import cors from 'cors';
import { hosts } from '../common.js'

const routes = express.Router()

function getLocalStatus(ip, port) {
  return fetch(`http://${ip}:${port}/status`,)
    .then(res => res.json())
}

routes.get('/status', cors(), async (req, res) => {
  // Fetch data from every host
  const collectedHost = hosts.reduce((acc, elem) => {
    acc.push(getLocalStatus(elem.ip, elem.port))
    return acc
  }, [])

  // Await result and add metadata
  for (let i = 0; i < collectedHost.length; i++) {
    collectedHost[i] = await collectedHost[i]

    // Add Metadata, add more fields if necessary
    collectedHost[i].name = hosts[i].name
  }

  const status = {
    average: {
      cpu: {
        load: collectedHost.reduce((acc, elem) => {
          acc += elem.cpu.load.main
          return acc
        }, 0.0) / collectedHost.length
      }
    },
    hosts: collectedHost
  }

  res.status(200).json(status)
})


export default routes