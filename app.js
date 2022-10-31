import env from './environment.js'

import express from 'express';
import session from 'express-session';
import fileupload from 'express-fileupload';
import bodyParser from 'body-parser';

import status from './routes/status.js';
import capacities from './routes/capacities.js';
import news from './routes/news.js';
import stats from './routes/stats.js'
import { memoryStore } from './common.js'
import { keycloak } from './security.js';
import cors from 'cors'

const app = express();

app.use(cors())
app.use(session({
    secret: env.sessionSecret,
    resave: false,
    saveUninitialized: true,
    store: memoryStore
}))
app.use(fileupload());
app.use(keycloak.middleware({ logout: '/'}))
app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use('/', status, capacities, news, stats);

const port = env.LANDING_PORT || 8080;

app.listen(port);
console.log('server started on port: ' + port);

