<script lang="ts">
	import TopBar from '../lib/components/MobileTopBar.svelte'
	import WaitingScreen from '../lib/components/WaitingScreen.svelte'
	import ChoosingScreen from '../lib/components/ChoosingScreen.svelte'

	import type { Team } from 'src/types/preview'

	let current: boolean = true
	let selected: boolean = true
	let team: Team = {
		name: 'muumel Team',
		school: 'KMUTT',
	}

	const handleSelected = (selected: boolean) => {
		selected = true
		console.log(selected)
	}
</script>

<main class=" flex justify-center items-center flex-col">
	<div class=" w-[390px] h-[844px] bg-[#1B2D51]">
		<TopBar {selected} {team} {handleSelected} />

		<div>
			{#if current == true}
				<ChoosingScreen {handleSelected} {selected} />
			{:else}
				<WaitingScreen />
			{/if}
		</div>
	</div>
</main>
