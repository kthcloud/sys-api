import express from 'express'
import cors from 'cors';
import { db } from '../common.js';
import { body, param, validationResult } from 'express-validator';
import { v4 as uuid } from 'uuid';
import { keycloak } from '../security.js';

const routes = express.Router()

const newsCollection = db.collection('news')

routes.get(
    '/news',
    cors(),
    async (_req, res) => {
        newsCollection
            .find({}, { projection: { _id: 0 } })
            .toArray()
            .then(news => res.status(200).json(news))
            .catch(err => { res.status(500).json({ msg: `Failed to get news from database. Details: ${err}` }) })
    })

routes.post(
    '/news',
    keycloak.protect(),
    body('title').isString().isLength({ min: 1, max: 500 }),
    body('description').isString().isLength({ max: 5000 }),
    body('image'),
    cors(),
    async (req, res) => {
        const errors = validationResult(req);
        if (!errors.isEmpty()) {
            return res.status(400).json({ errors: errors.array() });
        }

        const body = req.body

        const newsPiece = {
            id: uuid(),
            title: body.title,
            description: body.description,
            image: req.files ? req.files.image.data : null,
            postedAt: new Date(new Date().toUTCString())
        }

        newsCollection
            .insertOne(newsPiece)
            .then(res.status(201).json({ id: newsPiece.id }))
            .catch(err => { res.status(500).json({ msg: `Failed to insert item in database. Details: ${err}` }) })
    })

routes.delete(
    '/news/:id',
    keycloak.protect(),
    param('id').isUUID(4),
    cors(),
    async (req, res) => {
        const errors = validationResult(req);
        if (!errors.isEmpty()) {
            return res.status(400).json({ errors: errors.array() });
        }

        const deleteQuery = { id: req.params.id }
        newsCollection
            .deleteOne(deleteQuery)
            .then(res.sendStatus(204))
            .catch(err => { res.status(500).json({ msg: `Failed to delete item in database. Details: ${err}` }) })
    })

export default routes
