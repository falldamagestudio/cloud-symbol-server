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
              v-bind:key="pat.token"
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

<script setup lang="ts">

import { ref } from 'vue'

import PATListEntry from './PATListEntry.vue'
import { api } from '../adminApi'
import { GetTokenResponse } from '../generated/api'

const pats = ref([] as GetTokenResponse[])

async function fetch() {

  try {
    const response = await api.getTokens()
    pats.value = response.data
  } catch (error) {
    console.log(error)
  }
}

function refresh() {
  fetch()
}

fetch()

</script>