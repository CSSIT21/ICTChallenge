<script lang="ts">
	import AdminContainer from 'src/lib/components/AdminContainer.svelte'
	import AdminManager from 'src/lib/components/AdminManager.svelte'
	let mode =
		JSON.parse(localStorage.getItem('mode')) === null
			? true
			: JSON.parse(localStorage.getItem('mode'))

	$: localStorage.setItem('mode', JSON.stringify(mode))
</script>

<main
	class="h-screen w-screen overflow-hidden bg-gradient-to-b from-[#3DC3B6] via-[#4F68BF] to-[#1B2D51]"
>
	<div
		class="fixed bottom-0 left-[calc(100vw - 150px)] p-2 right-[100px] flex items-center justify-center"
	>
		<button
			on:click={() => (mode = !mode)}
			class="p-2 rounded-lg bg-white hover:bg-gray-200 font-semi"
			>Switch
		</button>
	</div>
	{#if mode}
		<AdminContainer />
	{:else}
		<AdminManager />
	{/if}
</main>
