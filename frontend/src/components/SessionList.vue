<template>
  <div>

    <!-- Filters -->

    <v-row align="end">

      <!-- Start date picker -->

      <v-col
        cols="12"
        sm="3"
        md="2"
      >

        <v-menu
          v-model="startDatePickerActive"
          :close-on-content-click="false"
          transition="scale-transition"
          offset-y
          min-width="auto"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="startDate"
              label="Start date"
              prepend-icon="mdi-calendar"
              readonly
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>

          <v-date-picker
            v-model="startDate"
            @input="startDatePickerActive = false"
          >
            <v-spacer></v-spacer>

            <v-btn
              text
              color="primary"
              @click="startDate = ''; startDatePickerActive = false"
            >
              Clear
            </v-btn>

          </v-date-picker>
        </v-menu>

      </v-col>

      <!-- End date picker -->

      <v-col
        cols="12"
        sm="3"
        md="2"
      >

        <v-menu
          v-model="endDatePickerActive"
          :close-on-content-click="false"
          transition="scale-transition"
          offset-y
          min-width="auto"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="endDate"
              label="End date"
              prepend-icon="mdi-calendar"
              readonly
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>

          <v-date-picker
            v-model="endDate"
            @input="endDatePickerActive = false"
          >
            <v-spacer></v-spacer>

            <v-btn
              text
              color="primary"
              @click="endDate = ''; endDatePickerActive = false"
            >
              Clear
            </v-btn>

          </v-date-picker>
        </v-menu>
      </v-col>

      <!-- Maps picker -->

      <v-col
        cols="12"
        sm="6"
        md="4"
      >

        <v-select
          v-model="selectedMaps"
          :items="availableMaps"
          attach
          chips
          label="Maps"
          multiple
        ></v-select>

      </v-col>

    </v-row>

    <!-- Results -->

    <v-row>
      <template v-for="session in sessions">
        <v-col v-bind:key="session.id" cols="12">

          <SessionPreview :session="session"/>

        </v-col>
      </template>
    </v-row>
  </div>
</template>

<script lang="ts">

import Vue from 'vue'
import type firebase from 'firebase'
import { db } from '../firebase'
import SessionPreview from './SessionPreview.vue'
import { DateTime } from 'luxon'

interface Data {
  sessions: firebase.firestore.QueryDocumentSnapshot<firebase.firestore.DocumentData>[]
  startDatePickerActive: boolean
  startDate: string
  endDatePickerActive: boolean
  endDate: string
  availableMaps: string[]
  selectedMaps: string[]
}

export default Vue.extend({

  components: {
    SessionPreview
  },

  data (): Data {
    return {
      sessions: [ ],
      startDatePickerActive: false,
      startDate: "",
      endDatePickerActive: false,
      endDate: "",
      availableMaps: [ ],
      selectedMaps: [ ],
    }
  },

  watch: {
    startDate: function (val) {
      this.fetch()
    },

    endDate: function (val) {
      this.fetch()
    },

    selectedMaps: function (val) {
      this.fetch()
    },
  },

  methods: {

    fetch() {
      let query = db.collection('sessions') as firebase.firestore.Query

      if (this.startDate) {
        // Hack: Timestamps are stored as regular strings in the database. Use string comparison when filtering them.
        //query = query.where('Timestamp', '>=', new Date(this.startDate))
        query = query.where('Timestamp', '>=', this.startDate)
      }

      if (this.endDate) {
        // Hack: Timestamps are stored as regular strings in the database. Use string comparison when filtering them.
        // query = query.where('Timestamp', '<', DateTime.fromISO(this.endDate, {zone: 'utc'}).plus({ days: 1 }).toJSDate())
        query = query.where('Timestamp', '<', DateTime.fromISO(this.endDate, {zone: 'utc'}).plus({ days: 1 }).toISO())
      }

      if (this.selectedMaps.length !== 0) {
        query = query.where('Map', 'in', this.selectedMaps)
      }

      query.get().then((sessions) => {
        this.sessions = sessions.docs
      })
    },

    fetchMaps() {
      db.collection('availableMaps').get().then((availableMaps) => {
        this.availableMaps = availableMaps.docs.map(availableMap => availableMap.get('DisplayName'))
      })
    },
  },

  created () {
    this.fetchMaps()
    this.fetch()
  },

})

</script>