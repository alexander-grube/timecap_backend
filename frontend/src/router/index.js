import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/LoginView.vue'),
    meta: { guest: true },
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: () => import('../views/SignUpView.vue'),
    meta: { guest: true },
  },
  {
    path: '/',
    component: () => import('../components/AppLayout.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('../views/DashboardView.vue'),
      },
      {
        path: 'tickets',
        name: 'Tickets',
        component: () => import('../views/TicketsView.vue'),
      },
      {
        path: 'tickets/new',
        name: 'CreateTicket',
        component: () => import('../views/CreateTicketView.vue'),
      },
      {
        path: 'tickets/:id',
        name: 'TicketDetail',
        component: () => import('../views/TicketDetailView.vue'),
        props: true,
      },
      {
        path: 'projects/new',
        name: 'CreateProject',
        component: () => import('../views/CreateProjectView.vue'),
      },
      {
        path: 'account',
        name: 'Account',
        component: () => import('../views/AccountView.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    next({ name: 'Login' })
  } else if (to.meta.guest && token) {
    next({ name: 'Dashboard' })
  } else {
    next()
  }
})

export default router
