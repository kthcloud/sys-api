import qs from 'qs'
import crypto from 'crypto'
import url from 'url'
import * as fs from 'fs'
import k8sClient from 'kubernetes-client'
import Request from 'kubernetes-client/backends/request/index.js'
import mongoDb, { ServerApiVersion } from 'mongodb'

const { MongoClient } = mongoDb
import session from 'express-session'

import env from './environment.js'

function createCloudstackUrl (baseUrl, command, apiKey, secretKey) {
  let queryDict = {
    'apiKey': apiKey,
    'command': command,
    'response': 'json',
  }

  let hmac = crypto.createHmac('sha1', secretKey)
  let orderedQuery = qs.stringify(queryDict, { encode: true })
    .replace(/\%5B(\D*?)\%5D/g, '.$1')
    .replace(/\%5B(\d*?)\%5D/g, '[$1]')
    .split('&')
    .sort()
    .join('&')
    .toLowerCase()
  hmac.update(orderedQuery)
  const signature = hmac.digest('base64')

  queryDict['signature'] = signature

  let apiURL = url.parse(baseUrl)
  apiURL.path += qs.stringify(queryDict, { encode: true })
    .replace(/\%5B(\D*?)\%5D/g, '.$1')

  return `${apiURL.protocol}//${apiURL.host}${apiURL.path}`

}

function __loadHosts () {
  const path = process.env.LANDING_HOSTS_PATH

  let hosts = []

  if (path) {
    hosts = JSON.parse(fs.readFileSync(process.env.LANDING_HOSTS_PATH))
  }

  console.log('Using hosts: ')
  console.log(hosts)

  return hosts
}

function __createConfig (url, authorityData, certificateData, secret) {
  const configJson = {
    apiVersion: 'v1',
    clusters: [
      {
        cluster: {
          'insecure-skip-tls-verify': true,
          'certificate-authority-data': authorityData,
          server: url,
        },
        name: 'kubernetes',
      },
    ],
    contexts: [
      {
        context: {
          cluster: 'kubernetes',
          user: 'kubernetes-admin',
        },
        name: 'kubernetes-admin@kubernetes',
      },
    ],
    'current-context': 'kubernetes-admin@kubernetes',
    kind: 'Config',
    users: [
      {
        name: 'kubernetes-admin',
        user: {
          'client-certificate-data': certificateData,
          'client-key-data': secret,
        },
      },
    ],
  }
  const kubeconfig = new k8sClient.KubeConfig()
  kubeconfig.loadFromString(JSON.stringify(configJson))
  return kubeconfig
}

function __createClient (kubeconfig) {
  const backend = new Request({ kubeconfig, insecureSkipTlsVerify: true })
  return new k8sClient.Client({ backend, version: '1.13' })
}

async function __connectDb (connectionString, dbName) {
  const client = new MongoClient(connectionString)

  return client
    .connect()
    .then(connectedClient => {
      console.log(`Sucessfully connected to database`)
      return connectedClient.db(dbName)
    })
    .catch(err => {
      console.error(`Failed to connect to database. Details: ${err}`)
      return null
    })
}

const hosts = __loadHosts()

const k8sConfigs = {
  sys: __createConfig(env.k8s.sys.url, env.k8s.sys.certAuthorityData,
    env.k8s.sys.certData, env.k8s.sys.secret),
  prod: __createConfig(env.k8s.prod.url, env.k8s.prod.certAuthorityData,
    env.k8s.prod.certData, env.k8s.prod.secret),
  dev: __createConfig(env.k8s.dev.url, env.k8s.dev.certAuthorityData,
    env.k8s.dev.certData, env.k8s.dev.secret),
}

const __internalk8sClients = {
  sys: __createClient(k8sConfigs.sys),
  prod: __createClient(k8sConfigs.prod),
  dev: __createClient(k8sConfigs.dev),
}

async function __getAvailableK8sClients () {
  const clients = []

  if (__internalk8sClients.sys) {
    const result = await __internalk8sClients.sys.api.v1.namespaces.get()
      .then(_ => true)
      .catch(_ => false)

    if (result) {
      clients.push(__internalk8sClients.sys)
    }
  }

  if (__internalk8sClients.prod) {
    const result = await __internalk8sClients.prod.api.v1.namespaces.get()
      .then(_ => true)
      .catch(_ => false)

    if (result) {
      clients.push(__internalk8sClients.prod)
    }
  }

  if (__internalk8sClients.dev) {
    const result = await __internalk8sClients.dev.api.v1.namespaces.get()
      .then(_ => true)
      .catch(_ => false)

    if (result) {
      clients.push(__internalk8sClients.dev)
    }
  }

  return clients
}

const k8sClients = await __getAvailableK8sClients()

const db = await __connectDb(env.db.url, env.db.name)

const memoryStore = new session.MemoryStore()

export { createCloudstackUrl, hosts, k8sClients, db, memoryStore }