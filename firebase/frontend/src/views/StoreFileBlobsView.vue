<template>
  <div>
    <h1>Blobs</h1>

    <v-breadcrumbs
      :items="breadcrumbItems"
    />

    <StoreFileBlobList :store="$route.params.store" :file="$route.params.file"/>
  </div>
</template>

<script setup lang="ts">

import StoreFileBlobList from '../components/StoreFileBlobList.vue';
import router from '../router/index'

const breadcrumbItems = [
  {
    text: 'Stores',
    disabled: false,
    href: router.match({
      name: 'stores',
    }).path,
  },
  {
    text: router.currentRoute.params.store,
    disabled: false,
    href: router.match({
      name: 'storeFiles',
      params: {
        store: router.currentRoute.params.store,
      },
    }).path,
  },
  {
    text: router.currentRoute.params.file,
    disabled: false,
    href: router.match({
      name: 'storeFileBlobs',
      params: {
        store: router.currentRoute.params.store,
        file: router.currentRoute.params.file,
      },
    }).path,
  },
]


</script>
