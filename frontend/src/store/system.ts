// system.ts
// An extremely simple external store
import { writable } from 'svelte/store'

export const number = writable(0)
