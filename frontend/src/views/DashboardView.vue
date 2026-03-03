<template>
  <v-container fluid>
    <v-row class="mb-4">
      <v-col>
        <h1 class="text-h4 font-weight-bold">Dashboard</h1>
        <p class="text-subtitle-1 text-medium-emphasis">Welcome back, {{ authStore.fullName }}</p>
      </v-col>
    </v-row>

    <!-- Stats Cards -->
    <v-row>
      <v-col cols="12" sm="6" md="3">
        <v-card color="primary" variant="tonal" class="rounded-lg">
          <v-card-text class="d-flex align-center">
            <v-icon size="48" class="mr-4">mdi-ticket-outline</v-icon>
            <div>
              <div class="text-h4 font-weight-bold">{{ tickets.length }}</div>
              <div class="text-subtitle-2">Total Tickets</div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="6" md="3">
        <v-card color="info" variant="tonal" class="rounded-lg">
          <v-card-text class="d-flex align-center">
            <v-icon size="48" class="mr-4">mdi-new-box</v-icon>
            <div>
              <div class="text-h4 font-weight-bold">{{ newCount }}</div>
              <div class="text-subtitle-2">New</div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="6" md="3">
        <v-card color="warning" variant="tonal" class="rounded-lg">
          <v-card-text class="d-flex align-center">
            <v-icon size="48" class="mr-4">mdi-progress-clock</v-icon>
            <div>
              <div class="text-h4 font-weight-bold">{{ inProgressCount }}</div>
              <div class="text-subtitle-2">In Progress</div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" sm="6" md="3">
        <v-card color="success" variant="tonal" class="rounded-lg">
          <v-card-text class="d-flex align-center">
            <v-icon size="48" class="mr-4">mdi-check-circle</v-icon>
            <div>
              <div class="text-h4 font-weight-bold">{{ resolvedCount }}</div>
              <div class="text-subtitle-2">Resolved</div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Recent Tickets -->
    <v-row class="mt-4">
      <v-col>
        <v-card class="rounded-lg">
          <v-card-title class="d-flex align-center">
            <v-icon start>mdi-clock-outline</v-icon>
            Recent Tickets
            <v-spacer />
            <v-btn variant="text" color="primary" :to="{ name: 'Tickets' }">
              View All
              <v-icon end>mdi-arrow-right</v-icon>
            </v-btn>
          </v-card-title>
          <v-divider />

          <v-progress-linear v-if="ticketStore.loading" indeterminate color="primary" />

          <v-list v-if="recentTickets.length" lines="two">
            <v-list-item
              v-for="ticket in recentTickets"
              :key="ticket.id"
              :to="{ name: 'TicketDetail', params: { id: ticket.id } }"
              class="py-3"
            >
              <template v-slot:prepend>
                <v-avatar :color="ticketStore.typeColors[ticket.type]" size="40">
                  <v-icon color="white" size="20">
                    {{ ticket.type === 2 ? 'mdi-bug' : ticket.type === 3 ? 'mdi-lightbulb' : 'mdi-help-circle' }}
                  </v-icon>
                </v-avatar>
              </template>
              <v-list-item-title class="font-weight-medium">{{ ticket.topic }}</v-list-item-title>
              <v-list-item-subtitle>{{ ticket.description?.substring(0, 80) }}{{ ticket.description?.length > 80 ? '...' : '' }}</v-list-item-subtitle>
              <template v-slot:append>
                <div class="d-flex ga-2">
                  <v-chip :color="ticketStore.priorityColors[ticket.priority]" size="small" label>
                    {{ ticketStore.priorityLabels[ticket.priority] }}
                  </v-chip>
                  <v-chip :color="ticketStore.statusColors[ticket.status]" size="small" label>
                    {{ ticketStore.statusLabels[ticket.status] }}
                  </v-chip>
                </div>
              </template>
            </v-list-item>
          </v-list>

          <v-card-text v-else-if="!ticketStore.loading" class="text-center text-medium-emphasis py-8">
            <v-icon size="48" class="mb-2">mdi-ticket-outline</v-icon>
            <div>No tickets yet. Create your first ticket!</div>
            <v-btn color="primary" variant="tonal" class="mt-4" :to="{ name: 'CreateTicket' }">
              <v-icon start>mdi-plus</v-icon>
              New Ticket
            </v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useTicketStore } from '../stores/tickets'

const authStore = useAuthStore()
const ticketStore = useTicketStore()

const tickets = computed(() => ticketStore.tickets)
const recentTickets = computed(() => [...tickets.value].slice(-5).reverse())

const newCount = computed(() => tickets.value.filter((t) => t.status === 1).length)
const inProgressCount = computed(() => tickets.value.filter((t) => t.status === 3).length)
const resolvedCount = computed(() => tickets.value.filter((t) => t.status === 4).length)

onMounted(() => {
  ticketStore.fetchAll()
})
</script>
