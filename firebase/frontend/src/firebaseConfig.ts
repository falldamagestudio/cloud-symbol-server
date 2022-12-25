const firebaseConfig = JSON.parse(process.env.VUE_APP_FIREBASE_CONFIG)

const adminAPIEndpoint = process.env.VUE_APP_ADMIN_API_ENDPOINT
const downloadAPIEndpoint = process.env.VUE_APP_DOWNLOAD_API_ENDPOINT

export { firebaseConfig, adminAPIEndpoint, downloadAPIEndpoint }
