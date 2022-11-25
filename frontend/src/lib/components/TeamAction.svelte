<script lang="ts">
	export let id
	export let onScoreChanged: ((id: number, score: number) => void) | undefined

	export let choose = null

	function chooseTeam(score) {
		return () => {
			choose = score
			onScoreChanged(id, score)
		}
	}
</script>

<div
	class="flex justify-evenly items-center gap-1 h-[72px]"
	style="border-bottom: 1px solid rgb(217, 217, 217); "
>
	<button
		class="{choose === -1
			? 'opacity-100'
			: 'opacity-20'} bg-red-500 hover:bg-red-600 text-white flex p-2 w-[24px] h-[24px] justify-center items-center rounded-full"
		on:click={chooseTeam(-1)}>-</button
	>
	<button
		class="{choose === 0
			? 'opacity-100'
			: 'opacity-20'} bg-yellow-500 hover:bg-yellow-600 text-white flex p-2 w-[24px] h-[24px] justify-center items-center rounded-full"
		on:click={chooseTeam(0)}>o</button
	>
	<button
		class="{choose === 1
			? 'opacity-100'
			: 'opacity-20'} bg-teal-500 hover:bg-teal-600 text-white flex p-2 w-[24px] h-[24px] justify-center items-center rounded-full"
		on:click={chooseTeam(1)}>+</button
	>
</div>
