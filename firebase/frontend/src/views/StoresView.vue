<template>
  <div>
    <h1>Stores</h1>

    <v-breadcrumbs
      :items="breadcrumbItems"
    />

    <StoreList/>
  </div>
</template>

<script setup lang="ts">

import StoreList from '../components/StoreList.vue';
import router from '../router/index'

const breadcrumbItems = [
  {
    text: 'Stores',
    disabled: false,
    href: router.match({
      name: 'stores',
    }).path,
  },
]

</script>
