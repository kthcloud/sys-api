import env from './environment.js'

import express from 'express';
import bodyParser from 'body-parser';

import status from './routes/status.js';
import capacities from './routes/capacities.js';
import news from './routes/news.js';
import stats from './routes/stats.js'

const app = express();

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use('/', status, capacities, news, stats);

const port = env.LANDING_PORT || 8080;

app.listen(port);
console.log('server started on port: ' + port);

