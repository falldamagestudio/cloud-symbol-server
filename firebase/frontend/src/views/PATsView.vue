<template>
  <div>
    <h1>Personal Access Tokens</h1>
    <PATList :email="email"/>
  </div>
</template>

<script setup lang="ts">

import { useAuthUserStore } from '../stores/authUser'
import PATList from '../components/PATList.vue';

const authUserStore = useAuthUserStore()

const email = authUserStore.user!.email!

</script>
