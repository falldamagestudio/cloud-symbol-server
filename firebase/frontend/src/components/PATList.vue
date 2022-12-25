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
          :to="{ name: 'new-pat' }"
        >
          <v-icon>mdi-plus</v-icon>
          Generate new token
        </v-btn>

      </v-col>

    </v-row>

    <!-- Existing tokens -->

    <v-simple-table>

      <template
        v-slot:default
      >
        <thead>
          <tr>
            <th class="text-left">
              ID
            </th>
            <th class="text-left">
              Created
            </th>
            <th class="text-left">
              Description
            </th>
            <th class="text-right">
              Actions
            </th>
          </tr>
        </thead>

        <tbody>

          <template
            v-for="pat in pats"
          >
            <PATListEntry
              v-bind:key="pat.id"
              :email="email"
              :pat="pat"
              @refresh="refresh()"
            />
          </template>

        </tbody>
      </template>

    </v-simple-table>
  </div>
</template>

<script lang="ts">

import Vue from 'vue'

// v9 compat packages are API compatible with v8 code
import type firebase from 'firebase/compat/app'

import { db } from '../firebase'
import PATListEntry from './PATListEntry.vue'

interface Data {
  pats: firebase.firestore.QueryDocumentSnapshot<firebase.firestore.DocumentData>[]
}

export default Vue.extend({

  components: {
    PATListEntry
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
      const query = db.collection('users').doc(this.email).collection('pats') as firebase.firestore.Query

      query.get().then((pats) => {
        this.pats = pats.docs
      })
    },

    refresh() {
      this.fetch()
    },
  },

  created () {
    this.fetch()
  },

})

</script>