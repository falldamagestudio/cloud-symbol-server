<template>
  <div>

    <!-- List of file-hashes for file -->

    <v-data-table
      :headers="headers"
      :items="storeFileHashes"
      :server-items-length="total"
      :options="options"
      :page="options.page"
      :items-per-page="options.itemsPerPage"
      @pagination="updatePagination"
    >
      <template
        v-slot:item.name="{ item }"
      >
        {{ item.hash }}
      </template>    
      <template
        v-slot:item.status="{ item }"
      >
        {{ item.status }}
      </template>    
    </v-data-table>

  </div>
</template>

<script setup lang="ts">

import { computed, ref, watch } from 'vue'

import { api } from '../adminApi'

const props = defineProps<{
  store: string,
  file: string,
}>()

const headers = [
  {
    text: "Name",
    value: "name",
  },
  {
    text: "Status",
    value: "status",
  }
]

interface StoreFileHashEntry {
  hash: string
  status: string
}

let options = {
  page: 1,
  itemsPerPage: 5,
}

const storeFileHashes = ref([] as StoreFileHashEntry[])
const total = ref(1)

async function fetch() {

  try {
    const storeFileHashesResponse = await api.getStoreFileHashes(props.store, props.file, (options.page - 1) * options.itemsPerPage, options.itemsPerPage)
    storeFileHashes.value.length = 0
    if (storeFileHashesResponse.data.hashes) {
      for (const hash of storeFileHashesResponse.data.hashes) {
        storeFileHashes.value.push({
          hash: hash.hash,
          status: hash.status,
        })
      }
    }
    total.value = storeFileHashesResponse.data.pagination.total
  } catch (error) {
    console.log(error)
  }
}

function updatePagination(newOptions: {
  page: number,
  itemsPerPage: number,
  pageStart: number,
  pageStop: number,
  pageCount: number,
  itemsLength: number
}) {
  if ((options.page != newOptions.page) || (options.itemsPerPage != newOptions.itemsPerPage)) {
    options = newOptions
    fetch()
  }
}

fetch()

</script>