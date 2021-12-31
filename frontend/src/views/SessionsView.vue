<template>
  <div>
    <h1>Sessions</h1>
    <SessionList/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import SessionList from '../components/SessionList.vue';

export default Vue.extend({

  components: {
    SessionList,
  },

  data () {
    return {} 
  },

})
</script>
