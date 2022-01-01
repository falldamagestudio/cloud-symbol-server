<template>
  <div>

    <!-- Header -->

    <v-row>

      <v-col
      >
        Your tokens
      </v-col>

      <v-col
        class="text-right"
      >

        <v-btn
        >
          <v-icon>mdi-plus</v-icon>
          Generate new token
        </v-btn>

        <v-btn
          color="error--text"
        >
          Revoke all tokens
        </v-btn>

      </v-col>

    </v-row>

    <!-- Existing tokens -->

    <v-row>
      <template v-for="pat in pats">
        <v-col v-bind:key="pat.id" cols="12">

          <PAT :pat="pat"/>

        </v-col>
      </template>
    </v-row>
  </div>
</template>

<script lang="ts">

import Vue from 'vue'
import type firebase from 'firebase'
import { db } from '../firebase'
import PAT from './PAT.vue'

interface Data {
  pats: firebase.firestore.QueryDocumentSnapshot<firebase.firestore.DocumentData>[]
}

export default Vue.extend({

  components: {
    PAT
  },

  props: {
    email: String,
  },

  data (): Data {
    return {
      pats: [ ],
    }
  },

  watch: {
  },

  methods: {

    fetch() {
      let query = db.collection('users').doc(this.email).collection('pats') as firebase.firestore.Query

      query.get().then((pats) => {
        this.pats = pats.docs
      })
    },
  },

  created () {
    this.fetch()
  },

})

</script>