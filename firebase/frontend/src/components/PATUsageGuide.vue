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

<script setup lang="ts">

import { ref } from 'vue'

import { adminAPIEndpoint, downloadAPIEndpoint } from '../appConfig'
import { GetTokenResponse } from '../generated/api'

const props = defineProps<{
  email: string,
  pat: GetTokenResponse,
}>()

const emit = defineEmits<{
  (e: 'done'): void
}>()

const howToUseTab = ref(null)
const downloadConfigFileHref = "data:application/json;charset=utf-8," + encodeURI(JSON.stringify({
        'service-url': adminAPIEndpoint,
        'email': props.email,
        'pat': props.pat.token,
      }, null, 2))
const symbolServerDownloadAPIEndpoint = downloadAPIEndpoint.split('://')[0] + "://" + encodeURIComponent(props.email) + ':' + props.pat.token + '@' + downloadAPIEndpoint.split('://')[1]

function copyTextToClipboard(text: string) {
  navigator.clipboard.writeText(text)
}

function done() {
    emit('done')
}

</script>