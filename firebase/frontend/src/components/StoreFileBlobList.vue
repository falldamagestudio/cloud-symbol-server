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
      :footer-props="{
        showCurrentPage: true,
        showFirstLastPage: true,
        itemsPerPageOptions: [
          10,
          25,
          100,
        ],
      }"
      @pagination="updatePagination"
    >
      <template
        v-slot:item.blobIdentifier="{ item }"
      >
        {{ item.blobIdentifier }}

        <!-- Shortcut for copying identifier to clipboard -->
        <v-btn
          icon
          @click="copyTextToClipboard(item.blobIdentifier)"
        >
          <v-icon
            small
          >
            mdi-content-copy
          </v-icon>
        </v-btn>
      </template>    

      <template
        v-slot:item.uploadTimestamp="{ item }"
      >
        {{ timestampToDisplayString(item.uploadTimestamp) }}
      </template>    

      <template
        v-slot:item.type="{ item }"
      >
        {{ item.type }}
      </template>    

      <template
        v-slot:item.size="{ item }"
      >
        {{ item.size }}
      </template>    

      <template
        v-slot:item.contentHash="{ item }"
      >
        {{ abbreviateHash(item.contentHash) }}

        <!-- Shortcut for copying full hash to clipboard -->
        <v-btn
          icon
          @click="copyTextToClipboard(item.contentHash)"
        >
          <v-icon
            small
          >
            mdi-content-copy
          </v-icon>
        </v-btn>

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
import dayjs from 'dayjs'

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
    text: "Type",
    value: "type",
  },
  {
    text: "Size",
    value: "size",
  },
  {
    text: "Content SHA256 Hash",
    value: "contentHash",
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
  type: string,
  size: string,
  contentHash: string,
  status: string
}

let options = {
  page: 1,
  itemsPerPage: 25,
}

const storeFileBlobs = ref([] as StoreFileBlobEntry[])
const total = ref(1)

function translateBlobType(blobType: string | undefined): string {
  if (blobType == 'pdb') {
    return 'PDB'
  } else if (blobType == 'pe') {
    return 'PE'
  } else {
    return "Undefined"
  }
}

function abbreviateHash(hash: string): string {
  return `${hash.slice(0, 4)}...${hash.slice(-4)}`
}

function timestampToDisplayString(timestamp: string): string {
  return dayjs(timestamp).format('YYYY-MM-DD HH:mm')
}

function copyTextToClipboard(text: string) {
  navigator.clipboard.writeText(text)
}

async function fetch() {

  try {
    const storeFileBlobsResponse = await api.getStoreFileBlobs(props.store, props.file, (options.page - 1) * options.itemsPerPage, options.itemsPerPage)
    storeFileBlobs.value.length = 0
    if (storeFileBlobsResponse.data.blobs) {
      for (const blob of storeFileBlobsResponse.data.blobs) {
        const size = blob.size ?? 0
        storeFileBlobs.value.push({
          blobIdentifier: blob.blobIdentifier,
          uploadTimestamp: blob.uploadTimestamp,
          type: translateBlobType(blob.type),
          size: (size ? size.toString() : "Unknown"),
          contentHash: blob.contentHash ?? "Unknown",
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