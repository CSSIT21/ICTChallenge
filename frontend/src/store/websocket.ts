import { onDestroy } from 'svelte'
import { writable } from 'svelte/store'

export type IMessageType =
	| 'cd/state'
	| 'lb/state'
	| 'lb/podium'
	| 'lb/turn'
	| 'cd/open'
	| 'cd/countdown'
	| 'lb/preview'
	| 'lb/preview/add'
	| 'st/turn'
	| 'lb/rankings'
	| 'lb/highlighted'

type IMessageEventHandler<T = any> = (
	payload: T
) => void | ((payload: T) => Promise<void>)

type IMessagePayload<T = any> = {
	event: string
	payload: T
}

export type WebSocketInstanceOptions = {
	url?: string
	log?: boolean
	name?: string
	reconnect?: boolean
	reconnectInterval?: number
}

class WebSocketInstance {
	subscribers: [IMessageType, IMessageEventHandler][] = []
	cleanUps: Function[] = []
	socket: WebSocket = null

	private isDev = false
	private name = 'ws'
	private url = 'ws://localhost:8080'
	private reconnect = false
	private reconnectInterval = 1000
	private static websocketId = 0

	constructor(options) {
		this.isDev = options?.log || false
		this.name = options?.name || `ws-${WebSocketInstance.websocketId++}`
		this.url = options?.url || this.url
		this.reconnect = options?.reconnect || this.reconnect
		this.reconnectInterval =
			options?.reconnectInterval || this.reconnectInterval
	}

	connect(url: string = this.url) {
		this.dispose()
		this.socket = new WebSocket(url)
		this.socket.onopen = () => {
			if (this.isDev) console.warn(`[WS:${this.name}] Connected to `, url)
		}

		this.socket.onmessage = (event) => {
			this.subscribers.forEach((subscriber) => {
				const [type, handler] = subscriber
				const { event: messageType, payload } = JSON.parse(
					event.data
				) as IMessagePayload
				if (this.isDev)
					console.warn(`[WS:${this.name}] Received: `, event, payload)
				if (type === messageType) {
					handler(payload)
				}
			})
		}

		this.socket.onclose = (event) => {
			if (this.isDev) console.warn(`[WS:${this.name}] Disconnected`)
			if (this.reconnect) {
				setTimeout(() => {
					this.connect()
				}, this.reconnectInterval)
			}
		}
	}

	subscribe<T = any>(
		event: IMessageType,
		subscriber: IMessageEventHandler<T>
	) {
		this.subscribers.push([event, subscriber])
		const cleanUp = () => {
			this.subscribers = this.subscribers.filter(
				([, s]) => s !== subscriber
			)
		}

		this.cleanUps.push(cleanUp)

		return cleanUp
	}

	dispose() {
		if (this.socket) {
			this.socket.close()
			this.cleanUps.forEach((cleanUp) => cleanUp())
			this.cleanUps = []
			this.subscribers = []
			this.socket = null
		}
	}
}

export class ArtWS {
	static connect(url: string, options: WebSocketInstanceOptions = {}) {
		const socketLike = new WebSocketInstance(options)
		socketLike.connect(url)
		onDestroy(() => {
			socketLike.dispose()
		})
		return socketLike
	}
}

// export const socket = writable(new WebSocketStore())
