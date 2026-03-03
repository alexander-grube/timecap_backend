import axios from 'axios'

const api = axios.create({
  baseURL: '/bugtracker',
  headers: {
    'Content-Type': 'application/json',
  },
})

// Attach JWT token to every request if available
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Handle 401/403 responses globally
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && (error.response.status === 401 || error.response.status === 403)) {
      localStorage.removeItem('token')
      localStorage.removeItem('account')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

// ── Auth ──
export const login = (email, password) =>
  api.post('/users/login', { account: { email, password } })

export const signUp = (firstname, lastname, email, password, organizationID) =>
  api.post('/users', {
    account: { firstname, lastname, email, password, organizationID },
  })

// ── Account ──
export const getCurrentAccount = () => api.get('/account')

export const getAccountByID = (id) => api.get(`/account/${id}`)

export const updateAccount = (data) =>
  api.put('/account', { account: data })

export const deleteAccount = () => api.delete('/account')

// ── Tickets ──
export const getAllTickets = () => api.get('/ticket')

export const getTicketByID = (id) => api.get(`/ticket/${id}`)

export const createTicket = (ticket) =>
  api.post('/ticket/new', { ticket })

export const updateTicket = (id, ticket) =>
  api.put(`/ticket/${id}`, { ticket: { id, ...ticket } })

// ── Projects ──
export const createProject = (project) =>
  api.post('/project/new', { project })

export default api
