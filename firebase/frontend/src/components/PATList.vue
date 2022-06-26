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
          v-on:click="generate()"
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
import type firebase from 'firebase'
import { db } from '../firebase'
import PATListEntry from './PATListEntry.vue'

interface Data {
  pats: firebase.firestore.QueryDocumentSnapshot<firebase.firestore.DocumentData>[]
}

// dec2hex :: Integer -> String
// i.e. 0-255 -> '00'-'ff'
function dec2hex (dec: number) : string {
  return dec.toString(16).padStart(2, "0")
}

// generateId :: Integer -> String
function generateId (len: number) {
  var arr = new Uint8Array((len || 40) / 2)
  window.crypto.getRandomValues(arr)
  return Array.from(arr, dec2hex).join('')
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
      let query = db.collection('users').doc(this.email).collection('pats') as firebase.firestore.Query

      query.get().then((pats) => {
        this.pats = pats.docs
      })
    },

    generate() {

      const id = generateId(32)

      db.collection('users').doc(this.email).collection('pats').doc(id).set({}).then(() => {
        this.fetch()
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