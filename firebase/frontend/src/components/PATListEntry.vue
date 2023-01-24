<template>

  <tr>
    <!-- Personal Access Token token-->

    <td>
      {{ abbreviateToken(pat.token) }}

      <!-- Shortcut for copying full PAT token to clipboard -->
      <v-btn
        icon
        @click="copyTextToClipboard(pat.token)"
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
      {{ timestampToDisplayString(pat.creationTimestamp) }}
    </td>

    <!-- Personal Access Token description -->

    <td>
      {{ pat.description }}
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
import dayjs from 'dayjs'

import PATUsageGuide from './PATUsageGuide.vue'
import { api } from '../adminApi'
import { GetTokenResponse } from '../generated/api'


const props = defineProps<{
  email: string,
  pat: GetTokenResponse,
}>()

const emit = defineEmits<{
  (e: 'refresh'): void
}>()

const useDialogVisible = ref(false)

function copyTextToClipboard(text: string) {
  navigator.clipboard.writeText(text)
}

async function revoke() {

  try {
    const response = await api.deleteToken(props.pat.token)
  } catch (error) {
    console.log(error)
  }

  emit('refresh')
}

function patUsageGuideDone() {
  useDialogVisible.value = false
}

function abbreviateToken(token: string): string {
  if (token.length > 8) {
    return `${token.slice(0, 4)}...${token.slice(-4)}`
  } else {
    return token
  }
}

function timestampToDisplayString(timestamp: string): string {
  return dayjs(timestamp).format('YYYY-MM-DD HH:mm')
}

</script>