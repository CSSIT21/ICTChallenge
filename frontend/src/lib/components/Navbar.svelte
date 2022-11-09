<script lang="ts">
	import { Link } from 'svelte-navigator'
	let currentPath = window.location.pathname
</script>

<div class="flex justify-between">
	<p class="text-white text-2xl">SIT ICT Challenge</p>
	<div class="flex gap-10">
		<Link
			to="/leaderboard"
			class={currentPath == '/leaderboard'
				? 'bg-[#4279E8] bg-opacity-20 p-2 text-center rounded-lg text-[#4279E8] text-2xl'
				: 'p-2 text-center text-[#75829B] text-opacity-70 text-2xl'}
			>Leaderboard</Link
		>
		<Link
			to="/question"
			class={currentPath == '/question'
				? 'bg-[#4279E8] bg-opacity-20 p-2 text-center rounded-lg text-[#4279E8] text-2xl'
				: 'p-2 text-center text-[#75829B] text-opacity-70 text-2xl'}
			>Questions</Link
		>
	</div>
</div>
