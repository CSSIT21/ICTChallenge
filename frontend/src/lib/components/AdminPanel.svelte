<script lang="ts">
	export let round = 0
</script>

<div
	class="w-[6rem] min-w-[6rem] h-full flex flex-col hover:bg-[rgba(0,0,0,.02)] opacity-80 hover:opacity-100"
	style="border-right: 1px solid rgba(217, 217, 217, .2)"
>
	<div
		class="flex w-full h-[64px] items-center justify-center font-semibold"
		style="border-bottom: 1px solid rgb(217, 217, 217); "
	>
		Round {round}
	</div>
	<div class="flex flex-1 flex-col">
		{#each [1, 2, 3, 4, 5, 6] as i}
			<div
				class="flex w-full h-[72px] flex-col items-center justify-center gap-2 px-2"
				style="border-bottom: 1px solid rgba(217, 217, 217, .5)"
			>
				<input
					type="number"
					placeholder="Score"
					class="rounded-full bg-white outline-none px-2 w-full text-center text-sm"
					style="border: 1px solid rgba(217, 217, 217,.4)"
				/>
				<input
					type="number"
					placeholder="Total"
					class="rounded-full bg-white outline-none px-2 w-full text-center text-sm"
					style="border: 1px solid rgba(217, 217, 217,.4)"
				/>
			</div>
		{/each}
	</div>
</div>
