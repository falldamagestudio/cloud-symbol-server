<template>
  <div>

    <!-- List of uploads in store -->

    <v-data-table
      :headers="headers"
      :items="storeUploads"
      :server-items-length="total"
      :options="options"
      :page="options.page"
      :items-per-page="options.itemsPerPage"
      @pagination="updatePagination"
    >
      <template
        v-slot:item.uploadId="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: storeUploads.indexOf(item).toString() } }">{{ item.uploadId }}</router-link>
      </template>

      <template
        v-slot:item.description="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: storeUploads.indexOf(item).toString() } }">{{ item.description }}</router-link>
      </template>

      <template
        v-slot:item.buildId="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: storeUploads.indexOf(item).toString() } }">{{ item.buildId }}</router-link>
      </template>

      <template
        v-slot:item.timestamp="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: storeUploads.indexOf(item).toString() } }">{{ timestampToDisplayString(item.timestamp) }}</router-link>
      </template>

      <template
        v-slot:item.status="{ item }"
      >
        <router-link :to="{ name: 'storeUpload', params: { store: store, upload: storeUploads.indexOf(item).toString() } }">{{ item.status }}</router-link>
      </template>


    </v-data-table>

  </div>
</template>

<script setup lang="ts">

import { computed, ref, watch } from 'vue'
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

let options = {
  page: 1,
  itemsPerPage: 5,
}

const storeUploads = ref([] as GetStoreUploadResponse[])
const total = ref(1)

function timestampToDisplayString(timestamp: string): string {
  return dayjs(timestamp).format('YYYY-MM-DD HH:mm')
}

async function fetch() {

  try {
    const storeUploadsResponse = await api.getStoreUploads(props.store, (options.page - 1) * options.itemsPerPage, options.itemsPerPage)
    storeUploads.value = storeUploadsResponse.data.uploads
    total.value = storeUploadsResponse.data.pagination.total
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