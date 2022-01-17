<template>

  <v-card>
    <v-card-title>
      How to use this PAT
    </v-card-title>

    <v-card-text>
      <p>To configure Visual Studio to access the symbol server using this PAT, do the following:
      </p>

      <ol>
        <li>
          Decide on a location to cache symbol files on your machine, for example
          <code>{{symbolCacheLocation}}&nbsp;
            <v-btn
              icon
              @click="copyTextToClipboard(symbolCacheLocation)"
            >
              <v-icon
                small
              >
                mdi-content-copy
              </v-icon>
            </v-btn>
          </code>
          . Create the folder if it does not exist.
        </li>
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
          Create or edit an environment variable for the current user with the following name:
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
        </li>
        <li>
          <p>Set the value of the variable to the following:
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
          </p>
          <p>If this variable already exists, be careful to merge the changes above with its previous settings.
          </p>
          <p>Adjust the symbol cache folder location if necessary.
          </p>
        </li>
      </ol>
      <p>For more information on _NT_SYMBOL_PATH, see <a href="https://docs.microsoft.com/en-us/windows/win32/debug/using-symsrv#setting-the-symbol-path">the MSDN documentation</a>.
      </p>
    </v-card-text>

    <v-divider></v-divider>

    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn
        color="primary"
        @click="done()"
      >
        OK
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">

import Vue from 'vue'
import { downloadAPIEndpoint } from '../firebaseConfig'

interface Data {
  symbolCacheLocation: string,
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
      symbolCacheLocation: 'C:\\Symbols',
      openEnvironmentVariablesCommand: 'rundll32 sysdm.cpl,EditEnvironmentVariables',
      ntSymbolPathName: '_NT_SYMBOL_PATH',
      ntSymbolPathValue: 'SRV*C:\\Symbols*' + downloadAPIEndpoint.split('://')[0] + "://" + encodeURIComponent(this.email) + ':' + this.pat.id + '@' + downloadAPIEndpoint.split('://')[1],
    }
  },

  methods: {
    copyTextToClipboard(text: string) {
      navigator.clipboard.writeText(text)
    },

    done() {
      this.$emit('done')
    },
  }

})

</script>