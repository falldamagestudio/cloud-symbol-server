<template>
  <v-card
    :elevation="2"
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
        </v-col>

      </v-row>
    </v-card-text>
    
  </v-card>
</template>

<script lang="ts">

import Vue from 'vue'
import { db } from '../firebase'
import PATUsageGuide from './PATUsageGuide.vue'

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

  methods: {

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