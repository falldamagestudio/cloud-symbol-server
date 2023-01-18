<template>
  <div>

    <!-- List of file-blobs for file -->

    <v-data-table
      :headers="headers"
      :items="storeFileBlobs"
      :server-items-length="total"
      :options="options"
      :page="options.page"
      :items-per-page="options.itemsPerPage"
      @pagination="updatePagination"
    >
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

      <template
        v-slot:item.operations="{ item }"
      >
        <v-btn
          v-on:click="download(item.blobIdentifier)"
        >
          Download
        </v-btn>
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
    text: "Identifier",
    value: "blobIdentifier",
  },
  {
    text: "Uploaded at",
    value: "uploadTimestamp",
  },
  {
    text: "Status",
    value: "status",
  },
  {
    text: "Operations",
    value: "operations",
  },
]

interface StoreFileBlobEntry {
  blobIdentifier: string
  uploadTimestamp: string
  status: string
}

let options = {
  page: 1,
  itemsPerPage: 5,
}

const storeFileBlobs = ref([] as StoreFileBlobEntry[])
const total = ref(1)

async function fetch() {

  try {
    const storeFileBlobsResponse = await api.getStoreFileBlobs(props.store, props.file, (options.page - 1) * options.itemsPerPage, options.itemsPerPage)
    storeFileBlobs.value.length = 0
    if (storeFileBlobsResponse.data.blobs) {
      for (const blob of storeFileBlobsResponse.data.blobs) {
        storeFileBlobs.value.push({
          blobIdentifier: blob.blobIdentifier,
          uploadTimestamp: blob.uploadTimestamp,
          status: blob.status,
        })
      }
    }
    total.value = storeFileBlobsResponse.data.pagination.total
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

function downloadFileInBrowser(url: string, filename: string) {

  // This is based on https://blog.logrocket.com/programmatic-file-downloads-in-the-browser-9a5186298d5c/

  // Create a new anchor element
  const a = document.createElement('a')

  // Set the href and download attributes for the anchor element
  // You can optionally set other attributes like `title`, etc
  // Especially, if the anchor element will be attached to the DOM
  a.href = url
  a.download = filename

  // Click handler that removes the anchor element after the element has been clicked
  const clickHandler = () => {
    setTimeout(() => {
      a.remove()
    }, 150)
  };

  // Add the click event listener on the anchor element
  a.addEventListener('click', clickHandler, false)
  
  // Programmatically trigger a click on the anchor element
  a.click()
}

async function download(blobIdentifier: string) {

  try {
    const getStoreFileBlobDownloadUrlResponse = await api.getStoreFileBlobDownloadUrl(props.store, props.file, blobIdentifier)

    downloadFileInBrowser(getStoreFileBlobDownloadUrlResponse.data.url, props.file)
  } catch (error) {
    console.log(error)
  }
}

fetch()

</script>