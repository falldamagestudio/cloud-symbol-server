<template>
  <div>
    <h1>Files</h1>

    <v-breadcrumbs
      :items="breadcrumbItems"
    />

    <StoreFileList :store="$route.params.store"/>
  </div>
</template>

<script setup lang="ts">

import StoreFileList from '../components/StoreFileList.vue';
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
]

</script>
