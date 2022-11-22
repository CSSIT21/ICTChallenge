<script lang="ts">
	import { selected } from 'src/store/system'

	import backIcon from '../../assets/images/back-icon.png'

	export let name: string = ''
</script>

<div
	class="w-full h-11 pl-[30px] pr-[58px] flex justify-between items-center bg-[#1B2D51] font-semibold text-lg"
>
	<div class="flex justify-center items-center w-9">
		{#if $selected != -1}
			<button
				on:click={() => {
					selected.set(-1)
					document.body.scrollIntoView()
				}}
			>
				<img src={backIcon} alt="backIcon" />
			</button>
		{/if}
	</div>
	<div class="w-[300px] flex justify-center items-center">
		<h1 class="text-white">{name}</h1>
	</div>
</div>

<style>
	h1:focus-visible {
		outline: none;
	}
</style>
