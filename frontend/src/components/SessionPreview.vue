<template>
  <v-hover v-slot:default="{ hover }">
    <v-card
      :to="{ name: 'session', params: { id: session.id } }"
      :elevation="hover ? 6 : 2"
    >

      <v-card-text>
        <v-row>
          <v-col cols="2">
            {{ date }}
          </v-col>

          <v-col cols="2">
            cs:{{ session.get('ChangeSet') }}
          </v-col>
          <v-col cols="2">
            {{ session.get('Map') }}
          </v-col>

          <v-col cols="4">
            Perf:
            {{ formatMS(session.get('PerfMinMS')) }}
            / 
            {{ formatMS(session.get('PerfAvgMS')) }}
            /
            {{ formatMS(session.get('PerfMaxMS')) }}
            ms
          </v-col>

        </v-row>
      </v-card-text>
      
    </v-card>
  </v-hover>
</template>

<script lang="ts">

import Vue from 'vue'
import { DateTime } from 'luxon'

interface Data {
  date: string
}

export default Vue.extend({

  props: {
    session: Object,
  },

  methods: {
    formatMS(ms: number) {
      return ms.toFixed(2)
    },
  },

  data (): Data {
    return {
        // Hack: Timestamps are stored as regular strings in the database.
        // date: DateTime.fromMillis(this.session.get('Timestamp').seconds * 1000).toFormat('yyyy-MM-dd')
        date: DateTime.fromISO(this.session.get('Timestamp'), {zone: 'utc'}).toFormat('yyyy-MM-dd'),
    }
  },

})

</script>