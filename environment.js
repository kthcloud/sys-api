import dotenv from 'dotenv'

if (process.env.LANDING_ENV_FILE) {
    dotenv.config({ path: process.env.LANDING_ENV_FILE })
}
else {
    dotenv.config({ path: `../.env` })
}

const env = {
    hostsPath: process.env.LANDING_HOSTS_PATH,
    sessionSecret: process.env.LANDING_SESSION_SECRET,
    cloudstack: {
        api: {
            url: process.env.LANDING_CS_BASE_API_URL,
            key: process.env.LANDING_CS_API_KEY,
            secret: process.env.LANDING_CS_API_SECRET
        }
    },
    k8s: {
        sys: {
            url: process.env.LANDING_K8S_SYS_URL,
            certAuthorityData: process.env.LANDING_K8S_SYS_CERT_AUTHORITY_DATA,
            certData: process.env.LANDING_K8S_SYS_CERT_DATA,
            secret: process.env.LANDING_K8S_SYS_SECRET,
        },
        prod: {
            url: process.env.LANDING_K8S_PROD_URL,
            certAuthorityData: process.env.LANDING_K8S_PROD_CERT_AUTHORITY_DATA,
            certData: process.env.LANDING_K8S_PROD_CERT_DATA,
            secret: process.env.LANDING_K8S_PROD_SECRET,

        },
        dev: {
            url: process.env.LANDING_K8S_DEV_URL,
            certAuthorityData: process.env.LANDING_K8S_DEV_CERT_AUTHORITY_DATA,
            certData: process.env.LANDING_K8S_DEV_CERT_DATA,
            secret: process.env.LANDING_K8S_DEV_SECRET,
        }
    },
    db: {
        url: process.env.LANDING_DB_URL,
        name: process.env.LANDING_DB_NAME,
        username: process.env.LANDING_DB_USERNAME,
        password: process.env.LANDING_DB_PASSWORD
    }
}

export default env