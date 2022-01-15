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

        <!-- "Client tools" modal dialog box -->

        <v-dialog
          v-model="clientToolsDialogVisible"
          width="500"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              icon
            >
              <v-icon
                v-bind="attrs"
                v-on="on"
                large
              >
                mdi-console
              </v-icon>
            </v-btn>
          </template>

          <v-card>
            <v-card-title>
              CLI tools
            </v-card-title>

            <v-tabs
              v-model="clientToolsDialogTab"
            >
              <v-tabs-slider></v-tabs-slider>
              <v-tab>
                Windows CLI
              </v-tab>
              <v-tab>
                Linux CLI
              </v-tab>
            </v-tabs>

            <v-tabs-items v-model="clientToolsDialogTab">
              <v-tab-item>
                <v-card flat>
                  <v-card-text>
                    Download the <a href="cloud-symbol-server-cli-win64.exe" download="cloud-symbol-server-cli-win64.exe">Windows CLI tool</a>.
                  </v-card-text>
                </v-card>
              </v-tab-item>
              <v-tab-item>
                <v-card flat>
                  <v-card-text>
                    Download the <a href="cloud-symbol-server-cli-linux" download="cloud-symbol-server-cli-linux">Linux CLI tool</a>.
                  </v-card-text>
                </v-card>
              </v-tab-item>
            </v-tabs-items>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="primary"
                @click="clientToolsDialogVisible = false"
              >
                OK
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>


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
import firebase from 'firebase/app'
import Vue from 'vue';
import store, { LoginState } from './store/index'
import { googleProvider } from './google-auth'
import { version } from './appConfig'

interface Data {
  version: string,
  clientToolsDialogVisible: boolean,
  clientToolsDialogTab: any,
}

export default Vue.extend({
  name: 'App',

  data (): Data {
    return {
      version: version,
      clientToolsDialogVisible: false,
      clientToolsDialogTab: null,
    }
  },

  methods: {
    login (): void {
      const provider = googleProvider()
      firebase.auth().signInWithRedirect(provider)
    },

    logout (): void {
      firebase.auth().signOut()
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
