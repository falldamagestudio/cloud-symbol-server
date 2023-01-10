<template>
  <div>
    <h1>Hashes</h1>

    <v-breadcrumbs
      :items="breadcrumbItems"
    />

    <StoreFileHashList :store="$route.params.store" :file="$route.params.file"/>
  </div>
</template>

<script setup lang="ts">

import StoreFileHashList from '../components/StoreFileHashList.vue';
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
      name: 'storeFileHashes',
      params: {
        store: router.currentRoute.params.store,
        file: router.currentRoute.params.file,
      },
    }).path,
  },
]


</script>
