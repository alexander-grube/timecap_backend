<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-card class="elevation-12 rounded-lg">
          <v-toolbar color="primary" flat>
            <v-toolbar-title>
              <v-icon start>mdi-bug</v-icon>
              Bug Tracker — Login
            </v-toolbar-title>
          </v-toolbar>

          <v-card-text class="pa-6">
            <v-alert v-if="error" type="error" variant="tonal" class="mb-4" closable @click:close="error = ''">
              {{ error }}
            </v-alert>

            <v-form ref="form" v-model="valid" @submit.prevent="handleLogin" fast-fail>
              <v-text-field
                v-model="email"
                label="Email"
                prepend-inner-icon="mdi-email"
                type="email"
                :rules="emailRules"
                variant="outlined"
                class="mb-2"
                required
              />
              <v-text-field
                v-model="password"
                label="Password"
                prepend-inner-icon="mdi-lock"
                :type="showPassword ? 'text' : 'password'"
                :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                @click:append-inner="showPassword = !showPassword"
                :rules="passwordRules"
                variant="outlined"
                class="mb-2"
                required
              />
              <v-btn
                type="submit"
                color="primary"
                size="large"
                block
                :loading="loading"
                :disabled="!valid"
              >
                Sign In
              </v-btn>
            </v-form>
          </v-card-text>

          <v-card-actions class="px-6 pb-4">
            <v-spacer />
            <span class="text-body-2">Don't have an account?</span>
            <v-btn variant="text" color="primary" :to="{ name: 'SignUp' }">
              Sign Up
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const valid = ref(false)
const loading = ref(false)
const error = ref('')
const showPassword = ref(false)

const emailRules = [
  (v) => !!v || 'Email is required',
  (v) => /.+@.+\..+/.test(v) || 'Must be a valid email',
]
const passwordRules = [(v) => !!v || 'Password is required']

async function handleLogin() {
  loading.value = true
  error.value = ''
  try {
    await authStore.login(email.value, password.value)
    router.push({ name: 'Dashboard' })
  } catch (err) {
    error.value = err.response?.data?.errors?.body?.[0] || 'Invalid email or password'
  } finally {
    loading.value = false
  }
}
</script>
