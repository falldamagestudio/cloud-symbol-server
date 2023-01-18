<template>
  <div>

    <td>
      <tr>
        Description: {{ storeUpload?.description }}
      </tr>
      <tr>
        Build ID: {{ storeUpload?.buildId }}
      </tr>
      <tr>
        Timestamp: {{ storeUpload?.timestamp }}
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

import { computed, ref, watch } from 'vue'

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

const storeUpload = ref(null as GetStoreUploadResponse)

async function fetch() {

  try {
    const storeUploadResponse = await api.getStoreUpload(props.upload, props.store)
    storeUpload.value = storeUploadResponse.data
  } catch (error) {
    console.log(error)
  }
}

fetch()

</script>