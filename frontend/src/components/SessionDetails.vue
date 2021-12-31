<template>
  <v-card :loading="loading">

    <template v-if="!this.loading">

      <v-card-subtitle>Date: {{ this.date }}</v-card-subtitle>
      <v-card-subtitle>Changeset: {{ this.session.get('ChangeSet') }}</v-card-subtitle>
      <v-card-subtitle>Map: {{ this.session.get('Map') }}</v-card-subtitle>
      <v-card-subtitle>Test Machine: {{ this.session.get('TestMachineName') }}</v-card-subtitle>
      <v-card-subtitle>Official run: {{ this.session.get('Official') ? "Yes" : "No" }}</v-card-subtitle>

      <v-divider/>

      <v-data-table
        :headers="headers"
        :items="sessionProbesData"
        :items-per-page="100"
        class="elevation-1"
      >
        <template v-slot:item.DrawThreadMS="{ item }">
          {{ formatMS(item.DrawThreadMS) }}
        </template>

        <template v-slot:item.GameThreadMS="{ item }">
          {{ formatMS(item.GameThreadMS) }}
        </template>

        <template v-slot:item.GpuMS="{ item }">
          {{ formatMS(item.GpuMS) }}
        </template>

        <template v-slot:item.RawFrameMS="{ item }">
          {{ formatMS(item.RawFrameMS) }}
        </template>

      </v-data-table>

    </template>
    <template v-else>

      <v-card-subtitle/>
    </template>
  </v-card>
</template>

<script lang="ts">

import Vue from 'vue'
import type firebase from 'firebase'
import { DataTableHeader } from 'vuetify'
import { db } from '../firebase'
import { DateTime } from 'luxon'

interface ProbeData {
  Name: string
  DrawThreadMS: string
  GameThreadMS: string
  GpuMS: string
  RawFrameMS: string
}

interface Data {
  sessionProbesData: ProbeData[]
  hasSession: boolean
  session: firebase.firestore.DocumentSnapshot<firebase.firestore.DocumentData> | null
  headers: DataTableHeader[]
  date: string
  loading: boolean
}

export default Vue.extend({

  props: {
    sessionId: String,
  },

  data (): Data {
    return {
      sessionProbesData: [],
      hasSession: false,
      session: null,
      headers: [
        {
          text: "Probe",
          value: "Name",
        },
        {
          text: "DrawThreadMS",
          value: 'DrawThreadMS',
        },
        {
          text: "GameThreadMS",
          value: 'GameThreadMS',
        },
        {
          text: "GpuMS",
          value: 'GpuMS',
        },
        {
          text: "RawFrameMS",
          value: 'RawFrameMS',
        },
      ],
      date: "",
      loading: true,
    }
  },

  methods: {
    signalLoadingPartiallyComplete() {
      if (this.session && this.sessionProbesData)
        this.loading = false
    },

    formatMS(ms: number) {
      return ms.toFixed(2)
    },
  },

  created () {
    db.collection('sessions').doc(this.sessionId).get().then((session) => {
      this.session = session

      // Hack: Timestamps are stored as regular strings in the database.
      // date: DateTime.fromMillis(this.session.get('Timestamp').seconds * 1000).toFormat('yyyy-MM-dd')
      this.date = DateTime.fromISO(this.session.get('Timestamp'), {zone: 'utc'}).toFormat('yyyy-MM-dd')

      this.signalLoadingPartiallyComplete()
    })

    db.collection('sessions').doc(this.sessionId).collection('probes').get().then((probes) => {
      this.sessionProbesData = probes.docs.map((probe) => ({
          DrawThreadMS: probe.get('DrawThreadMS'),
          GameThreadMS: probe.get('GameThreadMS'),
          GpuMS: probe.get('GpuMS'),
          Name: probe.get('Name'),
          RawFrameMS: probe.get('RawFrameMS'),
        }) as ProbeData)
      this.signalLoadingPartiallyComplete()
    })
  },

})

</script>