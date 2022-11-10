<script lang="ts">
	import TopBar from '../lib/components/MobileTopBar.svelte'
	import WaitingScreen from '../lib/components/WaitingScreen.svelte'
	import ChoosingScreen from '../lib/components/ChoosingScreen.svelte'

	import type { Team } from 'src/types/preview'
	import { current } from 'src/store/system'

	let team: Team = {
		name: 'muumel Team',
		school: 'KMUTT',
	}
</script>

<main class=" flex justify-center items-center flex-col">
	<div class=" w-[390px] h-[880px] pt-1 bg-[#1B2D51]">
		<TopBar {team} />
		<div>
			{#if $current == 1}
				<ChoosingScreen />
			{:else}
				<WaitingScreen />
			{/if}
		</div>
	</div>
</main>
