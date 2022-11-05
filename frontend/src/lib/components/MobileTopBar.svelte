<script lang="ts">
	import type { Team } from 'src/types/preview'

	import backIcon from '../../assets/images/back-icon.png'

	export let team: Team
	export let selected: boolean
	export let handleSelected: (selected: boolean) => void
</script>

<div
	class="w-full h-11 pl-[30px] pr-[30px] flex justify-between items-center bg-[#1B2D51] font-semibold text-lg"
>
	<div class="flex justify-center items-center w-9">
		{#if selected == true}
			<button
				on:click={() => {
					handleSelected(selected)
				}}
			>
				<img src={backIcon} alt="backIcon" />
			</button>
		{/if}
	</div>
	<div class="w-[300px] flex justify-center items-center">
		<span class="text-white">{team.name}</span>
	</div>
</div>
