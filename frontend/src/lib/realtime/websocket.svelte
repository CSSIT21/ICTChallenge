<script lang="ts">
	import { onMount } from 'svelte'
	export let url: string

	$: ws = new WebSocket(url)
	function onUrlChange(url: string) {
		if (ws) {
			ws.close()
		}
		ws = new WebSocket(url)
		ws.onmessage = (event) => {
			console.log(event.data)
		}

		ws.onopen = () => {
			console.log('connected')
		}

		ws.onclose = () => {
			console.log('disconnected')
		}
	}

	$: onUrlChange(url)
</script>
