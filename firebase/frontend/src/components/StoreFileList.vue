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

    <v-data-table
      :headers="headers"
      :items="storeFiles"
      :server-items-length="total"
      :options="options"
      :page="options.page"
      :items-per-page="options.itemsPerPage"
      @pagination="updatePagination"
    >
      <template
        v-slot:item.name="{ item }"
      >
        <a :href="generateStoreFileHref(item)">{{ item.name }}</a>
      </template>    
    </v-data-table>

  </div>
</template>

<script setup lang="ts">

import { computed, ref, watch } from 'vue'

import { api } from '../adminApi'

const props = defineProps<{
  store: string,
}>()

const headers = [
  {
    text: "Name",
    value: "name",
  }
]

interface StoreFileEntry {
  name: string
}

let options = {
  page: 1,
  itemsPerPage: 5,
}

const storeFiles = ref([] as StoreFileEntry[])
const total = ref(1)

function generateStoreFileHref(file: StoreFileEntry): string {
  const storeId = "blah"
  return `http://localhost:8080/stores/${storeId}/files/${file.name}`
}

async function fetch() {

  try {
    const storeFilesResponse = await api.getStoreFiles(props.store, (options.page - 1) * options.itemsPerPage, options.itemsPerPage)
    storeFiles.value.length = 0
    if (storeFilesResponse.data.files) {
      for (const storeFileId of storeFilesResponse.data.files) {
        storeFiles.value.push({
          name: storeFileId
        })
      }
    }
    total.value = storeFilesResponse.data.pagination.total
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