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

import { useAuthUserStore } from '../stores/authUser'
import { api } from '../adminApi'

const authUserStore = useAuthUserStore()

const description = ref('')
const email = authUserStore.user!.email!
const isFormValid = ref(false)

const router = useRouter()

function validateDescription(description: string): boolean {
    return Boolean(description)
}

async function generate() {

  try {
    const createTokenResponse = await api.createToken()
    const updateTokenResponse = await api.updateToken(createTokenResponse.data.token!, {
      description: description.value
    })
  } catch (error) {
    console.log(error)
  }

  router.push({ name: 'pats' })
}

</script>
