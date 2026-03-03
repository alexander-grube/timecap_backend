<template>
  <v-container fluid>
    <v-row class="mb-4">
      <v-col>
        <v-btn variant="text" :to="{ name: 'Dashboard' }" class="mb-2">
          <v-icon start>mdi-arrow-left</v-icon>
          Back to Dashboard
        </v-btn>
        <h1 class="text-h4 font-weight-bold">Create New Project</h1>
      </v-col>
    </v-row>

    <v-row justify="center">
      <v-col cols="12" md="6">
        <v-card class="rounded-lg">
          <v-card-text class="pa-6">
            <v-alert v-if="error" type="error" variant="tonal" class="mb-4" closable>{{ error }}</v-alert>
            <v-alert v-if="success" type="success" variant="tonal" class="mb-4" closable>{{ success }}</v-alert>

            <v-form ref="form" v-model="valid" @submit.prevent="handleCreate">
              <v-text-field
                v-model="project.name"
                label="Project Name"
                prepend-inner-icon="mdi-folder"
                :rules="[(v) => !!v || 'Name is required']"
                variant="outlined"
                class="mb-2"
                required
              />

              <v-textarea
                v-model="project.description"
                label="Description"
                prepend-inner-icon="mdi-text-long"
                :rules="[(v) => !!v || 'Description is required']"
                variant="outlined"
                rows="4"
                class="mb-2"
                required
              />

              <v-btn
                type="submit"
                color="primary"
                size="large"
                :loading="loading"
                :disabled="!valid"
              >
                <v-icon start>mdi-plus</v-icon>
                Create Project
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { createProject } from '../api'

const valid = ref(false)
const loading = ref(false)
const error = ref('')
const success = ref('')

const project = reactive({
  name: '',
  description: '',
})

async function handleCreate() {
  loading.value = true
  error.value = ''
  success.value = ''
  try {
    await createProject({ name: project.name, description: project.description })
    success.value = `Project "${project.name}" created successfully!`
    project.name = ''
    project.description = ''
  } catch (err) {
    error.value = err.response?.data?.errors?.body?.[0] || 'Failed to create project'
  } finally {
    loading.value = false
  }
}
</script>
