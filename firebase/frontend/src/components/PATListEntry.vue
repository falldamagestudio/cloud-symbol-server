<template>

  <tr>
    <!-- Personal Access Token ID -->

    <td>
      {{ pat.id | abbreviateId }}

      <!-- Shortcut for copying full PAT ID to clipboard -->
      <v-btn
        icon
        @click="copyTextToClipboard(pat.id)"
      >
        <v-icon
          small
        >
          mdi-content-copy
        </v-icon>
      </v-btn>
    </td>

    <!-- Personal Access Token creation timestamp -->

    <td>
      {{ pat.get('creationTimestamp') | timestampToDisplayString }}
    </td>

    <!-- Personal Access Token description -->

    <td>
      {{ pat.get('description') }}
    </td>

    <td>
      <div class="text-right">

        <!-- "Use" modal dialog box -->

        <v-dialog
          v-model="useDialogVisible"
          width="1000"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              v-bind="attrs"
              v-on="on"
            >
              Use
            </v-btn>
          </template>

          <PATUsageGuide
            :email="email"
            :pat="pat"
            @done="patUsageGuideDone"
          />

        </v-dialog>

        <!-- "Revoke" button -->

        <v-btn
          color="error--text"
          v-on:click="revoke()"
        >
          Revoke
        </v-btn>
      </div>

    </td>

  </tr>

</template>

<script lang="ts">

import Vue from 'vue'
import firebase from 'firebase/app'
import { db } from '../firebase'
import PATUsageGuide from './PATUsageGuide.vue'

import dayjs from 'dayjs'


interface Data {
  useDialogVisible: boolean,
}

export default Vue.extend({

  components: {
    PATUsageGuide,
  },

  props: {
    email: String,
    pat: Object,
  },

  data (): Data {
    return {
      useDialogVisible: false,
    }
  },

  filters: {
    
    abbreviateId: function (id: string): string {
      return `${id.slice(0, 4)}...${id.slice(-4)}`;
    },

    timestampToDisplayString: function (timestamp: firebase.firestore.Timestamp): string {
      return dayjs(timestamp.toDate()).format('YYYY-MM-DD HH:mm');
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
    },

    patUsageGuideDone() {
      this.useDialogVisible = false
    },
  }

})

</script>