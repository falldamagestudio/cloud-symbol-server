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

    <v-pagination
      :length="totalPages"
      v-model="currentPage"
    >
    </v-pagination>
  </div>
</template>

<script setup lang="ts">

import { computed, ref, watch } from 'vue'

import { api } from '../adminApi'

const props = defineProps<{
  store: string,
}>()

const storeFiles = ref([] as string[])

const currentPage = ref(1)
const limit = ref(1)
const totalPages = ref(1)

watch(currentPage, (newPage) => {
  refresh()
})

watch(limit, (newLimit) => {
  currentPage.value = 1
})

const total = ref(undefined as number)
watch([total, limit], ([newTotal, newLimit]) => {
  totalPages.value = Math.ceil(newTotal / newLimit)
})

async function fetch() {

  try {
    storeFiles.value.length = 0
    const storeFilesResponse = await api.getStoreFiles(props.store, (currentPage.value - 1) * limit.value, limit.value)
    for (const storeFileId of storeFilesResponse.data.files) {
      storeFiles.value.push(storeFileId)
    }
    total.value = storeFilesResponse.data.pagination.total
  } catch (error) {
    console.log(error)
  }
}

function refresh() {
  fetch()
}

fetch()

</script>