import express from 'express'
import cors from 'cors';

const routes = express.Router()

routes.get('/news', cors(), async (req, res) => {
    // id: faker.datatype.uuid(),
    // title: faker.name.jobTitle(),
    // description: faker.name.jobTitle(),
    // image: `/static/mock-images/covers/cover_${index + 1}.jpg`,
    // postedAt: faker.date.recent(),

    let news = [
        {
            id: "87ea469b-098f-4945-b62d-cf348eac2e9f",
            title: "Lead Response Consultant but from the server ",
            description: "Product Branding Designer",
            image: "/static/mock-images/covers/cover_1.jpg",
            postedAt: "2022-10-17T11:50:02.762Z"
        },
        {
            id: "3551d8b6-9dcb-4089-af96-4a57acfbf9ae",
            title: "Legacy Functionality Producer",
            description: "Dynamic Solutions Facilitator",
            image: "/static/mock-images/covers/cover_2.jpg",
            postedAt: "2022-10-17T07:52:00.428Z"
        },
        {
            id: "a919f152-0981-4b5f-827e-277fadd7d0c7",
            title: "Product Program Consultant",
            description: "Investor Communications Planner",
            image: "/static/mock-images/covers/cover_3.jpg",
            postedAt: "2022-10-17T19:06:52.245Z"
        },
        {
            id: "3af3aa44-e213-4c3c-a589-e922abafe471",
            title: "Lead Configuration Planner",
            description: "International Optimization Director",
            image: "/static/mock-images/covers/cover_4.jpg",
            postedAt: "2022-10-17T18:52:43.629Z"
        },
        {
            id: "2878be35-2a7e-47c2-a385-bbe8ae30fa88",
            title: "Principal Creative Engineer",
            description: "International Interactions Architect",
            image: "/static/mock-images/covers/cover_5.jpg",
            postedAt: "2022-10-17T14:00:51.999Z"
        }
    ]

    res.status(200).json(news)
})

export default routes
