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
    csApiSecret: process.env.LANDING_CS_API_SECRET,
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
    }
}

export default env