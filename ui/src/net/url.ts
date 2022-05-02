import { config } from '../config'

const { root, push, webSocket, notifications } = config.endpointConfig
const { httpScheme, wsScheme, address, port, prefix } = config.proxyConfig
const host = `${address}:${port}`

type Scheme = typeof httpScheme | typeof wsScheme

const getUrl = (scheme: Scheme, target: string): string => `${scheme}://${host}${prefix}${target}`

export const rootUrl = getUrl(httpScheme, root)
export const pushUrl = getUrl(httpScheme, push)
export const wsUrl = getUrl(wsScheme, webSocket)
export const notificationsUrl = getUrl(httpScheme, notifications)
