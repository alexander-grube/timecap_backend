<template>
  <v-container fluid>
    <v-row class="mb-4" align="center">
      <v-col>
        <h1 class="text-h4 font-weight-bold">Tickets</h1>
      </v-col>
      <v-col cols="auto">
        <v-btn color="primary" :to="{ name: 'CreateTicket' }">
          <v-icon start>mdi-plus</v-icon>
          New Ticket
        </v-btn>
      </v-col>
    </v-row>

    <!-- Filters -->
    <v-row class="mb-2">
      <v-col cols="12" sm="4">
        <v-text-field
          v-model="search"
          label="Search tickets..."
          prepend-inner-icon="mdi-magnify"
          variant="outlined"
          density="compact"
          clearable
          hide-details
        />
      </v-col>
      <v-col cols="6" sm="3">
        <v-select
          v-model="filterStatus"
          :items="statusOptions"
          label="Status"
          variant="outlined"
          density="compact"
          clearable
          hide-details
        />
      </v-col>
      <v-col cols="6" sm="3">
        <v-select
          v-model="filterPriority"
          :items="priorityOptions"
          label="Priority"
          variant="outlined"
          density="compact"
          clearable
          hide-details
        />
      </v-col>
      <v-col cols="6" sm="2">
        <v-select
          v-model="filterType"
          :items="typeOptions"
          label="Type"
          variant="outlined"
          density="compact"
          clearable
          hide-details
        />
      </v-col>
    </v-row>

    <v-progress-linear v-if="ticketStore.loading" indeterminate color="primary" class="mb-4" />

    <!-- Ticket Table -->
    <v-card class="rounded-lg">
      <v-data-table
        :headers="headers"
        :items="filteredTickets"
        :search="search"
        hover
        class="cursor-pointer"
        @click:row="(_, { item }) => goToTicket(item)"
      >
        <template v-slot:item.type="{ item }">
          <v-chip :color="ticketStore.typeColors[item.type]" size="small" label>
            {{ ticketStore.typeLabels[item.type] }}
          </v-chip>
        </template>
        <template v-slot:item.priority="{ item }">
          <v-chip :color="ticketStore.priorityColors[item.priority]" size="small" label>
            {{ ticketStore.priorityLabels[item.priority] }}
          </v-chip>
        </template>
        <template v-slot:item.status="{ item }">
          <v-chip :color="ticketStore.statusColors[item.status]" size="small" label>
            {{ ticketStore.statusLabels[item.status] }}
          </v-chip>
        </template>
        <template v-slot:item.description="{ item }">
          {{ item.description?.substring(0, 60) }}{{ item.description?.length > 60 ? '...' : '' }}
        </template>
        <template v-slot:no-data>
          <div class="text-center py-8">
            <v-icon size="48" class="mb-2 text-medium-emphasis">mdi-ticket-outline</v-icon>
            <div class="text-medium-emphasis">No tickets found</div>
          </div>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTicketStore } from '../stores/tickets'

const router = useRouter()
const ticketStore = useTicketStore()

const search = ref('')
const filterStatus = ref(null)
const filterPriority = ref(null)
const filterType = ref(null)

const statusOptions = [
  { title: 'New', value: 1 },
  { title: 'Open', value: 2 },
  { title: 'In Progress', value: 3 },
  { title: 'Resolved', value: 4 },
]
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

const headers = [
  { title: 'ID', key: 'id', width: '80px' },
  { title: 'Topic', key: 'topic' },
  { title: 'Description', key: 'description', sortable: false },
  { title: 'Type', key: 'type', width: '120px' },
  { title: 'Priority', key: 'priority', width: '120px' },
  { title: 'Status', key: 'status', width: '140px' },
]

const filteredTickets = computed(() => {
  let result = ticketStore.tickets
  if (filterStatus.value) result = result.filter((t) => t.status === filterStatus.value)
  if (filterPriority.value) result = result.filter((t) => t.priority === filterPriority.value)
  if (filterType.value) result = result.filter((t) => t.type === filterType.value)
  return result
})

function goToTicket(item) {
  router.push({ name: 'TicketDetail', params: { id: item.id } })
}

onMounted(() => {
  ticketStore.fetchAll()
})
</script>

<style scoped>
.cursor-pointer :deep(tbody tr) {
  cursor: pointer;
}
</style>
