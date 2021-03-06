import firebase from 'firebase/app'
import 'firebase/auth'

export function googleProvider(): firebase.auth.GoogleAuthProvider {
  const provider = new firebase.auth.GoogleAuthProvider()

  // Reference: https://firebase.google.com/docs/reference/js/firebase.auth.GoogleAuthProvider#setcustomparameters
  provider.setCustomParameters({

    // Optimize the login process for accounts at a particular domain
    hd: "falldamagestudio.com",

    // Ensure that the account selector is shown, even if the user only has provided a single account so far
    prompt: 'select_account'
  })

  return provider
}
