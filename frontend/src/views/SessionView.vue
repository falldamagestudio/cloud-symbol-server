<template>
  <div>
    <h1>Session details</h1>
    <SessionDetails :sessionId="$route.params.id"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import SessionDetails from '../components/SessionDetails.vue';

export default Vue.extend({

  components: {
    SessionDetails,
  },

  data () {
    return {} 
  },

})
</script>
