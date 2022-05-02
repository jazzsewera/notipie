import { NotificationCardHandlers } from './handler'

export type Notification = {
  appName: string
  timestamp: string
  appImgUri?: string
  title: string
  subtitle?: string
  body?: string
  id?: string
  read?: boolean
  extUri?: string
  readUri?: string
  archiveUri?: string
  relativeTime?: string
}

export type NotificationWithHandlers = {
  notification: Notification
  handlers: NotificationCardHandlers
}
