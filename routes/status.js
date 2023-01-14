import express from "express";
import fetch from "node-fetch";
import cors from "cors";
import { hosts } from "../common.js";
import { clearInvalid } from "../lib/utils.js";

const routes = express.Router();

function getLocalStatus(ip, port) {
  return fetch(`http://${ip}:${port}/status`)
    .then((res) => res.json())
    .catch((err) => {
      console.error(`failed to fetch hosts\' status. Details: ${err}`);
    });
}

routes.get("/status", cors(), async (req, res) => {
  // Fetch data from every host
  const statusPromises = hosts.map((host) =>
    getLocalStatus(host.ip, host.port)
  );

  let statusesPerHost = await Promise.all(statusPromises);

  for (let i = 0; i < statusesPerHost.length; i++) {
    if (statusesPerHost[i]) {
      // Add Metadata, add more fields if necessary
      statusesPerHost[i].name = hosts[i].name;
    }
  }

  statusesPerHost = clearInvalid(statusesPerHost);

  const status = {
    average: {
      cpu: {
        load:
          statusesPerHost.reduce((acc, elem) => {
            acc += elem.cpu.load.main;
            return acc;
          }, 0.0) / statusesPerHost.length,
      },
    },
    hosts: statusesPerHost,
  };

  res.status(200).json(status);
});

export default routes;
