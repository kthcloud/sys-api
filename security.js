import Keycloak from 'keycloak-connect';
import { memoryStore } from "./common.js";

const keycloak = new Keycloak({
    store: memoryStore
});

export { keycloak }