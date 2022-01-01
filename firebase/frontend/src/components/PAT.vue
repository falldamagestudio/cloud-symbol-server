<template>
  <v-hover v-slot:default="{ hover }">
    <v-card
      :elevation="hover ? 6 : 2"
    >

      <v-card-text>
        <v-row>

          <!-- Personal Access Token ID -->

          <v-col>
            {{ pat.id }}
          </v-col>

          <v-spacer/>

          <v-col
            class="text-right"
          >

            <!-- "Use" modal dialog box -->

            <v-dialog
              v-model="useDialogVisible"
              width="500"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  v-bind="attrs"
                  v-on="on"
                >
                  Use
                </v-btn>
              </template>

              <v-card>
                <v-card-title>
                  How to use this PAT
                </v-card-title>

                <v-card-text>
                  <p>To configure Visual Studio to access the symbol store using this PAT, do the following:
                  </p>

                  <ol>
                    <li>
                      Run the following command to open the Environment Variables editor:
                      <code>{{openEnvironmentVariablesCommand}}&nbsp;
                        <v-btn
                          icon
                          @click="copyTextToClipboard(openEnvironmentVariablesCommand)"
                        >
                          <v-icon
                            small
                          >
                            mdi-content-copy
                          </v-icon>
                        </v-btn>
                      </code>
                    </li>
                    <li>
                      <p>Create a new Environment Variable for the current user with the following name:
                      <code>{{ntSymbolPathName}}&nbsp;
                        <v-btn
                          icon
                          @click="copyTextToClipboard(ntSymbolPathName)"
                        >
                          <v-icon
                            small
                          >
                            mdi-content-copy
                          </v-icon>
                        </v-btn>
                      </code>
                      </p>
                      <p>If this variable already exists, be careful to merge the changes below with the previous settings.
                      </p>
                    </li>
                    <li>
                      Set the value of the variable to the following:
                      <code>{{ntSymbolPathValue}}&nbsp;
                        <v-btn
                          icon
                          @click="copyTextToClipboard(ntSymbolPathValue)"
                        >
                          <v-icon
                            small
                          >
                            mdi-content-copy
                          </v-icon>
                        </v-btn>
                      </code>
                    </li>
                    <li>
                      If the local symbols folder mentioned in the previous doesn't exist on your machine, create it.
                    </li>
                  </ol>
                </v-card-text>

                <v-divider></v-divider>

                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn
                    color="primary"
                    @click="useDialogVisible = false"
                  >
                    OK
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>

            <!-- "Revoke" button -->

            <v-btn
              color="error--text"
              v-on:click="revoke()"
            >
              Revoke
            </v-btn>
          </v-col>

        </v-row>
      </v-card-text>
      
    </v-card>
  </v-hover>
</template>

<script lang="ts">

import Vue from 'vue'
import { db } from '../firebase'
import { downloadAPIProtocol, downloadAPIHost } from '../firebaseConfig'

interface Data {
  useDialogVisible: boolean,
  openEnvironmentVariablesCommand: string,
  ntSymbolPathName: string,
  ntSymbolPathValue: string,
}

export default Vue.extend({

  props: {
    email: String,
    pat: Object,
  },

  data(): Data {
    return {
      useDialogVisible: false,
      openEnvironmentVariablesCommand: 'rundll32 sysdm.cpl,EditEnvironmentVariables',
      ntSymbolPathName: '_NT_SYMBOL_PATH',
      ntSymbolPathValue: 'SRV*C:\\Symbols*' + downloadAPIProtocol + '://' + encodeURIComponent(this.email) + ':' + this.pat.id + '@' + downloadAPIHost,
    }
  },

  methods: {
    copyTextToClipboard(text: string) {
      navigator.clipboard.writeText(text)
    },

    revoke() {
      db.collection('users').doc(this.email).collection('pats').doc(this.pat.id).delete().then(() => {
        this.$emit('refresh')
      })
    }
  }

})

</script>