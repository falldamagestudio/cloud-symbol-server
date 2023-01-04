import axios from 'axios'

import { Configuration, ConfigurationParameters } from './generated/configuration'
import { DefaultApiFactory } from './generated/api'
import { adminAPIEndpoint } from './appConfig'
import { useAuthUserStore } from './stores/authUser'

const axiosInstance = axios.create({
  baseURL: adminAPIEndpoint,
  timeout: 5000,
});

// Use interceptor to inject the token to requests
axiosInstance.interceptors.request.use((request: any) => {
  const authUserStore = useAuthUserStore()

  return authUserStore.user!.getIdToken(false)
    .then((token) => {
      request.headers['Authorization'] = 'Bearer ' + token;
      return request;
    })
});

const apiConfigurationParameters = {
  basePath: adminAPIEndpoint
} as ConfigurationParameters

const apiConfiguration = new Configuration(apiConfigurationParameters)

export const api = DefaultApiFactory(apiConfiguration, undefined, axiosInstance)
