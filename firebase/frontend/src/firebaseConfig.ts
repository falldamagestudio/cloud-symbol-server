const firebaseConfig = JSON.parse(process.env.VUE_APP_FIREBASE_CONFIG)

const authEmulatorUrl = process.env.VUE_APP_AUTH_EMULATOR_URL

const firestoreEmulator = (process.env.VUE_APP_FIRESTORE_EMULATOR_PORT ?
  {
    host: 'localhost',
    port: parseInt(process.env.VUE_APP_FIRESTORE_EMULATOR_PORT)
  }
  : null)

const downloadAPIEndpoint = process.env.VUE_APP_DOWNLOAD_API_ENDPOINT

export { firebaseConfig, authEmulatorUrl, firestoreEmulator, downloadAPIEndpoint }
