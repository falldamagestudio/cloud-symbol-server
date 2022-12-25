<template>
  <v-app>

    <!-- Display main UI when user is logged in -->

    <template v-if="isLoggedIn">
      <v-app-bar app color="primary" dark>

        <div class="d-flex align-center">

          Cloud Symbol Server

        </div>

        <v-spacer></v-spacer>

        <pre>{{version}}  </pre>

        <v-btn v-on:click="logout">Logout</v-btn>

      </v-app-bar>

      <v-main>
        <v-container fluid>
          <router-view/>
        </v-container>
      </v-main>
    </template>

    <!-- Display a spinner when the app doesn't yet know the user's login status -->

    <template v-else-if="isLoginStateUnknown">

      <v-container fill-height fluid>
        <v-row align="center" justify="center">
          <v-progress-circular indeterminate color="primary"></v-progress-circular>
        </v-row>
      </v-container>

    </template>

    <!-- Display a login prompt when user is logged out -->

    <template v-else>

      <v-container fill-height fluid>
        <v-row align="center" justify="center">
          <v-btn color="primary" v-on:click="login">Login</v-btn>
        </v-row>
      </v-container>

    </template>

  </v-app>
</template>

<script lang="ts">

import Vue from 'vue';

import { signInWithRedirect, signOut } from 'firebase/auth'

import { auth } from './firebase'
import store, { LoginState } from './store/index'
import { googleProvider } from './google-auth'
import { version } from './appConfig'

interface Data {
  version: string,
}

export default Vue.extend({
  name: 'App',

  data (): Data {
    return {
      version: version,
    }
  },

  methods: {
    login (): void {
      const provider = googleProvider()
      signInWithRedirect(auth, provider)
    },

    logout (): void {
      signOut(auth)
    },
  },

  computed: {
    isLoggedIn(): boolean {
      return store.state.user != null
    },

    isLoginStateUnknown(): boolean {
      return store.state.loginState == LoginState.Unknown
    },
}

});
</script>
