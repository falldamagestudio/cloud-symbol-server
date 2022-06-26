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

<script lang="ts">
import Vue from 'vue';
import firebase from 'firebase/app'
import store from '../store/index'
import { db } from '../firebase'

interface Data {
  description: string
  email: string
  isFormValid: boolean
}

// dec2hex :: Integer -> String
// i.e. 0-255 -> '00'-'ff'
function dec2hex (dec: number) : string {
  return dec.toString(16).padStart(2, "0")
}

// generateId :: Integer -> String
function generateId (len: number): string {
  var arr = new Uint8Array((len || 40) / 2)
  window.crypto.getRandomValues(arr)
  return Array.from(arr, dec2hex).join('')
}

export default Vue.extend({

  components: {
  },


  methods: {
    validateDescription(description: string): boolean {
        return Boolean(description)
    },

    generate() {
      const id = generateId(32)

      const patFields = {
        description: this.description,
        creationTimestamp: firebase.firestore.FieldValue.serverTimestamp()
      }

      db.collection('users').doc(this.email).collection('pats').doc(id).set(patFields).then(() => {
        this.$router.push({ name: 'pats' })
      })
    },
  },

  data () : Data {
    return {
      description: '',
      email: store.state.user!.email!,
      isFormValid: false,
    } 
  },

})
</script>
