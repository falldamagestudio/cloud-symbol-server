<template>
  <div>
    <h1>Personal Access Tokens</h1>
    <PATList :email="email"/>
  </div>
</template>

<script setup lang="ts">

import store from '../store/index'
import PATList from '../components/PATList.vue';

const email = store.state.user!.email!

</script>
