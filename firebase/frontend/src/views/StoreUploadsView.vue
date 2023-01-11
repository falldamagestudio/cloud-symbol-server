<template>
  <div>
    <h1>Uploads</h1>

    <v-breadcrumbs
      :items="breadcrumbItems"
    />

    <StoreUploadList :store="$route.params.store"/>
  </div>
</template>

<script setup lang="ts">

import StoreUploadList from '../components/StoreUploadList.vue';
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
]

</script>
