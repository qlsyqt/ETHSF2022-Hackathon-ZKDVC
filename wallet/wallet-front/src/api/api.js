import axios from 'axios';

//Get identity
export async function fetchPolygonId() {
    const resp = await axios.get(`${process.env.REACT_APP_WALLET_BACKEND}/api/identity/fetch`)
    return resp.data
}

//Get polygon address
export async function fetchPolygonWalletInfo() {
    const resp = await axios.get(`${process.env.REACT_APP_WALLET_BACKEND}/api/wallet/fetch`)
    return resp.data
}

//Get polygon address
export async function fetchClaims() {
    const resp = await axios.get(`${process.env.REACT_APP_WALLET_BACKEND}/api/claims`)
    return resp.data
}

//
export async function fetchBadges() {
    const resp = await axios.get(`${process.env.REACT_APP_WALLET_BACKEND}/api/badge/fetch`)
    return resp.data
}


export async function sendAuthProof(authData) {
    const payload = JSON.parse(authData);
    console.log(payload);
    const resp = await axios.post(`${process.env.REACT_APP_WALLET_BACKEND}/api/authenticate`, payload)
    return resp.data
}

export async function acquirerBadge(hIndex) {
    const resp = await axios.post(`${process.env.REACT_APP_WALLET_BACKEND}/api/badge/claim?hIndex=${hIndex}`)
    return resp.data
}
