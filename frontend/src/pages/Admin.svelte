<script lang="ts">
	import AdminContainer from 'src/lib/components/AdminContainer.svelte'
	import AdminManager from 'src/lib/components/AdminManager.svelte'
	let mode = false
</script>

<main
	class="h-screen w-screen overflow-hidden bg-gradient-to-b from-[#3DC3B6] via-[#4F68BF] to-[#1B2D51]"
>
	{#if mode}
		<AdminContainer />
	{:else}
		<AdminManager />
	{/if}
</main>
