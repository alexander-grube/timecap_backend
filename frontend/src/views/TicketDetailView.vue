<template>
  <v-container fluid>
    <v-row class="mb-4">
      <v-col>
        <v-btn variant="text" :to="{ name: 'Tickets' }" class="mb-2">
          <v-icon start>mdi-arrow-left</v-icon>
          Back to Tickets
        </v-btn>
        <h1 class="text-h4 font-weight-bold">Ticket #{{ id }}</h1>
      </v-col>
    </v-row>

    <v-progress-linear v-if="ticketStore.loading" indeterminate color="primary" class="mb-4" />

    <v-row v-if="ticket">
      <v-col cols="12" md="8">
        <v-card class="rounded-lg">
          <v-card-title class="text-h5">{{ ticket.topic }}</v-card-title>
          <v-divider />
          <v-card-text>
            <h3 class="text-subtitle-1 font-weight-bold mb-2">Description</h3>
            <div class="text-body-1" style="white-space: pre-wrap;">{{ ticket.description }}</div>
          </v-card-text>
        </v-card>

        <!-- Edit Form -->
        <v-card class="rounded-lg mt-4">
          <v-card-title>
            <v-icon start>mdi-pencil</v-icon>
            Update Ticket
          </v-card-title>
          <v-divider />
          <v-card-text>
            <v-alert v-if="error" type="error" variant="tonal" class="mb-4" closable>{{ error }}</v-alert>
            <v-alert v-if="success" type="success" variant="tonal" class="mb-4" closable>{{ success }}</v-alert>

            <v-form @submit.prevent="handleUpdate">
              <v-text-field
                v-model="editForm.topic"
                label="Topic"
                variant="outlined"
                class="mb-2"
              />
              <v-textarea
                v-model="editForm.description"
                label="Description"
                variant="outlined"
                rows="4"
                class="mb-2"
              />
              <v-row>
                <v-col cols="4">
                  <v-select
                    v-model="editForm.priority"
                    :items="priorityOptions"
                    label="Priority"
                    variant="outlined"
                  />
                </v-col>
                <v-col cols="4">
                  <v-select
                    v-model="editForm.type"
                    :items="typeOptions"
                    label="Type"
                    variant="outlined"
                  />
                </v-col>
                <v-col cols="4">
                  <v-select
                    v-model="editForm.status"
                    :items="statusOptions"
                    label="Status"
                    variant="outlined"
                  />
                </v-col>
              </v-row>
              <v-btn type="submit" color="primary" :loading="updating">
                <v-icon start>mdi-content-save</v-icon>
                Save Changes
              </v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- Sidebar -->
      <v-col cols="12" md="4">
        <v-card class="rounded-lg">
          <v-card-title>Details</v-card-title>
          <v-divider />
          <v-list>
            <v-list-item>
              <template v-slot:prepend>
                <v-icon>mdi-flag</v-icon>
              </template>
              <v-list-item-title>Priority</v-list-item-title>
              <template v-slot:append>
                <v-chip :color="ticketStore.priorityColors[ticket.priority]" size="small" label>
                  {{ ticketStore.priorityLabels[ticket.priority] }}
                </v-chip>
              </template>
            </v-list-item>
            <v-list-item>
              <template v-slot:prepend>
                <v-icon>mdi-shape</v-icon>
              </template>
              <v-list-item-title>Type</v-list-item-title>
              <template v-slot:append>
                <v-chip :color="ticketStore.typeColors[ticket.type]" size="small" label>
                  {{ ticketStore.typeLabels[ticket.type] }}
                </v-chip>
              </template>
            </v-list-item>
            <v-list-item>
              <template v-slot:prepend>
                <v-icon>mdi-list-status</v-icon>
              </template>
              <v-list-item-title>Status</v-list-item-title>
              <template v-slot:append>
                <v-chip :color="ticketStore.statusColors[ticket.status]" size="small" label>
                  {{ ticketStore.statusLabels[ticket.status] }}
                </v-chip>
              </template>
            </v-list-item>
          </v-list>
        </v-card>

        <!-- Assigned Account -->
        <v-card v-if="ticket.account" class="rounded-lg mt-4">
          <v-card-title>Assigned To</v-card-title>
          <v-divider />
          <v-card-text class="d-flex align-center">
            <v-avatar color="primary" class="mr-3">
              <span class="text-h6">{{ ticket.account.firstname?.[0] }}{{ ticket.account.lastname?.[0] }}</span>
            </v-avatar>
            <div>
              <div class="font-weight-medium">{{ ticket.account.firstname }} {{ ticket.account.lastname }}</div>
              <div class="text-caption text-medium-emphasis">{{ ticket.account.email }}</div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-card v-else-if="!ticketStore.loading" class="rounded-lg text-center py-8">
      <v-icon size="64" class="text-medium-emphasis mb-4">mdi-ticket-outline</v-icon>
      <div class="text-h6 text-medium-emphasis">Ticket not found</div>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue'
import { useTicketStore } from '../stores/tickets'

const props = defineProps({
  id: { type: [String, Number], required: true },
})

const ticketStore = useTicketStore()
const ticket = ref(null)
const error = ref('')
const success = ref('')
const updating = ref(false)

const editForm = reactive({
  topic: '',
  description: '',
  priority: 1,
  type: 1,
  status: 1,
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

async function loadTicket() {
  try {
    const t = await ticketStore.fetchByID(props.id)
    ticket.value = t
    editForm.topic = t.topic
    editForm.description = t.description
    editForm.priority = t.priority
    editForm.type = t.type
    editForm.status = t.status
  } catch (e) {
    error.value = 'Failed to load ticket'
  }
}

async function handleUpdate() {
  updating.value = true
  error.value = ''
  success.value = ''
  try {
    await ticketStore.update(Number(props.id), {
      topic: editForm.topic,
      description: editForm.description,
      priority: editForm.priority,
      type: editForm.type,
      status: editForm.status,
    })
    success.value = 'Ticket updated successfully'
    await loadTicket()
  } catch (e) {
    error.value = e.response?.data?.errors?.body?.[0] || 'Failed to update ticket'
  } finally {
    updating.value = false
  }
}

onMounted(loadTicket)
watch(() => props.id, loadTicket)
</script>
