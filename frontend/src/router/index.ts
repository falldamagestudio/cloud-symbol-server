import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import SessionsView from '../views/SessionsView.vue'
import SessionView from '../views/SessionView.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    redirect: '/sessions'
  },
  {
    path: '/sessions',
    name: 'sessions',
    component: SessionsView
  },
  {
    path: '/sessions/:id',
    name: 'session',
    component: SessionView
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
