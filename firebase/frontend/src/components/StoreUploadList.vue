<template>
  <div>

    <!-- List of uploads in store -->

    <v-data-table
      :headers="headers"
      :items="storeUploads"
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
        v-slot:item.uploadId="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: item.uploadId } }">{{ item.uploadId }}</router-link>
      </template>

      <template
        v-slot:item.description="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: item.uploadId } }">{{ item.description }}</router-link>
      </template>

      <template
        v-slot:item.buildId="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: item.uploadId } }">{{ item.buildId }}</router-link>
      </template>

      <template
        v-slot:item.timestamp="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: item.uploadId } }">{{ timestampToDisplayString(item.timestamp) }}</router-link>
      </template>

      <template
        v-slot:item.status="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: item.uploadId } }">{{ item.status }}</router-link>
      </template>


    </v-data-table>

  </div>
</template>

<script setup lang="ts">

import { ref } from 'vue'
import dayjs from 'dayjs'

import { api } from '../adminApi'
import { GetStoreUploadResponse } from '../generated/api'

const props = defineProps<{
  store: string,
}>()

const headers = [
{
    text: "Upload ID",
    value: "uploadId",
  },
  {
    text: "Description",
    value: "description",
  },
  {
    text: "Build ID",
    value: "buildId",
  },
  {
    text: "Timestamp",
    value: "timestamp",
  },
  {
    text: "Status",
    value: "status",
  }
]

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

const storeUploads = ref([] as GetStoreUploadResponse[])
const total = ref(1)

function timestampToDisplayString(timestamp: string): string {
  return dayjs(timestamp).format('YYYY-MM-DD HH:mm')
}

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
    const storeUploadsResponse = await api.getStoreUploads(props.store, getSortIndex(), (options.page - 1) * options.itemsPerPage, options.itemsPerPage)
    storeUploads.value = storeUploadsResponse.data.uploads
    total.value = storeUploadsResponse.data.pagination.total
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