import { NotificationCardHandlers } from './handler'

export type Notification = {
  appName: string
  timestamp: string
  appImgUri?: string
  title: string
  subtitle?: string
  body?: string
  id?: string
  relativeTime?: string
  read?: boolean
}

export type NotificationWithHandlers = {
  notification: Notification
  handlers: NotificationCardHandlers
}
