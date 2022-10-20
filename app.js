import express from 'express';
import bodyParser from 'body-parser';

import statusRoutes from './routes/status.js';
import capcityRoutes from './routes/capacities.js';
import newsRoutes from './routes/news.js';

const app = express();

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use('/', statusRoutes, capcityRoutes, newsRoutes);

const port = process.env.LANDING_PORT || 8080;

app.listen(port);
console.log('server started on port: ' + port);

