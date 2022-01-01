<template>
  <v-hover v-slot:default="{ hover }">
    <v-card
      :elevation="hover ? 6 : 2"
    >

      <v-card-text>
        <v-row>

          <v-col>
            {{ pat.id }}
          </v-col>

          <v-spacer/>

          <v-col
            class="text-right"
          >
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

export default Vue.extend({

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