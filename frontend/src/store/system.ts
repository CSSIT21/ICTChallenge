// system.ts
// An extremely simple external store
import { writable } from 'svelte/store'

export const number = writable(0)
export const selected = writable(-1)
export const current = writable(0)
