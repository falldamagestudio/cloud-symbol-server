<template>
  <v-navigation-drawer app clipped>
    <v-list>
      <v-list-item
        v-for="item in items"
        :key="item.title"
        link
        :to="item.route"
      >
        <v-list-item-content>
          {{item.title}}
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script setup lang="ts">

interface NavigationDrawerEntry {
  title: string,
  route: { 
    name: string
  },
}

const items: NavigationDrawerEntry[] = [
  {
    title: 'Stores',
    route: { name: 'stores' },
  },
  {
    title: 'My Access Tokens',
    route: { name: 'pats' },
  }
]

</script>