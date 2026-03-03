<template>
  <v-container fluid>
    <v-row class="mb-4">
      <v-col>
        <v-btn variant="text" :to="{ name: 'Tickets' }" class="mb-2">
          <v-icon start>mdi-arrow-left</v-icon>
          Back to Tickets
        </v-btn>
        <h1 class="text-h4 font-weight-bold">Create New Ticket</h1>
      </v-col>
    </v-row>

    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-card class="rounded-lg">
          <v-card-text class="pa-6">
            <v-alert v-if="error" type="error" variant="tonal" class="mb-4" closable>{{ error }}</v-alert>

            <v-form ref="form" v-model="valid" @submit.prevent="handleCreate">
              <v-text-field
                v-model="ticket.topic"
                label="Topic"
                prepend-inner-icon="mdi-text-short"
                :rules="[(v) => !!v || 'Topic is required']"
                variant="outlined"
                class="mb-2"
                required
              />

              <v-textarea
                v-model="ticket.description"
                label="Description"
                prepend-inner-icon="mdi-text-long"
                :rules="[(v) => !!v || 'Description is required']"
                variant="outlined"
                rows="5"
                class="mb-2"
                required
              />

              <v-row>
                <v-col cols="12" sm="4">
                  <v-select
                    v-model="ticket.priority"
                    :items="priorityOptions"
                    label="Priority"
                    prepend-inner-icon="mdi-flag"
                    :rules="[(v) => !!v || 'Priority is required']"
                    variant="outlined"
                    required
                  />
                </v-col>
                <v-col cols="12" sm="4">
                  <v-select
                    v-model="ticket.type"
                    :items="typeOptions"
                    label="Type"
                    prepend-inner-icon="mdi-shape"
                    :rules="[(v) => !!v || 'Type is required']"
                    variant="outlined"
                    required
                  />
                </v-col>
                <v-col cols="12" sm="4">
                  <v-select
                    v-model="ticket.status"
                    :items="statusOptions"
                    label="Status"
                    prepend-inner-icon="mdi-list-status"
                    :rules="[(v) => !!v || 'Status is required']"
                    variant="outlined"
                    required
                  />
                </v-col>
              </v-row>

              <v-text-field
                v-model.number="ticket.project_id"
                label="Project ID"
                prepend-inner-icon="mdi-folder"
                type="number"
                :rules="[(v) => !!v || 'Project ID is required']"
                variant="outlined"
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
                Create Ticket
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
import { useRouter } from 'vue-router'
import { useTicketStore } from '../stores/tickets'

const router = useRouter()
const ticketStore = useTicketStore()

const valid = ref(false)
const loading = ref(false)
const error = ref('')

const ticket = reactive({
  topic: '',
  description: '',
  priority: 2,
  type: 1,
  status: 1,
  project_id: 1,
})

const priorityOptions = [
  { title: 'Low', value: 1 },
  { title: 'Medium', value: 2 },
  { title: 'High', value: 3 },
]
const typeOptions = [
  { title: 'Other', value: 1 },
  { title: 'Bug', value: 2 },
  { title: 'Feature', value: 3 },
]
const statusOptions = [
  { title: 'New', value: 1 },
  { title: 'Open', value: 2 },
  { title: 'In Progress', value: 3 },
  { title: 'Resolved', value: 4 },
]

async function handleCreate() {
  loading.value = true
  error.value = ''
  try {
    await ticketStore.create({ ...ticket })
    router.push({ name: 'Tickets' })
  } catch (err) {
    error.value = err.response?.data?.errors?.body?.[0] || 'Failed to create ticket'
  } finally {
    loading.value = false
  }
}
</script>
