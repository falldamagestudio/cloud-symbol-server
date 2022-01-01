<template>
  <div>

    <!-- Header -->

    <v-row>

      <v-col
      >
        Personal Access Tokens
      </v-col>

      <v-col
        class="text-right"
      >

        <v-btn
        >
          Generate new token
        </v-btn>

        <v-btn
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