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

import { collection, DocumentData, getDocs, query, QueryDocumentSnapshot } from 'firebase/firestore'

import { db } from '../firebase'
import PATListEntry from './PATListEntry.vue'

import { api } from '../adminApi'

interface Data {
  pats: QueryDocumentSnapshot<DocumentData>[]
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

    async fetch() {

      try {
        const response = await api.getTokens()
        console.log(response)
      } catch (error) {
        console.log(error)
      }

      //response.then(value => { this.pats = value })

      // const patsRef = collection(db, 'users', this.email, 'pats')
      // const patsQuery = query(patsRef)
      // const patsSnapshot = await getDocs(patsQuery)
      // this.pats = patsSnapshot.docs
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