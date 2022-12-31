<template>
  <div>

    <v-form
      v-model="isFormValid"
    >

      <h1>Generate new Personal Access Token</h1>

      <p>Personal Access Tokens are used to fetch files from the symbol store.</p>

      <v-text-field
        label="Description"
        v-model="description"
        :rules="[validateDescription]"
      >
      </v-text-field>

      <v-btn
        :disabled="!isFormValid"
        v-on:click="generate()"
      >
        Generate
      </v-btn>

      <v-btn
        :to="{ name: 'pats' }"
      >
        Cancel
      </v-btn>

    </v-form>

  </div>
</template>

<script setup lang="ts">

import { ref } from 'vue'
import { useRouter } from 'vue-router/composables'
import { doc, serverTimestamp, setDoc }  from 'firebase/firestore'

import { useAuthUserStore } from '../stores/authUser'
import { db } from '../firebase'

const authUserStore = useAuthUserStore()

const description = ref('')
const email = authUserStore.user!.email!
const isFormValid = ref(false)

const router = useRouter()

// dec2hex :: Integer -> String
// i.e. 0-255 -> '00'-'ff'
function dec2hex (dec: number) : string {
  return dec.toString(16).padStart(2, "0")
}

// generateId :: Integer -> String
function generateId (len: number): string {
  const arr = new Uint8Array((len || 40) / 2)
  window.crypto.getRandomValues(arr)
  return Array.from(arr, dec2hex).join('')
}

function validateDescription(description: string): boolean {
    return Boolean(description)
}

async function generate() {
  const id = generateId(32)

  const patFields = {
    description: description.value,
    creationTimestamp: serverTimestamp()
  }

  const patDocRef = doc(db, 'users', email, 'pats', id)
  await setDoc(patDocRef, patFields)
  router.push({ name: 'pats' })
}

</script>
