import { defineStore } from 'pinia'

import { User } from 'firebase/auth'

export enum LoginState {
  // When the Firebase SDK initializes, it does initally not know whether it
  //  has a user since the previous session (or whether this is the page reload
  //  just after a sign-in/sign-out operation)
  // Once the SDK completes initialization, it will trigger an onAuthStateChanged()
  //  callback, and provide either a user object or null.
  // We use the login state of LoginState.Unknown to signal to the rest of our application
  //  that it should not yet show any application UI.
  Unknown,
  LoggedOut,
  LoggedIn,
}

interface State {
  user: null | User,
  loginState: LoginState
}

export const useAuthUserStore = defineStore('authUser', {

  state: (): State => ({
    user: null,
    loginState: LoginState.Unknown,
  }),

  actions: {
    setUser (user: null | User) {
      this.user = user
      if (this.user) {
        this.loginState = LoginState.LoggedIn
      } else {
        this.loginState = LoginState.LoggedOut
      }
    },

    setLoginStateUnknown () {
      this.loginState = LoginState.Unknown
    },
  },
})
