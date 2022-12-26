import { Configuration, ConfigurationParameters } from './generated/configuration'
import { DefaultApiFactory } from './generated/api'
import { adminAPIEndpoint } from './appConfig'

const apiConfigurationParameters = {
    basePath: adminAPIEndpoint
} as ConfigurationParameters

const apiConfiguration = new Configuration(apiConfigurationParameters)

export const api = DefaultApiFactory(apiConfiguration)
