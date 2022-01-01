<template>
  <v-hover v-slot:default="{ hover }">
    <v-card
      :elevation="hover ? 6 : 2"
    >

      <v-card-text>
        <v-row>

          <v-col>
            {{ pat.id }}
          </v-col>

          <v-spacer/>

          <v-col
            class="text-right"
          >
            <v-btn>
              Revoke
            </v-btn>
          </v-col>

        </v-row>
      </v-card-text>
      
    </v-card>
  </v-hover>
</template>

<script lang="ts">

import Vue from 'vue'

export default Vue.extend({

  props: {
    pat: Object,
  },

})

</script>