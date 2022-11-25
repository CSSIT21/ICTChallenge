<script lang="ts">
	export let round = 0
	export let teams: any
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
		{#each teams as team}
			<div
				class="flex w-full h-[72px] flex-col items-center justify-center gap-2 px-2"
				style="border-bottom: 1px solid rgba(217, 217, 217, .5)"
			>
				<input
					type="number"
					placeholder="Score"
					value={team.scores[round - 1].change || 0}
					class="rounded-full bg-white outline-none px-2 w-full text-center text-sm"
					style="border: 1px solid rgba(217, 217, 217,.4)"
					readonly
				/>
				<input
					type="number"
					placeholder="Total"
					value={team.scores[round - 1].total || 0}
					class="rounded-full bg-white outline-none px-2 w-full text-center text-sm"
					style="border: 1px solid rgba(217, 217, 217,.4)"
					readonly
				/>
			</div>
		{/each}
	</div>
</div>
