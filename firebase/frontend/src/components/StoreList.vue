<template>
  <div>

    <!-- List of stores -->

    <v-simple-table>

      <template
        v-slot:default
      >
        <thead>
          <tr>
            <th class="text-left">
              Store name
            </th>
          </tr>
        </thead>

        <tbody>

          <template
            v-for="store in stores"
          >
            <tr
              v-bind:key="store"
            >
              <td>
                {{store}} -
                <router-link :to="{ name: 'storeFiles', params: { store: store } }">Files</router-link>, 
                <router-link :to="{ name: 'storeUploads', params: { store: store } }">Uploads</router-link>
              </td>
            </tr>
          </template>

        </tbody>
      </template>

    </v-simple-table>
  </div>
</template>

<script setup lang="ts">

import { ref } from 'vue'

import { api } from '../adminApi'

const stores = ref([] as string[])

async function fetch() {

  try {
    const response = await api.getStores()
    stores.value = response.data
  } catch (error) {
    console.log(error)
  }
}

fetch()

</script>