<script lang="ts">
	import { ArtWS } from 'src/store/websocket'
	import { onDestroy } from 'svelte'

	const client1 = ArtWS.connect(
		'ws://localhost:3000/ws/projector/card?token=wdvXuDOytfx84J8d',
		{
			log: true, // Log for console.warning
			reconnect: true, // Reconnect on close
			reconnectInterval: 1000, // Reconnect interval in ms
			name: 'client1', // Name for the client
		}
	)

	const unsubscribeClient1 = client1.subscribe('cd/state', (payload) => {
		console.log(payload)
	})

	onDestroy(() => {
		unsubscribeClient1()
	})

	const client2 = ArtWS.connect(
		'ws://localhost:3000/ws/projector/leaderboard?token=wdvXuDOytfx84J8d',
		{
			log: true, // Log for console.warning
			reconnect: true, // Reconnect on close
			reconnectInterval: 1000, // Reconnect interval in ms
			name: 'client2', // Name for the client
		}
	)

	const unsubscribeClient2 = client2.subscribe('lb/state', (payload) => {
		console.log(payload)
	})

	onDestroy(() => {
		unsubscribeClient2()
	})
</script>

<main>
	<h1>Hello this is Home</h1>
</main>
