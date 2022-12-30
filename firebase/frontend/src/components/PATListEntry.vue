<template>

  <tr>
    <!-- Personal Access Token ID -->

    <td>
      {{ abbreviateId(pat.id) }}

      <!-- Shortcut for copying full PAT ID to clipboard -->
      <v-btn
        icon
        @click="copyTextToClipboard(pat.id)"
      >
        <v-icon
          small
        >
          mdi-content-copy
        </v-icon>
      </v-btn>
    </td>

    <!-- Personal Access Token creation timestamp -->

    <td>
      {{ timestampToDisplayString(pat.get('creationTimestamp')) }}
    </td>

    <!-- Personal Access Token description -->

    <td>
      {{ pat.get('description') }}
    </td>

    <td>
      <div class="text-right">

        <!-- "Use" modal dialog box -->

        <v-dialog
          v-model="useDialogVisible"
          width="1000"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              v-bind="attrs"
              v-on="on"
            >
              Use
            </v-btn>
          </template>

          <PATUsageGuide
            :email="email"
            :pat="pat"
            @done="patUsageGuideDone"
          />

        </v-dialog>

        <!-- "Revoke" button -->

        <v-btn
          color="error--text"
          v-on:click="revoke()"
        >
          Revoke
        </v-btn>
      </div>

    </td>

  </tr>

</template>

<script setup lang="ts">

import { ref } from 'vue'
import { doc, deleteDoc, Timestamp } from 'firebase/firestore'
import dayjs from 'dayjs'

import { db } from '../firebase'
import PATUsageGuide from './PATUsageGuide.vue'

const props = defineProps<{
  email: string,
  pat: any,
}>()

const emit = defineEmits<{
  (e: 'refresh'): void
}>()

const useDialogVisible = ref(false)

function copyTextToClipboard(text: string) {
  navigator.clipboard.writeText(text)
}

async function revoke() {
  const patDocRef = doc(db, 'users', props.email, 'pats', props.pat.id)
  await deleteDoc(patDocRef)
  emit('refresh')
}

function patUsageGuideDone() {
  useDialogVisible.value = false
}

function abbreviateId(id: string): string {
  return `${id.slice(0, 4)}...${id.slice(-4)}`
}

function timestampToDisplayString(timestamp: Timestamp): string {
  return dayjs(timestamp.toDate()).format('YYYY-MM-DD HH:mm')
}

</script>