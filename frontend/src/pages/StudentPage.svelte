<script lang="ts">
	import TopBar from '../lib/components/MobileTopBar.svelte'
	import WaitingScreen from '../lib/components/WaitingScreen.svelte'
	import ChoosingScreen from '../lib/components/ChoosingScreen.svelte'

	import { ArtWS } from 'src/store/websocket'
	import { current } from 'src/store/system'
	import { onDestroy } from 'svelte'

	let name: string = ''
	const client = ArtWS.connect(
		'ws://ictc-int.sit.kmutt.ac.th:3000/ws/student?token=XrO3ole8bS83OQ3p',
		{
			log: true, // Log for console.warning
			reconnect: true, // Reconnect on close
			reconnectInterval: 1000, // Reconnect interval in ms
			name: 'client', // Name for the client
		}
	)

	const unsubscribeclient = client.subscribe('st/turn', (payload) => {
		console.log('payload')
		console.log(payload)
		name = payload.name
		if (payload.current) {
			$current = 1
		} else {
			$current = -1
		}
	})

	onDestroy(() => {
		unsubscribeclient()
	})
</script>

<main class=" flex justify-center items-center flex-col">
	<div class=" w-[390px] h-[880px] pt-1 bg-[#1B2D51]">
		<TopBar {name} />
		<div>
			{#if $current == 1}
				<ChoosingScreen />
			{:else}
				<WaitingScreen />
			{/if}
		</div>
	</div>
</main>
