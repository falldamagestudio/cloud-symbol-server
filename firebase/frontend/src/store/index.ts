// v9 compat packages are API compatible with v8 code
import firebase from 'firebase/compat/app'

import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from "vuex-persistedstate"

Vue.use(Vuex)

export enum LoginState {
  Unknown,
  LoggedOut,
  LoggedIn,
}

export default new Vuex.Store({
  plugins: [createPersistedState()],

  state: {
    user: null as null | firebase.User,
    loginState: LoginState.LoggedOut,
  },
  getters: {
    getUser: state => {
        return state.user;
    }
  },
  mutations: {
    setUser (state, user: null | firebase.User) {
      state.user = user
      if (state.user) {
        state.loginState = LoginState.LoggedIn
      } else {
        state.loginState = LoginState.LoggedOut
      }
    },

    setLoginStateUnknown (state) {
      state.loginState = LoginState.Unknown
    },
  },
  actions: {

  },
  modules: {
  }
})
