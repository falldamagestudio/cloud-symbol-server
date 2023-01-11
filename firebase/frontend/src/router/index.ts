import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import PATsView from '../views/PATsView.vue'
import NewPATView from '../views/NewPATView.vue'
import StoresView from '../views/StoresView.vue'
import StoreFilesView from '../views/StoreFilesView.vue'
import StoreFileHashesView from '../views/StoreFileHashesView.vue'
import StoreUploadsView from '../views/StoreUploadsView.vue'
import StoreUploadView from '../views/StoreUploadView.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    redirect: '/stores'
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
  {
    path: '/stores/:store/files/:file',
    name: 'storeFileHashes',
    component: StoreFileHashesView
  },
  {
    path: '/stores/:store/uploads',
    name: 'storeUploads',
    component: StoreUploadsView
  },
  {
    path: '/stores/:store/uploads/:upload',
    name: 'storeUpload',
    component: StoreUploadView
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
