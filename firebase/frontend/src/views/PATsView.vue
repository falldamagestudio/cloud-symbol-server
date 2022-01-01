<template>
  <div>
    <h1>Personal Access Tokens</h1>
    <PATList :email="email"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import PATList from '../components/PATList.vue';
import store from '../store/index'

interface Data {
  email: string
}

export default Vue.extend({

  components: {
    PATList,
  },

  data () : Data {
    return {
      email: store.state.user!.email!
    } 
  },

})
</script>
