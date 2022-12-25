import { GoogleAuthProvider } from 'firebase/auth'

export function googleProvider(): GoogleAuthProvider {
  const provider = new GoogleAuthProvider()

  // Reference: https://firebase.google.com/docs/reference/js/firebase.auth.GoogleAuthProvider#setcustomparameters
  provider.setCustomParameters({

    // Optimize the login process for accounts at a particular domain
    hd: "falldamagestudio.com",

    // Ensure that the account selector is shown, even if the user only has provided a single account so far
    prompt: 'select_account'
  })

  return provider
}
