<template>
  <v-app>

    <!-- Display main UI when user is logged in -->

    <template v-if="isLoggedIn()">
      <v-app-bar app clipped-left color="primary" dark>

        <div class="d-flex align-center">

          Cloud Symbol Server

        </div>

        <v-spacer></v-spacer>

        <pre>{{version}}  </pre>

        <v-btn v-on:click="logout">Logout</v-btn>

      </v-app-bar>

      <NavigationDrawer/>

      <v-main>
        <v-container fluid>
          <router-view/>
        </v-container>
      </v-main>
    </template>

    <!-- Display a spinner when the app doesn't yet know the user's login status -->

    <template v-else-if="isLoginStateUnknown()">

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

<script setup lang="ts">

import { signInWithRedirect, signOut } from 'firebase/auth'
import { auth } from './firebase'

import { useAuthUserStore, LoginState } from './stores/authUser'
import { googleProvider } from './google-auth'
import { version } from './appConfig'

import NavigationDrawer from './components/NavigationDrawer.vue'

const authUserStore = useAuthUserStore()

function login(): void {
  const provider = googleProvider()
  signInWithRedirect(auth, provider)
}

function logout(): void {
  signOut(auth)
}

function isLoggedIn(): boolean {
   return authUserStore.user != null
}

function isLoginStateUnknown(): boolean {
    return authUserStore.loginState == LoginState.Unknown
}

</script>
