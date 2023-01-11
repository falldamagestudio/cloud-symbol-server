<template>
  <div>
    <h1>Uploads</h1>

    <v-breadcrumbs
      :items="breadcrumbItems"
    />

    <StoreUpload :store="$route.params.store" :upload="$route.params.upload"/>
  </div>
</template>

<script setup lang="ts">

import StoreUpload from '../components/StoreUpload.vue';
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
      name: 'storeUploads',
      params: {
        store: router.currentRoute.params.store,
      },
    }).path,
  },
  {
    text: router.currentRoute.params.upload,
    disabled: false,
    href: router.match({
      name: 'storeUpload',
      params: {
        store: router.currentRoute.params.store,
        upload: router.currentRoute.params.upload,
      },
    }).path,
  },
]

</script>
