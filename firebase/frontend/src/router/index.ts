import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import PATsView from '../views/PATsView.vue'
import NewPATView from '../views/NewPATView.vue'
import StoresView from '../views/StoresView.vue'
import StoreFilesView from '../views/StoreFilesView.vue'

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
  {
    path: '/stores',
    name: 'stores',
    component: StoresView
  },
  {
    path: '/stores/:store/files',
    name: 'storeFiles',
    component: StoreFilesView
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
