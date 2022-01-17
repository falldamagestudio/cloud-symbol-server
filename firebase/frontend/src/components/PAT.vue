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

          <PATUsageGuide
            :email="email"
            :pat="pat"
          />

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

export default Vue.extend({

  components: {
    PATUsageGuide,
  },

  props: {
    email: String,
    pat: Object,
  },

  methods: {

    revoke() {
      db.collection('users').doc(this.email).collection('pats').doc(this.pat.id).delete().then(() => {
        this.$emit('refresh')
      })
    }
  }

})

</script>