import axios from 'axios'
import createAuthRefreshInterceptor from 'axios-auth-refresh'
import { getIdToken } from 'firebase/auth'

import { Configuration, ConfigurationParameters } from './generated/configuration'
import { DefaultApiFactory } from './generated/api'
import { adminAPIEndpoint } from './appConfig'
import { useAuthUserStore } from './stores/authUser'

const axiosInstance = axios.create({
  baseURL: adminAPIEndpoint,
  timeout: 5000,
});

// Obtain the fresh token each time the function is called
function getAccessToken() {
  return localStorage.getItem('token');
}

// Use interceptor to inject the token to requests
axiosInstance.interceptors.request.use((request: any) => {
  const authUserStore = useAuthUserStore()
  console.log("request interceptor called")

  return authUserStore.user!.getIdToken(false)
    .then((token) => {
      console.log("getIdToken callback resulted in token: " + token)
      request.headers['Authorization'] = 'Bearer ' + token;
      return request;
    })
});

// Function that will be called to refresh authorization
// const refreshAuthLogic = (failedRequest: any) => {
//   const authUserStore = useAuthUserStore()
//   console.log("refreshAuthLogic called")

//   return authUserStore.user!.getIdToken(false)
//     .then((token) => {
//       console.log("getIdToken callback resulted in token: " + token)
//       failedRequest.response.config.headers['Authorization'] = 'Bearer ' + token;
//       return Promise.resolve();
//     })
// }

// Instantiate the interceptor
// createAuthRefreshInterceptor(axiosInstance, refreshAuthLogic);

const apiConfigurationParameters = {
  basePath: adminAPIEndpoint
} as ConfigurationParameters

const apiConfiguration = new Configuration(apiConfigurationParameters)

export const api = DefaultApiFactory(apiConfiguration, undefined, axiosInstance)
