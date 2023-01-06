<template>
  <div>
    <h1>Stores</h1>
    <StoreFileList :store="$route.params.store"/>
  </div>
</template>

<script setup lang="ts">

import StoreFileList from '../components/StoreFileList.vue';

</script>
