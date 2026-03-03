<template>
  <v-container fluid>
    <v-row class="mb-4">
      <v-col>
        <h1 class="text-h4 font-weight-bold">My Account</h1>
      </v-col>
    </v-row>

    <v-row>
      <!-- Profile Card -->
      <v-col cols="12" md="4">
        <v-card class="rounded-lg">
          <v-card-text class="text-center pa-6">
            <v-avatar color="primary" size="96" class="mb-4">
              <span class="text-h3">{{ initials }}</span>
            </v-avatar>
            <h2 class="text-h5 font-weight-bold">{{ authStore.fullName }}</h2>
            <p class="text-medium-emphasis mb-2">{{ authStore.account?.email }}</p>
            <v-chip color="primary" label>{{ authStore.roleLabel }}</v-chip>
            <div v-if="authStore.account?.organization" class="mt-3 text-caption text-medium-emphasis">
              Organization #{{ authStore.account.organization }}
            </div>
          </v-card-text>
        </v-card>

        <!-- Danger Zone -->
        <v-card class="rounded-lg mt-4" color="error" variant="outlined">
          <v-card-title class="text-error">
            <v-icon start color="error">mdi-alert</v-icon>
            Danger Zone
          </v-card-title>
          <v-divider />
          <v-card-text>
            <p class="text-body-2 mb-4">Once you delete your account, there is no going back.</p>
            <v-btn color="error" variant="tonal" @click="confirmDelete = true" block>
              <v-icon start>mdi-delete</v-icon>
              Delete Account
            </v-btn>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Edit Profile -->
      <v-col cols="12" md="8">
        <v-card class="rounded-lg">
          <v-card-title>
            <v-icon start>mdi-pencil</v-icon>
            Edit Profile
          </v-card-title>
          <v-divider />
          <v-card-text class="pa-6">
            <v-alert v-if="error" type="error" variant="tonal" class="mb-4" closable>{{ error }}</v-alert>
            <v-alert v-if="success" type="success" variant="tonal" class="mb-4" closable>{{ success }}</v-alert>

            <v-form @submit.prevent="handleUpdate">
              <v-row>
                <v-col cols="6">
                  <v-text-field
                    v-model="form.firstname"
                    label="First Name"
                    prepend-inner-icon="mdi-account"
                    variant="outlined"
                  />
                </v-col>
                <v-col cols="6">
                  <v-text-field
                    v-model="form.lastname"
                    label="Last Name"
                    prepend-inner-icon="mdi-account"
                    variant="outlined"
                  />
                </v-col>
              </v-row>

              <v-text-field
                v-model="form.email"
                label="Email"
                prepend-inner-icon="mdi-email"
                type="email"
                variant="outlined"
                class="mb-2"
              />

              <v-text-field
                v-model="form.password"
                label="New Password (leave blank to keep current)"
                prepend-inner-icon="mdi-lock"
                :type="showPassword ? 'text' : 'password'"
                :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                @click:append-inner="showPassword = !showPassword"
                variant="outlined"
                class="mb-2"
              />

              <v-btn type="submit" color="primary" :loading="updating">
                <v-icon start>mdi-content-save</v-icon>
                Save Changes
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Delete Confirmation Dialog -->
    <v-dialog v-model="confirmDelete" max-width="400">
      <v-card class="rounded-lg">
        <v-card-title class="text-error">Delete Account</v-card-title>
        <v-card-text>
          Are you sure you want to permanently delete your account? This action cannot be undone.
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="confirmDelete = false">Cancel</v-btn>
          <v-btn color="error" variant="tonal" @click="handleDelete" :loading="deleting">
            Delete
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const error = ref('')
const success = ref('')
const updating = ref(false)
const deleting = ref(false)
const confirmDelete = ref(false)
const showPassword = ref(false)

const form = reactive({
  firstname: '',
  lastname: '',
  email: '',
  password: '',
})

const initials = computed(() => {
  const f = authStore.account?.firstname?.[0] || ''
  const l = authStore.account?.lastname?.[0] || ''
  return (f + l).toUpperCase()
})

onMounted(async () => {
  try {
    await authStore.fetchCurrentAccount()
  } catch { /* ignore */ }
  form.firstname = authStore.account?.firstname || ''
  form.lastname = authStore.account?.lastname || ''
  form.email = authStore.account?.email || ''
})

async function handleUpdate() {
  updating.value = true
  error.value = ''
  success.value = ''
  try {
    const updates = {}
    if (form.firstname) updates.firstname = form.firstname
    if (form.lastname) updates.lastname = form.lastname
    if (form.email) updates.email = form.email
    if (form.password) updates.password = form.password
    await authStore.updateProfile(updates)
    success.value = 'Profile updated successfully'
    form.password = ''
  } catch (err) {
    error.value = err.response?.data?.errors?.body?.[0] || 'Failed to update profile'
  } finally {
    updating.value = false
  }
}

async function handleDelete() {
  deleting.value = true
  try {
    await authStore.removeAccount()
    router.push({ name: 'Login' })
  } catch (err) {
    error.value = err.response?.data?.errors?.body?.[0] || 'Failed to delete account'
    confirmDelete.value = false
  } finally {
    deleting.value = false
  }
}
</script>
