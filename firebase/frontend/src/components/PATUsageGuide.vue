<template>

  <v-card>
    <v-card-title>
      How to use this PAT
    </v-card-title>

    <v-tabs
      v-model="howToUseTab"
    >
      <v-tabs-slider></v-tabs-slider>
      <v-tab>
        Visual Studio
      </v-tab>
      <v-tab>
        Windows CLI
      </v-tab>
      <v-tab>
        Linux CLI
      </v-tab>
    </v-tabs>

    <v-tabs-items v-model="howToUseTab">
      <v-tab-item>
        <v-card flat>
          <v-card-text>
            <p>To configure Visual Studio to access the symbol server using this PAT, visit <code>Tools | Options... | Debugging | Symbols</code>, and add the following symbol file location:
              <code>{{symbolServerDownloadAPIEndpoint}}&nbsp;
                <v-btn
                  icon
                  @click="copyTextToClipboard(symbolServerDownloadAPIEndpoint)"
                >
                  <v-icon
                    small
                  >
                    mdi-content-copy
                  </v-icon>
                </v-btn>
              </code>
            </p>
            <p>If you rather wish to use <code>_NT_SYMBOL_PATH</code> directly, see <a href="https://docs.microsoft.com/en-us/windows/win32/debug/using-symsrv#setting-the-symbol-path">the MSDN documentation</a>.
            </p>
          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item>
        <v-card flat>
          <v-card-text>
            Download the <a href="cloud-symbol-server-cli-win64.exe" download="cloud-symbol-server-cli-win64.exe">Windows CLI tool</a>.
            Also, download <a :href="downloadConfigFileHref" download="cloud-symbol-server-cli.config.json" target="_blank">your personalized config file</a> and place it next to the CLI tool.
          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item>
        <v-card flat>
          <v-card-text>
            Download the <a href="cloud-symbol-server-cli-linux" download="cloud-symbol-server-cli-linux">Linux CLI tool</a>.
            Also, download <a :href="downloadConfigFileHref" download="cloud-symbol-server-cli.config.json" target="_blank">your personalized config file</a> and place it next to the CLI tool.
          </v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs-items>

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
import { adminAPIEndpoint, downloadAPIEndpoint } from '../appConfig'

interface Data {
  howToUseTab: any,
  downloadConfigFileHref: string,
  symbolServerDownloadAPIEndpoint: string,
}

export default Vue.extend({

  props: {
    email: String,
    pat: Object,
  },

  data(): Data {
    return {
      howToUseTab: null,
      downloadConfigFileHref: "data:application/json;charset=utf-8," + encodeURI(JSON.stringify({
        'service-url': adminAPIEndpoint,
        'email': this.email,
        'pat': this.pat.id,
      }, null, 2)),
      symbolServerDownloadAPIEndpoint: downloadAPIEndpoint.split('://')[0] + "://" + encodeURIComponent(this.email) + ':' + this.pat.id + '@' + downloadAPIEndpoint.split('://')[1],
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