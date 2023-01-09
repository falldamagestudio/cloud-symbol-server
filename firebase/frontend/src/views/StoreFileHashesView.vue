<template>
  <div>
    <h1>Hashes</h1>
    <StoreFileHashList :store="$route.params.store" :file="$route.params.file"/>
  </div>
</template>

<script setup lang="ts">

import StoreFileHashList from '../components/StoreFileHashList.vue';

</script>
