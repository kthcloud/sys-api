import dotenv from 'dotenv'

if (process.env.LANDING_ENV_FILE) {
    dotenv.config({ path: process.env.LANDING_ENV_FILE })
}
else {
    dotenv.config({ path: `../.env` })
}

const env = {
    hostsPath: process.env.LANDING_HOSTS_PATH,
    csBaseApiUrl: process.env.LANDING_CS_BASE_API_URL,
    csApiKey: process.env.LANDING_CS_API_KEY,
    csApiSecret: process.env.LANDING_CS_API_SECRET
}

export default env