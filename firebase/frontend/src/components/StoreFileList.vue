<template>
  <div>

    <!-- Header -->

    <v-row>

      <v-col
      >
        Store Files
      </v-col>

    </v-row>

    <!-- List of files in store -->

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
            v-for="file in storeFiles"
          >
            <tr
              v-bind:key="file"
            >
              <td>
                {{file}}
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

const props = defineProps<{
  store: string,
}>()

const storeFiles = ref([] as string[])

async function fetch() {

  try {
    storeFiles.value.length = 0
    const storeFileIdsResponse = await api.getStoreFileIds(props.store)
    for (const storeFileId of storeFileIdsResponse.data) {
      storeFiles.value.push(storeFileId)
    }
  } catch (error) {
    console.log(error)
  }
}

function refresh() {
  fetch()
}

fetch()

</script>