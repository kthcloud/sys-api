import express from 'express'
import * as fs from 'fs'
import fetch from 'node-fetch';
import cors from 'cors';

const routes = express.Router()

function loadHosts() {
  const path = process.env.LANDING_HOSTS_PATH

  let hosts = []

  if (path) {
    hosts = JSON.parse(fs.readFileSync(process.env.LANDING_HOSTS_PATH))
  }

  console.log('Using hosts: ');
  console.log(hosts);

  return hosts
}

const hosts = loadHosts()

function getLocalStatus(ip, port) {
  return fetch(`http://${ip}:${port}/status`,)
    .then(res => res.json())
}

routes.get('/status', cors(), async (req, res) => {
  // Fetch data from every host
  const all = hosts.reduce((acc, elem) => {
    acc.push(
      getLocalStatus(elem.ip, elem.port)
    )
    return acc
  }, [])

  // Await result and add metadata
  for (let i = 0; i < all.length; i++) {
    all[i] = await all[i]

    // Add Metadata, add more fields if necessary
    all[i].name = hosts[i].name
  }

  res.status(200).json(all)
})


export default routes