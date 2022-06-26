import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import PATsView from '../views/PATsView.vue'
import NewPATView from '../views/NewPATView.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    redirect: '/pats'
  },
  {
    path: '/pats',
    name: 'pats',
    component: PATsView
  },
  {
    path: '/pats/new',
    name: 'new-pat',
    component: NewPATView
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
