import { defineStore } from 'pinia'
import { ref } from 'vue'
import {
  getAllTickets as apiGetAll,
  getTicketByID as apiGetByID,
  createTicket as apiCreate,
  updateTicket as apiUpdate,
} from '../api'

export const useTicketStore = defineStore('tickets', () => {
  const tickets = ref([])
  const currentTicket = ref(null)
  const loading = ref(false)

  const priorityLabels = { 1: 'Low', 2: 'Medium', 3: 'High' }
  const priorityColors = { 1: 'success', 2: 'warning', 3: 'error' }
  const typeLabels = { 1: 'Other', 2: 'Bug', 3: 'Feature' }
  const typeColors = { 1: 'grey', 2: 'error', 3: 'info' }
  const statusLabels = { 1: 'New', 2: 'Open', 3: 'In Progress', 4: 'Resolved' }
  const statusColors = { 1: 'info', 2: 'primary', 3: 'warning', 4: 'success' }

  async function fetchAll() {
    loading.value = true
    try {
      const { data } = await apiGetAll()
      tickets.value = data.ticket || []
    } finally {
      loading.value = false
    }
  }

  async function fetchByID(id) {
    loading.value = true
    try {
      const { data } = await apiGetByID(id)
      currentTicket.value = data.ticket
      return data.ticket
    } finally {
      loading.value = false
    }
  }

  async function create(ticket) {
    const { data } = await apiCreate(ticket)
    await fetchAll()
    return data.ticket
  }

  async function update(id, ticket) {
    const { data } = await apiUpdate(id, ticket)
    await fetchAll()
    return data.ticket
  }

  return {
    tickets,
    currentTicket,
    loading,
    priorityLabels,
    priorityColors,
    typeLabels,
    typeColors,
    statusLabels,
    statusColors,
    fetchAll,
    fetchByID,
    create,
    update,
  }
})
