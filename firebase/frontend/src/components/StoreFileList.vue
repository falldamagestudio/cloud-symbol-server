<template>
  <div>

    <!-- List of files in store -->

    <v-data-table
      :headers="headers"
      :items="storeFiles"
      :server-items-length="total"
      :options="options"
      :footer-props="{
        showCurrentPage: true,
        showFirstLastPage: true,
        itemsPerPageOptions: [
          10,
          25,
          100,
        ],
      }"
      @update:options="updateOptions"
    >
      <template
        v-slot:item.fileName="{ item }"
      >
        <router-link :to="{ name: 'storeFileBlobs', params: { store: store, file: item.fileName } }">{{ item.fileName }}</router-link>
      </template>    
    </v-data-table>

  </div>
</template>

<script setup lang="ts">

import { ref } from 'vue'

import { api } from '../adminApi'

const props = defineProps<{
  store: string,
}>()

const headers = [
  {
    text: "Name",
    value: "fileName",
  }
]

interface StoreFileEntry {
  fileName: string
}

interface DataTableOptions {
  page: number,
  itemsPerPage: number,
  sortBy: string[],
  sortDesc: boolean[],
  groupBy: string[],
  groupDesc: boolean[],
  multiSort: boolean,
  mustSort: boolean
}

let options: DataTableOptions = {
  page: 1,
  itemsPerPage: 25,
  sortBy: [],
  sortDesc: [],
  groupBy: [],
  groupDesc: [],
  multiSort: false,
  mustSort: false,
}

const storeFiles = ref([] as StoreFileEntry[])
const total = ref(1)

function getSortIndex(): string | undefined {
  if (options.sortBy.length == 0) {
    return undefined
  } else {
    const sortDirection = (options.sortDesc[0] ? "-" : "")
    const sortKey = options.sortBy[0]
    return `${sortDirection}${sortKey}`
  }
}

async function fetch() {

  try {
    const storeFilesResponse = await api.getStoreFiles(props.store, getSortIndex(), (options.page - 1) * options.itemsPerPage, options.itemsPerPage)
    storeFiles.value.length = 0
    if (storeFilesResponse.data.files) {
      for (const storeFileId of storeFilesResponse.data.files) {
        storeFiles.value.push({
          fileName: storeFileId
        })
      }
    }
    total.value = storeFilesResponse.data.pagination.total
  } catch (error) {
    console.log(error)
  }
}

function updateOptions(newOptions: DataTableOptions) {
  if ((options.page != newOptions.page)
    || (options.itemsPerPage != newOptions.itemsPerPage)
    || (options.sortBy != newOptions.sortBy)
    || (options.sortDesc != newOptions.sortDesc) ) {
    options = newOptions
    fetch()
  }
}

fetch()

</script>