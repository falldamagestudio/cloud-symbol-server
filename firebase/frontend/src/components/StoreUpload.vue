<template>
  <div>

    <td>
      <tr>
        Upload ID: {{ storeUpload?.uploadId }}
      </tr>
      <tr>
        Description: {{ storeUpload?.description }}
      </tr>
      <tr>
        Build ID: {{ storeUpload?.buildId }}
      </tr>
      <tr>
        Uploaded at: {{ timestampToDisplayString(storeUpload?.uploadTimestamp) }}
      </tr>
      <tr>
        Expired at: {{ storeUpload?.expiryTimestamp && timestampToDisplayString(storeUpload?.expiryTimestamp) || "" }}
      </tr>
      <tr>
        Status: {{ storeUpload?.status }}
      </tr>
    </td>

    <!-- List of files in upload -->

    <v-data-table
      :headers="headers"
      :items="storeUpload?.files"
    >
      <template
        v-slot:item.fileName="{ item }"
      >
        {{ item.fileName }}
      </template>    

      <template
        v-slot:item.blobIdentifier="{ item }"
      >
        {{ item.blobIdentifier }}
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

import { ref } from 'vue'
import dayjs from 'dayjs'

import { api } from '../adminApi'
import { GetStoreUploadResponse } from '../generated/api'

const props = defineProps<{
  store: string,
  upload: string,
}>()

const headers = [
  {
    text: "File Name",
    value: "fileName",
  },
  {
    text: "Blob Identifier",
    value: "blobIdentifier",
  },
  {
    text: "Status",
    value: "status",
  }
]

const storeUpload = ref(null as (null | GetStoreUploadResponse))

function timestampToDisplayString(timestamp: string): string {
  return dayjs(timestamp).format('YYYY-MM-DD HH:mm')
}

async function fetch() {

  try {
    const storeUploadResponse = await api.getStoreUpload(props.upload as unknown as number, props.store)
    storeUpload.value = storeUploadResponse.data
  } catch (error) {
    console.log(error)
  }
}

fetch()

</script>