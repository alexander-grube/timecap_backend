import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as apiLogin, signUp as apiSignUp, getCurrentAccount, updateAccount as apiUpdateAccount, deleteAccount as apiDeleteAccount } from '../api'

export const useAuthStore = defineStore('auth', () => {
  const account = ref(JSON.parse(localStorage.getItem('account') || 'null'))
  const token = ref(localStorage.getItem('token') || '')

  const isAuthenticated = computed(() => !!token.value)

  const roleLabel = computed(() => {
    if (!account.value) return ''
    const roles = { 1: 'Admin', 2: 'Developer', 3: 'Manager', 4: 'User' }
    return roles[account.value.role] || 'Unknown'
  })

  const fullName = computed(() => {
    if (!account.value) return ''
    return `${account.value.firstname} ${account.value.lastname}`
  })

  async function login(email, password) {
    const { data } = await apiLogin(email, password)
    const acc = data.account
    token.value = acc.token
    account.value = acc
    localStorage.setItem('token', acc.token)
    localStorage.setItem('account', JSON.stringify(acc))
    return acc
  }

  async function signUp(firstname, lastname, email, password, organizationID) {
    const { data } = await apiSignUp(firstname, lastname, email, password, organizationID)
    const acc = data.account
    token.value = acc.token
    account.value = acc
    localStorage.setItem('token', acc.token)
    localStorage.setItem('account', JSON.stringify(acc))
    return acc
  }

  async function fetchCurrentAccount() {
    const { data } = await getCurrentAccount()
    account.value = { ...account.value, ...data.account }
    localStorage.setItem('account', JSON.stringify(account.value))
    return data.account
  }

  async function updateProfile(updates) {
    const { data } = await apiUpdateAccount(updates)
    account.value = { ...account.value, ...data.account }
    localStorage.setItem('account', JSON.stringify(account.value))
    return data.account
  }

  async function removeAccount() {
    await apiDeleteAccount()
    logout()
  }

  function logout() {
    token.value = ''
    account.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('account')
  }

  return {
    account,
    token,
    isAuthenticated,
    roleLabel,
    fullName,
    login,
    signUp,
    fetchCurrentAccount,
    updateProfile,
    removeAccount,
    logout,
  }
})
