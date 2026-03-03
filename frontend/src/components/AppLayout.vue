<template>
  <div>
    <!-- App Bar -->
    <v-app-bar color="surface" flat border="b">
      <v-app-bar-nav-icon @click="drawer = !drawer" />
      <v-toolbar-title class="font-weight-bold">
        <v-icon start color="primary">mdi-bug</v-icon>
        Bug Tracker
      </v-toolbar-title>
      <v-spacer />

      <!-- Theme Toggle -->
      <v-btn icon @click="toggleTheme">
        <v-icon>{{ isDark ? 'mdi-weather-sunny' : 'mdi-weather-night' }}</v-icon>
      </v-btn>

      <!-- Account Menu -->
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" variant="text" class="ml-2">
            <v-avatar color="primary" size="32" class="mr-2">
              <span class="text-caption">{{ initials }}</span>
            </v-avatar>
            {{ authStore.fullName }}
            <v-icon end>mdi-chevron-down</v-icon>
          </v-btn>
        </template>
        <v-list density="compact" nav>
          <v-list-item :to="{ name: 'Account' }" prepend-icon="mdi-account-circle">
            <v-list-item-title>My Account</v-list-item-title>
          </v-list-item>
          <v-divider />
          <v-list-item @click="handleLogout" prepend-icon="mdi-logout" class="text-error">
            <v-list-item-title>Logout</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <!-- Navigation Drawer -->
    <v-navigation-drawer v-model="drawer" :rail="rail" permanent @click="rail = false">
      <template v-slot:prepend>
        <v-list-item
          :prepend-avatar="undefined"
          class="pa-4"
        >
          <template v-slot:prepend>
            <v-avatar color="primary">
              <span>{{ initials }}</span>
            </v-avatar>
          </template>
          <v-list-item-title class="font-weight-bold">{{ authStore.fullName }}</v-list-item-title>
          <v-list-item-subtitle>{{ authStore.roleLabel }}</v-list-item-subtitle>
          <template v-slot:append>
            <v-btn icon variant="text" size="small" @click.stop="rail = !rail">
              <v-icon>{{ rail ? 'mdi-chevron-right' : 'mdi-chevron-left' }}</v-icon>
            </v-btn>
          </template>
        </v-list-item>
      </template>

      <v-divider />

      <v-list nav density="comfortable">
        <v-list-item
          v-for="item in navItems"
          :key="item.title"
          :to="item.to"
          :prepend-icon="item.icon"
          :title="item.title"
          :value="item.title"
          color="primary"
          rounded="lg"
        />
      </v-list>

      <v-divider />

      <v-list nav density="comfortable">
        <v-list-subheader v-if="!rail">ACTIONS</v-list-subheader>
        <v-list-item
          v-for="item in actionItems"
          :key="item.title"
          :to="item.to"
          :prepend-icon="item.icon"
          :title="item.title"
          :value="item.title"
          color="primary"
          rounded="lg"
        />
      </v-list>
    </v-navigation-drawer>

    <!-- Main Content -->
    <v-main>
      <router-view />
    </v-main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useTheme } from 'vuetify'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const theme = useTheme()

const drawer = ref(true)
const rail = ref(false)

const isDark = computed(() => theme.global.current.value.dark)
const initials = computed(() => {
  const f = authStore.account?.firstname?.[0] || ''
  const l = authStore.account?.lastname?.[0] || ''
  return (f + l).toUpperCase()
})

const navItems = [
  { title: 'Dashboard', icon: 'mdi-view-dashboard', to: { name: 'Dashboard' } },
  { title: 'Tickets', icon: 'mdi-ticket-outline', to: { name: 'Tickets' } },
  { title: 'Account', icon: 'mdi-account-circle', to: { name: 'Account' } },
]

const actionItems = [
  { title: 'New Ticket', icon: 'mdi-plus-circle', to: { name: 'CreateTicket' } },
  { title: 'New Project', icon: 'mdi-folder-plus', to: { name: 'CreateProject' } },
]

function toggleTheme() {
  theme.global.name.value = isDark.value ? 'light' : 'dark'
}

function handleLogout() {
  authStore.logout()
  router.push({ name: 'Login' })
}
</script>
